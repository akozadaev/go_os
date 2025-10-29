package main

import (
	"fmt"
	"os"
)

func main() {
	fileinfo, err := os.Stat("example.txt")
	if os.IsNotExist(err) {
		fmt.Println("Файл не существует.")
	} else {
		fmt.Println(fileinfo.Name())
		fmt.Println(fileinfo.Size())
		fmt.Println("Файл существует.")
	}
}
