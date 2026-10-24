package tools

import (
	"fmt"
)

func PrintMenu() {
	fmt.Println("1. Создать задачу")
	fmt.Println("2. Добавить комментарий к задаче")
	fmt.Println("3. Показать все задачи")
	fmt.Println("4. Найти задачу")
	fmt.Println("5. Изменить задачу")
	fmt.Println("6. Удалить задачу")
	fmt.Println("7. Фильтрация")
	fmt.Println("8. Статистика")
	fmt.Println("9. Сохранить")
	fmt.Println("10. Загрузить")
	fmt.Println("0. Выход")
}

func PrintSearchMenu() {
	fmt.Println("1. Поиск по ID")
	fmt.Println("2. Поиск по названию")
	fmt.Println("3. Поиск по исполнителю")
	fmt.Println("4. Поиск по тэгу")
	fmt.Println("5. Поиск по статусу")
	fmt.Println("0. Выход")
}

func PrintTaskMenu() {
	fmt.Println("1. Название")
	fmt.Println("2. Описание")
	fmt.Println("3. Автор")
	fmt.Println("4. Исполнитель")
	fmt.Println("5. Приоритет")
	fmt.Println("6. Deadline")
	fmt.Println("0. Выход")
}
