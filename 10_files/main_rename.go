package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Rename("oldfile.txt", "newfile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Файл переименован.")
}
