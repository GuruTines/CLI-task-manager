package tools

import (
	"bufio"
	"fmt"
)

func GetConsoleText(scanner *bufio.Scanner) string {
	if ok := scanner.Scan(); !ok {
		if err := scanner.Err(); err != nil {
			fmt.Printf("Возникла ошибка: %v\n", err)
			return ""
		} else {
			fmt.Println("Достигнут конец строки")
			return ""
		}
	}

	return scanner.Text()
}
