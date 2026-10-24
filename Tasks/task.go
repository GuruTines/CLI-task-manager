package tasks

import (
	"bufio"
	"errors"
	"fmt"
	tools "task-tracker/Tools"
	"time"
	"unicode/utf8"
)

type Comment struct {
	Author      string
	Desc        string
	DateComment time.Time
}

type ChangeInfo struct {
	ChangeDesc string
	ChangeDate time.Time
}

type Task struct {
	ID        int
	Title     string
	Desc      string
	Author    string
	Executor  string
	Priority  string
	Status    string
	DateStart time.Time
	DateEnd   time.Time

	Tags          []string
	Comments      []Comment
	ChangeHistory []ChangeInfo
}

var priorityArray = [3]string{
	"Low",
	"Medium",
	"High",
}

var statusArray = [5]string{
	"New",
	"In Progress",
	"Review",
	"Done",
	"Cancelled",
}

var tasksList []*Task = make([]*Task, 0)

func PrintPriorityTask() {
	for index, priority := range priorityArray {
		fmt.Println(index, priority)
	}
}

func PrintStatusTask() {
	for index, status := range statusArray {
		fmt.Println(index, status)
	}
}

// Валидация полей структуры задач
func ValidateShortText(text string) (string, error) {
	count := utf8.RuneCountInString(text)

	if count > 64 {
		return "", errors.New("Слишком длинное название")
	} else if count < 4 {
		return "", errors.New("Слишком короткое название")
	}

	return text, nil
}

func ValidateLongText(desc string) (string, error) {
	count := utf8.RuneCountInString(desc)

	if count > 255 {
		return "", errors.New("Описание слишком длинное")
	}

	return desc, nil
}

func ValidatePriorityIndex(indexPriority int) (int, error) {
	lenPriorityArray := len(priorityArray)

	if indexPriority < 0 || indexPriority > (lenPriorityArray-1) {
		return 0, errors.New("Индекс выходит за пределы массива приоритетов")
	}

	return indexPriority, nil
}

func ValidateDate(dateEnd time.Time) (time.Time, error) {
	if dateEnd.Before(time.Now()) {
		return time.Now(), errors.New("Дата дедлайна не может быть в прошлом")
	}
	return dateEnd, nil
}

// Добавление комментария в
// список комментариев структуры задачи
func (task *Task) AddComment(comment Comment) error {
	task.Comments = append(task.Comments, comment)
	return nil
}

// Конструктор комментария
func NewComment(author, desc string) Comment {
	comment := Comment{
		Author:      author,
		Desc:        desc,
		DateComment: time.Now(),
	}

	return comment
}

// Конструктор задачи
func NewTask(title,
	desc,
	author,
	executor string,
	priority int,
	dateEnd time.Time,
	tags []string) Task {

	task := Task{
		ID:        len(tasksList) + 1,
		Title:     title,
		Desc:      desc,
		Author:    author,
		Executor:  executor,
		Priority:  priorityArray[priority],
		Status:    statusArray[0],
		DateStart: time.Now(),
		DateEnd:   dateEnd,

		Tags: tags,
	}

	tasksList = append(tasksList, &task)

	return task
}

// Функции вывода информации
// об задаче/задачах
func PrintAllTasks() {
	fmt.Printf("%-3s %-30s %-15s %-15s %-10s\n", "ID", "Название", "Статус", "Приоритет", "Исполнитель")
	for _, task := range tasksList {
		fmt.Printf("%-3d %-30s %-15s %-15s %-10s\n", task.ID, task.Title, task.Status, task.Priority, task.Executor)
	}
	fmt.Println()
}

func PrintAllInfoTask(task Task) {
	fmt.Printf("%-15s %d\n", "ID", task.ID)
	fmt.Printf("%-15s %s\n", "Название", task.Title)
	fmt.Printf("%-15s %s\n", "Описание", task.Desc)
	fmt.Printf("%-15s %s\n", "Автор", task.Author)
	fmt.Printf("%-15s %s\n", "Исполнитель", task.Executor)
	fmt.Printf("%-15s %s\n", "Приоритет", task.Priority)
	fmt.Printf("%-15s %s\n", "Статус", task.Status)
	fmt.Printf("%-15s %s\n", "Дата создания", task.DateStart.Format("02-01-2006 15:04"))
	fmt.Printf("%-15s %s\n", "Deadline", task.DateEnd.Format("02-01-2006 15:04"))

	fmt.Printf("%-15s %v\n", "Тэги", task.Tags)

	fmt.Printf("%-15s\n", "Комментарии")
	if len(task.Comments) == 0 {
		fmt.Println("Комментариев нет")
	} else {
		for index, comment := range task.Comments {
			fmt.Printf("%d. %-5s\t%-5s\t%v\n", index+1, comment.Author, comment.Desc, comment.DateComment.Format("02-01-2006 15:04"))
		}
	}

	fmt.Printf("%-15s %v\n", "История изменений", task.ChangeHistory)
	if len(task.ChangeHistory) == 0 {
		fmt.Println("Изменений нет")
	} else {
		for index, change := range task.ChangeHistory {
			fmt.Printf("%d. %-5s\t%v\n", index+1, change.ChangeDesc, change.ChangeDate.Format("02-01-2006 15:04"))
		}
	}
	fmt.Println()
}

// Функции поиска по критериям
func SearchByID(taskID int) ([]*Task, error) {
	foundTask := []*Task{}

	if count := len(tasksList); taskID <= 0 || taskID > count {
		return foundTask, errors.New("ID выходит за пределы числа задач")
	}

	for _, task := range tasksList {
		if task.ID == taskID {
			foundTask = append(foundTask, task)
			return foundTask, nil
		}
	}

	return foundTask, errors.New("Задачи с таким ID нет")
}

func SearchByTitle(title string) ([]*Task, error) {
	foundTask := []*Task{}

	for _, task := range tasksList {
		if task.Title == title {
			foundTask = append(foundTask, task)
			return foundTask, nil
		}
	}

	return foundTask, errors.New("По указанному названию задачи нет")
}

func SearchByExecutor(executor string) ([]*Task, error) {
	foundTask := []*Task{}

	for _, task := range tasksList {
		if task.Executor == executor {
			foundTask = append(foundTask, task)
			return foundTask, nil
		}
	}

	return foundTask, errors.New("Данный исполнитель не выполняет задачи")
}

func SearchByTag(tag string) ([]*Task, error) {
	foundTask := []*Task{}

	for _, task := range tasksList {
		for _, currTag := range task.Tags {
			if len(task.Tags) == 0 {
				break
			}

			if currTag == tag {
				foundTask = append(foundTask, task)
				break
			}
		}
	}

	return foundTask, nil
}

// Функции изменения задачи
func ChangeTitleField(scanner *bufio.Scanner, changeTask *Task) error {
	changeTitle := tools.GetConsoleText(scanner)

	validTitle, err := ValidateShortText(changeTitle)
	if err != nil {
		return fmt.Errorf("%s\n", err)
	}

	changeTask.Title = validTitle
	return nil
}

func ChangeDescField(scanner *bufio.Scanner, changeTask *Task) error {
	changeDesc := tools.GetConsoleText(scanner)

	valiDesc, err := ValidateLongText(changeDesc)
	if err != nil {
		return fmt.Errorf("%s\n", err)
	}

	changeTask.Title = valiDesc
	return nil
}
