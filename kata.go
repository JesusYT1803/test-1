package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	render := bufio.NewReader(os.Stdin)

	// Карта для перевода римских цифр в арабские
	romeToArab := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	// Ошибки
	err1 := fmt.Errorf("Выдача паники, так как используются одновременно разные системы счисления.")
	err2 := fmt.Errorf("Выдача паники, так как в римской системе нет отрицательных чисел.")
	err4 := fmt.Errorf("Выдача паники, так как строка не является математической операцией.")
	err5 := fmt.Errorf("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).\n")

	for {
		fmt.Print("Введите значение ")
		text, _ := render.ReadString('\n') // Ждет ввода данных в формате строки
		text = strings.TrimSpace(text)     // Очищает все пустоты (пробелы, табуляции)
		parts := strings.Split(text, " ")  // Разделяем строку на части

		if len(parts) <= 1 {
			fmt.Println(err4)
			break
		}
		if len(parts) > 3 {
			fmt.Println(err5)
			break
		}

		operator := parts[1] // Вытащили оператор

		_, status1 := romeToArab[parts[0]]
		_, status2 := romeToArab[parts[2]]

		if status1 == true && status2 == true {
			num1 := romeToArab[parts[0]]
			num2 := romeToArab[parts[2]]
			result := calculate(operator, num1, num2)

			if result >= 1 {
				fmt.Println(arabicToRoman(result)) // Вывод результата
			} else {
				fmt.Println(err2)
				break
			}
		} else if status1 == false && status2 == false {
			num1, _ := strconv.Atoi(parts[0])         // Преобразование строки в число 1
			num2, _ := strconv.Atoi(parts[2])         // Преобразование строки в число 2
			result := calculate(operator, num1, num2) // Вычисляем результат операции с заданным оператором и числами
			fmt.Println(result)                       // Вывод результата
		} else {
			fmt.Println(err1)
			break
		}
	}

}

// Функция для вычисления
func calculate(operator string, num1 int, num2 int) int {
	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}
	return result
}

// Функция для преобразования арабских цифр в римские
func arabicToRoman(number int) string {
	// Карта для перевода арабских цифр в римские
	arabToRome := map[int]string{
		100: "C",
		90:  "XC",
		50:  "L",
		40:  "XL",
		10:  "X",
		5:   "V",
		1:   "I",
	}

	result := ""

	// Срез значений для итерации по карте в порядке убывания
	values := []int{100, 90, 50, 40, 10, 5, 1}

	for _, value := range values {
		for number >= value {
			result += arabToRome[value]
			number -= value
		}
	}

	return result
}
