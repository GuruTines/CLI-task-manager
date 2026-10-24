package handlers

import (
	"bufio"
	"fmt"
	"strings"
	tasks "task-tracker/Tasks"
	tools "task-tracker/Tools"
	"time"
)

func CreateTaskHandler(scanner *bufio.Scanner) tasks.Task {
	var err error = nil

	var title string = ""
	var description string = ""
	var author string = ""
	var executor string = ""
	var priority int = 0
	var dateEnd time.Time

	for {
		fmt.Println("Название задачи: ")
		title, err = WriteAndValidateShortText(scanner)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	for {
		fmt.Println("Описание задачи: ")
		description, err = WriteAndValidateLongText(scanner)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	for {
		fmt.Println("Автор задачи: ")
		author, err = WriteAndValidateShortText(scanner)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	for {
		fmt.Println("Исполнитель задачи: ")
		executor, err = WriteAndValidateShortText(scanner)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	for {
		fmt.Println("Приоритет задачи: ")
		tasks.PrintPriorityTask()
		priority, err = WriteAndValidateIntTypes(scanner)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	for {
		fmt.Println("Срок выполнения -> DD-MM-YYYY HH:MM: ")
		dateEnd, err = WriteAndValidateTime(scanner)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	fmt.Println("Тэги к задаче -> [TAG1 TAG2 ... TAGn]: ")
	tags := strings.Fields(tools.GetConsoleText(scanner))

	return tasks.NewTask(title, description, author, executor, priority, dateEnd, tags)
}

func SearchTask(scanner *bufio.Scanner) []*(tasks.Task) {
	var err error = nil
	task := []*(tasks.Task){}
	searchOption := tools.GetConsoleText(scanner)

	switch searchOption {
	case "1":
		for {
			fmt.Println("Введите ID задачи: ")
			taskID, err := WriteAndValidateIntTypes(scanner)
			if err != nil {
				fmt.Println(err)
				continue
			}

			task, err = tasks.SearchByID(taskID)
			if err != nil {
				fmt.Println(err)
				continue
			}

			break
		}

	case "2":
		for {
			fmt.Println("Введите название задачи: ")
			taskTitle := tools.GetConsoleText(scanner)

			task, err = tasks.SearchByTitle(taskTitle)
			if err != nil {
				fmt.Println(err)
				continue
			}

			break
		}

	case "3":
		for {
			fmt.Println("Введите исполнителя задачи: ")
			taskExecutor := tools.GetConsoleText(scanner)

			task, err = tasks.SearchByExecutor(taskExecutor)
			if err != nil {
				fmt.Println(err)
				continue
			}

			break
		}

	case "4":
		for {
			fmt.Println("Введите тэг задачи: ")
			taskTag := tools.GetConsoleText(scanner)

			task, err = tasks.SearchByTag(taskTag)
			if err != nil {
				fmt.Println(err)
				continue
			}

			break
		}

	case "0":

	}

	return task
}
