package main

import (
	"fmt"
)

func sis(s string) int {
	switch s {
	case "r":
		return 1
	case "ar":
		return 0
	default:
		return -1
	}
} // Функция проверяет корректность выбора системы счисления(РИМСКАЯ/АРАБСКАЯ)

func valid_Ar_Value_X_Y(x, y int) (int, bool) {
	var res bool
	var value int
	if -10 <= x && x <= 10 {
		if -10 <= y && y <= 10 {
			res = true
			value = 0
		} else {
			res = false
			value = y
		}
	} else {
		res = false
		value = x
	}
	return value, res
} // Функция проверяет вхождение введёных чисел в заданный интервал(АРАБСКИЕ цифры)

func valid_Value_Oper(z string) bool {
	var res bool
	a := [4]string{"+", "-", "*", "/"}
	for _, v := range a {
		if z == v {
			res = true
			break
		} else {
			res = false
		}
	}
	return res
} // Функция проверяет валидность математической операции

func result_Ar(z string, x, y int) int {
	var result int
	if z == "/" && y == 0 {
		result = 101
	} else {
		switch z {
		case "+":
			result = x + y
		case "-":
			result = x - y
		case "*":
			result = x * y
		case "/":
			result = x / y
		}
	}
	return result
} // Функция выполняет расчёт введёного выражения в АРАБСКИХ цифрах

func valid_R_Value_X_Y(x, y string) (string, bool) {
	var res bool
	var value string
	if roman_To_Arabic(x) != 0 {
		if roman_To_Arabic(y) != 0 {
			res = true
			value = "OK"
		} else {
			res = false
			value = y
		}
	} else {
		res = false
		value = x
	}
	return value, res
} // Функция проверяет вхождение введёных чисел в заданный интервал(РИМСКИЕ цифры)

func roman_To_Arabic(x string) int {
	m := map[string]int{
		"X":    10,
		"IX":   9,
		"VIII": 8,
		"VII":  7,
		"VI":   6,
		"V":    5,
		"IV":   4,
		"III":  3,
		"II":   2,
		"I":    1}
	return m[x]
} // Функция выполняет конвертацию из РИМСКИХ в АРАБСКИЕ цифры

func result_R(z, x, y string) string {
	roman_Digit := [9]string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	arabic_Digit := [9]int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman_Res := result_Ar(z, roman_To_Arabic(x), roman_To_Arabic(y))
	result := ""
	if x <= y && (z == "-" || z == "/") {
		fmt.Println("ошибка №_006: решение невозможно (в римской системе счисления нет 'ноля' и отрицательных чисел) перезапустите приложение и попробуйте снова")
		result = "нет"
	} else {
		for i := 0; roman_Res > 0; i++ {
			if roman_Res >= arabic_Digit[i] {
				result += roman_Digit[i]
				roman_Res -= arabic_Digit[i]
				i -= 1
			}
		}
	}
	return result
} // Функция выполняет расчёт введённого пользователем выражения в РИМСКИХ цифрах

func main() {
	var s, z, a, b string
	var x, y int

	fmt.Println("Здравствуйте!")
	fmt.Println("Выберите систему расчёта в римских(r) или арабских(ar) цифрах и нажмите 'Enter'. (r / ar) ?")
	fmt.Scanln(&s)

	for i := 0; sis(s) < 0; i++ {
		fmt.Println("ошибка №_001: недопустимое значение (только 'r' или 'ar'!) попробуйте снова:")
		fmt.Scanln(&s)
	} // Цикл останавливает выполнение программы до тех пор, пока пользователь не введёт корректные данные

	fmt.Println("Введите выражение в формате 'число №1' (+, -, *, /) 'число №2'")
	fmt.Println("Каждое значение отделяйте нажатием клавиши 'пробел'. После нажмите 'Enter'")

	if s == "ar" {
		fmt.Println("ВНИМАНИЕ! Диапазон вводимых чисел ограничен -10 <= 'число №1/число №2' <= 10")

		_, err := fmt.Scanln(&x, &z, &y)
		value_Ar, valid_Ar := valid_Ar_Value_X_Y(x, y)
		if err != nil {
			fmt.Println(fmt.Errorf("ошибка №_007: неверный формат ввода (%s) перезапустите приложение и введите выражение в формате 'число №1' (+, -, *, /) 'число №2'", err))
		} else {
			if valid_Value_Oper(z) == false {
				fmt.Printf("ошибка №_003: неизвестная операция (%s) допустимые операции (+, -, *, /) перезапустите приложение и попробуйте снова", z)
			} else if valid_Ar == false {
				fmt.Printf("ошибка №_002: число %d вне диапазона (-10 <= 'число №1/число №2' <= 10) перезапустите приложение и попробуйте снова", value_Ar)
			} else if result_Ar(z, x, y) == 101 {
				fmt.Println("ошибка №_008: деление на ноль запрещено!")
			} else {
				fmt.Printf("Результат: %d", result_Ar(z, x, y))
			}
		}
	} else if s == "r" {
		fmt.Println("ВНИМАНИЕ! Диапазон вводимых чисел ограничен I <= 'число №1/число №2' <= X")

		_, err := fmt.Scanln(&a, &z, &b)
		value_R, valid_R := valid_R_Value_X_Y(a, b)
		if err != nil {
			fmt.Println(fmt.Errorf("ошибка №_007: неверный формат ввода (%s) перезапустите приложение и введите выражение в формате 'число №1' (+, -, *, /) 'число №2'", err))
		} else {
			if valid_Value_Oper(z) == false {
				fmt.Printf("ошибка №_003: неизвестная операция (%s) допустимые операции (+, -, *, /) перезапустите приложение и попробуйте снова", z)
			} else if valid_R == false {
				fmt.Printf("ошибка №_004: число '%s' вне диапазона (I <= 'Число №1/Число №2' <= X) перезапустите приложение и попробуйте снова", value_R)
			} else {
				fmt.Printf("Результат: %s", result_R(z, a, b))
			}
		}

	} else {
		fmt.Println("ошибка №_005: (неизвестная ошибка) Вам удалось найти необработаный сценарий ошибки; по возможности сообщите об этом разработчику для устранения")
	} // Основной блок кода программы для расчёта в РИМСКИХ или АРАБСКИХ цифрах
}
