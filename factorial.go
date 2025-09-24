//Программа факториал числа
//Итоговая версия 0.2



package main

import (
	"fmt"
)


func main() {
	fmt.Println()
	fmt.Println("Приведствую! Программа Факториал числа запущенна.")
	fmt.Println("Для выхода введите ноль")
	fmt.Println("----------")
	for {
		var numberUser, summator int
		summator = 1
		fmt.Println("Введи число, а программа выведет его вакториал:")
		_, err := fmt.Scan(&numberUser)

		if err != nil {
			fmt.Println("Увы... Произшла ошибка чтения", err)
			fmt.Println("Пожалуйста введите числовое значение!")
			fmt.Println("----------")

			var bak string
			fmt.Scan(&bak)
			continue
		}

		if numberUser == 0 {
			fmt.Println("Выход из программы...")
			break
		}

		for i := 1; i <= numberUser; i++ {
			summator *= i
		}
		fmt.Println("----------")
		fmt.Println("Факториал числа", numberUser, "=", summator)
		fmt.Println("----------")

	}
}
