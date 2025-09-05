package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type DiaryEntry struct {
	Date    time.Time
	Title   string
	Content string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var entries []DiaryEntry
	var userName string = "Анна"
	fmt.Printf("🗓️ Добро пожаловать в личный дневник, %s!\n", userName)
	currentDate := time.Now().Format("02.01.2006 15:04")
	fmt.Printf("📅 Дата: %s\n", currentDate)

	for {
		fmt.Println("\n=== МЕНЮ ДНЕВНИКА ===")
		fmt.Println("1. Добавить новую запись")
		fmt.Println("2. Показать все записи")
		fmt.Println("3. Выход")
		fmt.Print("Выберите действие (1-3): ")

		choice, _ := reader.ReadString('\n')
		choice = choice[:len(choice)-1]

		switch choice {
		case "1":
			fmt.Print("Введите заголовок записи: ")
			title, _ := reader.ReadString('\n')
			title = title[:len(title)-1]

			fmt.Print("Введите содержимое записи: ")
			content, _ := reader.ReadString('\n')
			content = content[:len(content)-1]

			newEntry := DiaryEntry{
				Date:    time.Now(),
				Title:   title,
				Content: content,
			}

			entries = append(entries, newEntry)
			entriesCount := len(entries)
			fmt.Printf("📊Запись добавлена! Количество записей: %d\n", entriesCount)
		case "2":
			if len(entries) == 0 {
				fmt.Println("Записей пока нет.")
			} else {
				fmt.Printf("📚 Всего записей:%d\n", len(entries))
				fmt.Println("==================")
				for i, entry := range entries {
					fmt.Printf("Запись №%d\n", i+1)
					fmt.Printf("Дата: %s\n", entry.Date.Format("02.01.2006 15:04"))
					fmt.Printf("Заголовок: %s\n", entry.Title)
					fmt.Printf("Содержание: %s\n", entry.Content)
					fmt.Println("-----------")
				}
			}
		case "3":
			fmt.Println("👋Пока-пока!")
			return
		default:
			fmt.Println("❌Неверный выбор. Попробуйте снова.")
		}
	}

}
