package main

import (
	"bufio"
	"fmt"
	"os"
	handlers "task-tracker/Handlers"
	tasks "task-tracker/Tasks"
	tools "task-tracker/Tools"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

outerLoop:
	for {
		tools.PrintMenu()
		text := tools.GetConsoleText(scanner)

		switch text {
		// Создать задачу
		case "1":
			_ = handlers.CreateTaskHandler(scanner)

		// Добавить комментарий
		case "2":
			fmt.Println("Выберите способ поиска задачи:\nПри использовании тэга, комментарий будет добавлен\nко всем полученным задачам")
			tools.PrintSearchMenu()

			taskPtr := handlers.SearchTask(scanner)
			count := len(taskPtr)

			if count == 0 {
				fmt.Println("Задача не выбрана")
				break
			}

			commentAuthor, commentText := "", ""
			var err error = nil

			for {
				fmt.Println("Автор комментария: ")
				commentAuthor, err = handlers.WriteAndValidateShortText(scanner)
				if err != nil {
					fmt.Println(err)
					continue
				}
				break
			}

			for {
				fmt.Println("Текст комментария: ")
				commentText, err = handlers.WriteAndValidateShortText(scanner)
				if err != nil {
					fmt.Println(err)
					continue
				}
				break
			}

			newComment := tasks.NewComment(commentAuthor, commentText)

			if count == 1 {
				taskPtr[0].AddComment(newComment)
			} else {
				for _, currTask := range taskPtr {
					currTask.AddComment(newComment)
				}
			}

		// Показать ВСЕ задачи
		case "3":
			tasks.PrintAllTasks()

		// Поиск задач по критериям
		case "4":
			fmt.Println("Выберите способ поиска задачи: ")
			tools.PrintSearchMenu()

			taskPtr := handlers.SearchTask(scanner)
			if count := len(taskPtr); count == 0 {
				fmt.Println("Задача не выбрана")
			} else {
				for _, currTask := range taskPtr {
					tasks.PrintAllInfoTask(*currTask)
					fmt.Println()
				}
			}

		case "5":
			fmt.Println("Выберите способ поиска задачи: ")
			tools.PrintSearchMenu()

			fmt.Println("Новое название задачи: ")

			taskPtr := handlers.SearchTask(scanner)
			if count := len(taskPtr); count == 0 {
				fmt.Println("Задача не выбрана")
			} else {
				for _, currTask := range taskPtr {
					tasks.ChangeTitleField(scanner, currTask)
					fmt.Println()
				}
			}

		case "0":
			break outerLoop
		}
	}
}
