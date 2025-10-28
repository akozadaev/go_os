package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("mydirectory")
	if os.IsNotExist(err) {
		fmt.Println("Каталог 'mydirectory' не существует.")
	} else {
		fmt.Println("Каталог 'mydirectory' существует.")
	}
}
