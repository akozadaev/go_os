package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("Hello")
	n, err := os.Stdout.Write(data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\nЗаписали %d байт", n)

}
