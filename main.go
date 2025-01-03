package GoThemePark

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ParkName    = "Happy Fun Park"
	MaxVisitors = 1000
	TicketPrice = 20.5
)

var (
	OpenTime  = time.Date(2024, time.December, 31, 9, 0, 0, 0, time.Local)
	CloseTime = time.Date(2024, time.December, 31, 22, 0, 0, 0, time.Local)
)

func main() {
	fmt.Println("Добро пожаловать в", ParkName)
	fmt.Printf("Часы работы: %v - %v\n", OpenTime.Format("15:04"), CloseTime.Format("15:04"))

	// Работа с примитивами
	visitorCount := 500
	fmt.Println("Текущие посетители:", visitorCount)

	// Работа с числами
	totalRevenue := calculateRevenue(visitorCount, TicketPrice)
	fmt.Printf("Общая выручка: %.2f$\n", totalRevenue)

	// Работа с массивами и слайсами
	attractions := []string{"Колесо обозрения", "Американские горки", "Комната страха", "Автодром"}
	printAttractions(attractions)

	// Работа с картами
	visitorAgeGroups := map[string]int{
		"Дети":       200,
		"Взрослые":   250,
		"Пенсионеры": 50,
	}
	printVisitorStats(visitorAgeGroups)

	// Работа с указателями
	highlightAttraction := "Американские горки"
	promoteAttraction(&highlightAttraction)

	// Работа с strconv
	visitorCountStr := strconv.Itoa(visitorCount)
	fmt.Println("Посетителей в строковом формате:", visitorCountStr)

	// Работа с strings
	parkDescription := strings.Join(attractions, ", ")
	fmt.Println("Описание парка:", parkDescription)

	// Работа с errors
	err := checkCapacity(visitorCount + 600)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	// Анонимная функция
	func(message string) {
		fmt.Println("Специальное объявление:", message)
	}("Скидки на билеты после 18:00!")

	// Обработчики ошибок
	handleErrors(0, 0)

	// Работа с библиотекой os
	logToFile("park.log", "Парк успешно открылся!")

	fmt.Println("\n*** Парк завершил работу! ***")
	os.Exit(0)
}

// Расчёт выручки
func calculateRevenue(visitors int, price float64) float64 {
	return float64(visitors) * price
}

// Печать аттракционов
func printAttractions(attractions []string) {
	fmt.Println("Доступные аттракционы:")
	for _, attraction := range attractions {
		fmt.Println("- ", attraction)
	}
}

// Печать статистики посетителей
func printVisitorStats(stats map[string]int) {
	fmt.Println("Статистика посетителей:")
	for group, count := range stats {
		fmt.Printf("%s: %d\n", group, count)
	}
}

// Промоция аттракциона
func promoteAttraction(attraction *string) {
	*attraction = *attraction + " - САМЫЙ ПОПУЛЯРНЫЙ!"
	fmt.Println("Промоция:", *attraction)
}

// Проверка вместимости
func checkCapacity(currentVisitors int) error {
	if currentVisitors > MaxVisitors {
		return errors.New("превышена максимальная вместимость парка")
	}
	return nil
}

// Обработчики ошибок
func handleErrors(num1, num2 int) {
	if num1 == 0 || num2 == 0 {
		fmt.Println("Ошибка: деление на ноль невозможно!")
		return
	}
	fmt.Println("Результат деления:", num1/num2)
}

// Логирование в файл
func logToFile(filename, message string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка записи лога:", err)
		return
	}
	defer f.Close()

	buffer := bytes.NewBufferString(message)
	_, err = f.Write(buffer.Bytes())
	if err != nil {
		fmt.Println("Ошибка записи лога:", err)
	}
}
