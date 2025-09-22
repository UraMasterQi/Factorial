//Программа факториал числа
//Итоговая версия 0.2

// Добавлен цвет

package main

import (
	"fmt"
)

const (
    Blue = "\033[34m"
	Red = "\033[31m"
	Green  = "\033[32m"
	Reset  = "\033[0m"
)

const (
	launs = "--------------------"
)

func main() {
	fmt.Println(Blue + launs)
	fmt.Println("Приведствую! Программа Факториал числа запущенна.")
	fmt.Println("Для выхода введите ноль")
	fmt.Println(launs + Reset)
	for {
		var numberUser, summator int
		summator = 1
		fmt.Println("Введи число, а программа выведет его вакториал:")
		_, err := fmt.Scan(&numberUser)

		if err != nil {
			fmt.Println(Red + launs)
			fmt.Println("Увы... Произшла ошибка чтения", err)
			fmt.Println("Пожалуйста введите числовое значение!")
			fmt.Println(launs + Reset)

			var bak string
			fmt.Scan(&bak)
			continue
		}

		if numberUser == 0 {
			fmt.Println(Blue + launs)
			fmt.Println("Выход из программы...")
			fmt.Println(launs + Blue)
			break
		}

		for i := 1; i <= numberUser; i++ {
			summator *= i
		}
		fmt.Println(Green + launs)
		fmt.Println("Факториал числа", numberUser, "=", summator)
		fmt.Println(launs + Reset)

	}
}
