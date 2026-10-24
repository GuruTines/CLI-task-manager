package handlers

import (
	"bufio"
	"fmt"
	"strconv"
	tasks "task-tracker/Tasks"
	tools "task-tracker/Tools"
	"time"
)

// Функция подходит для обработки
// Названия
// Автора
// Исполнителя
func WriteAndValidateShortText(scanner *bufio.Scanner) (string, error) {
	textFromConsole := tools.GetConsoleText(scanner)

	validText, err := tasks.ValidateShortText(textFromConsole)
	if err != nil {
		return "", fmt.Errorf("%s\n", err)
	}

	return validText, nil
}

// Функция подходит для обработки
// Описания
func WriteAndValidateLongText(scanner *bufio.Scanner) (string, error) {
	textFromConsole := tools.GetConsoleText(scanner)

	validText, err := tasks.ValidateLongText(textFromConsole)
	if err != nil {
		return "", fmt.Errorf("%s\n", err)
	}

	return validText, nil
}

// Функция подходит для обработки
// числовых типов
// Выбор приоритета
// ID задачи
func WriteAndValidateIntTypes(scanner *bufio.Scanner) (int, error) {
	textFromConsole := tools.GetConsoleText(scanner)

	intTextType, atoiErr := strconv.Atoi(textFromConsole)
	if atoiErr != nil {
		return 0, fmt.Errorf("%s\n", atoiErr)
	}

	return intTextType, nil
}

// Работа с датой и временем
func WriteAndValidateTime(scanner *bufio.Scanner) (time.Time, error) {
	textFromConsole := tools.GetConsoleText(scanner)

	parsDateEnd, errParse := time.Parse("02-01-2006 15:04", textFromConsole)
	if errParse != nil {
		return time.Time{}, fmt.Errorf("%s\n", errParse)
	}

	validDate, err := tasks.ValidateDate(parsDateEnd)
	if err != nil {
		return time.Time{}, fmt.Errorf("%s\n", err)
	}

	return validDate, nil
}
