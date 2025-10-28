package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	data := []byte("Hello, World!")
	n, err := writer.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Written %d bytes\n", n)

	writer.Flush() // Сбрасывает буфер на диск
}
