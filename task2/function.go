package main

import (
	"fmt"
	"time"
)

// Дополните функцию из п.1 возвратом собственной ошибки
// в случае возникновения панической ситуации.
// Собственная ошибка должна хранить время обнаружения панической ситуации.
// Критерием успешного выполнения задания является наличие
// обработки созданной ошибки в функции main и вывод ее состояния в консоль

type MyErr struct {
	text string
	time time.Time
}

func New(text string) error {
	return &MyErr{text: text, time: time.Now()}
}

func (e *MyErr) Error() string {
	return fmt.Sprintf("error: %s\ntime:%s", e.text, e.time)
}

func main() {
	names := []string{
		"Вася",
		"Женя",
		"Георгий",
	}

	getName(names, 3)
}

func getName(arr []string, num int) {
	defer panicHandler()
	fmt.Println("My favorite sea creature is:", arr[num])
}

func panicHandler() {
	if v := recover(); v != nil {
		err := New("my error")
		fmt.Println(err)
	}
}
