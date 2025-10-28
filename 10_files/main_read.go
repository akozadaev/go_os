package main

import (
	"fmt"
	"os"
)

func main() {
	// Открытие файла для чтения
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Чтение данных из файла
	data := make([]byte, 1024)
	n, err := file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Прочитано %d байт: %s\n", n, string(data[:n]))
}
