package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("example.txt")
	if os.IsNotExist(err) {
		fmt.Println("Файл не существует.")
	} else {
		fmt.Println("Файл существует.")
	}
}
