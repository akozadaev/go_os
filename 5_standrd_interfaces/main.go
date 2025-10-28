package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Пример использования io.Reader и io.Writer
	reader := strings.NewReader("Hello, World!")
	writer := &strings.Builder{}

	io.Copy(writer, reader)
	fmt.Println(writer.String())

	// Пример использования io.Closer
	var r io.Reader = strings.NewReader("Some resource to close")

	if closer, ok := r.(io.Closer); ok {
		closer.Close()
		fmt.Println("Resource closed")
	} else {
		fmt.Println("This reader does not need to be closed")
	}
	example()
}
func example() {
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // file реализует io.Closer
}
