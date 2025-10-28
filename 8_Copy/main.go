package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	sourceFile, err := os.Open("source.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create("destination.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer destFile.Close()

	n, err := io.Copy(destFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Copied %d bytes\n", n)
}
