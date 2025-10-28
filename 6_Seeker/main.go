package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Перемещение указателя в начало файла
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Чтение данных из начала файла
	data := make([]byte, 10)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Read data: %s\n", string(data))
}
