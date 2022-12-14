package main

// Импортируем библиотеки вывода и конвертации
import (
	"fmt"
	"strconv"
)

// Функция проверки комиссии
func check_commission(number string) uint64 {
	// Конвертируем присланное значение из строки в число с запятой во временную переменную
	tempValue, err := strconv.ParseFloat(number, 64)
	// Если конвертация не получилась
	if err != nil {
		// Скажем об этом
		fmt.Println("Введено не число:", number)
		// И выйдем из функции
		return 0
	}
	// Если конвертация удалась
	// Умножим значение на 100, тем самым сдвинем запятую
	tempValue = tempValue * 100
	// Преобразуем значение в нужный тип данных
	result := uint64(tempValue)
	// Возвращаем корректное значение
	return result
}

// Ввиду ограничений формы входящее значение от пользователя всегда в string
// Пеменная input будет эмулировать входящее значение
func main() {
	// Проверяем маленькое значение
	input := "0.01"
	fmt.Println("Первое значение:", input)
	fmt.Println("Преобразованное значение:", check_commission(input))

	// Проверяем большое значение
	input = "99.99"
	fmt.Println("\nВторое значение:", input)
	fmt.Println("Преобразованное значение:", check_commission(input))

	// Проверяем если ввели не число
	input = "неЧисло"
	fmt.Println("\nСтроковое значение:", input)
	fmt.Println("Преобразованное значение:", check_commission(input))
}
