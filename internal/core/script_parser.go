package core

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ScriptCommand представляет команды сценария
type ScriptCommand string

const (
	CMD_CHARACTER  ScriptCommand = "CHARACTER"
	CMD_NARRATOR   ScriptCommand = "NARRATOR"
	CMD_CENTER     ScriptCommand = "CENTER"
	CMD_BACKGROUND ScriptCommand = "BACKGROUND"
	CMD_POSITION   ScriptCommand = "POSITION"
)

// LoadScript загружает и парсит файл сценария
func (s *NovelScene) LoadScript(scriptPath string) error {
	file, err := os.Open(scriptPath)
	if err != nil {
		return fmt.Errorf("failed to open script file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentCharacter *Character
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		// Пропускаем пустые строки и комментарии
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Разбираем строку на команду и содержимое
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid script format at line %d: %s", lineNumber, line)
		}

		cmd := strings.TrimSpace(parts[0])
		content := strings.TrimSpace(parts[1])

		// Обработка команд
		err := s.processScriptCommand(ScriptCommand(cmd), content, &currentCharacter, lineNumber)
		if err != nil {
			return fmt.Errorf("error at line %d: %v", lineNumber, err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading script file: %v", err)
	}

	return nil
}

// processScriptCommand обрабатывает отдельные команды сценария
func (s *NovelScene) processScriptCommand(cmd ScriptCommand, content string, currentCharacter **Character, lineNumber int) error {
	switch cmd {
	case CMD_CHARACTER:
		return s.processCharacterDefinition(content, lineNumber)

	case CMD_BACKGROUND:
		return s.processBackgroundCommand(content)

	case CMD_POSITION:
		if *currentCharacter == nil {
			return fmt.Errorf("POSITION command without active character")
		}
		return s.processPositionCommand(content, *currentCharacter)

	case CMD_NARRATOR:
		s.addDialogueLine(DialogueLine{
			Type: NarratorDialogue,
			Text: content,
		})
		*currentCharacter = nil
		return nil

	case CMD_CENTER:
		s.addDialogueLine(DialogueLine{
			Type: CenterScreenText,
			Text: content,
		})
		*currentCharacter = nil
		return nil

	default:
		// Пробуем обработать как диалог персонажа
		return s.processCharacterDialogue(string(cmd), content, currentCharacter)
	}
}

// processCharacterDefinition обрабатывает определение нового персонажа
func (s *NovelScene) processCharacterDefinition(content string, lineNumber int) error {
	parts := strings.Split(content, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid CHARACTER definition format. Expected 'Name, ImagePath'")
	}

	name := strings.TrimSpace(parts[0])
	imagePath := strings.TrimSpace(parts[1])

	if name == "" || imagePath == "" {
		return fmt.Errorf("CHARACTER name and image path cannot be empty")
	}

	s.characters[name] = &Character{
		Name:      name,
		ImagePath: imagePath,
	}

	return nil
}

// processBackgroundCommand обрабатывает команду установки фона
func (s *NovelScene) processBackgroundCommand(content string) error {
	if content == "" {
		return fmt.Errorf("BACKGROUND image path cannot be empty")
	}

	s.background = &Background{
		ImagePath: content,
	}

	return nil
}

// processPositionCommand обрабатывает команду позиционирования персонажа
func (s *NovelScene) processPositionCommand(content string, character *Character) error {
	parts := strings.Split(content, ",")
	if len(parts) != 3 {
		return fmt.Errorf("invalid POSITION format. Expected 'CharacterName, X, Y'")
	}

	charName := strings.TrimSpace(parts[0])
	if charName != character.Name {
		return fmt.Errorf("character name in POSITION doesn't match current character")
	}

	x, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return fmt.Errorf("invalid X coordinate: %v", err)
	}

	y, err := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)
	if err != nil {
		return fmt.Errorf("invalid Y coordinate: %v", err)
	}

	// Обновляем позицию в последней добавленной строке диалога
	if len(s.dialogueLines) > 0 {
		lastLine := &s.dialogueLines[len(s.dialogueLines)-1]
		if lastLine.Character == character {
			lastLine.Position = Position{X: x, Y: y}
		}
	}

	return nil
}

// processCharacterDialogue обрабатывает диалог персонажа
func (s *NovelScene) processCharacterDialogue(characterName, content string, currentCharacter **Character) error {
	character, exists := s.characters[characterName]
	if !exists {
		return fmt.Errorf("undefined character: %s", characterName)
	}

	s.addDialogueLine(DialogueLine{
		Type:      CharacterDialogue,
		Character: character,
		Text:      content,
	})

	*currentCharacter = character
	return nil
}

// addDialogueLine добавляет новую строку диалога в сцену
func (s *NovelScene) addDialogueLine(line DialogueLine) {
	s.dialogueLines = append(s.dialogueLines, line)
}
