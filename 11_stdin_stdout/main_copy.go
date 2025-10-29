package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	sourceFile, err := os.Open("source.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sourceFile.Close()

	destFile, err := os.Create("destination.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer destFile.Close()

	n, err := io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Скопировано %d байт\n", n)
}
