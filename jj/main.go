package main

import (
	"fmt"
	"time"
)

// count читает числа из канала, возводит их в квадрат и выводит результат
func count(ch <-chan int) {
	for num := range ch {
		fmt.Printf("Number: %d, Square: %d\n", num, num*num)
	}
}

func main() {
	ch := make(chan int)

	// Запускаем функцию count в отдельной горутине
	go count(ch)

	// Отправляем числа в канал
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	// Закрываем канал
	close(ch)

	// Даем время завершиться горутине count
	time.Sleep(time.Second)
}
