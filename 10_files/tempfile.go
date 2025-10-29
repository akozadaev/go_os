// tempfile.go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Создаём временный файл в системной временной директории (dir = "")
	tempFile, err := os.CreateTemp(".", "log-*.txt")
	if err != nil {
		fmt.Printf("Ошибка создания временного файла: %v\n", err)
		return
	}
	// Гарантированно удаляем файл при выходе
	defer os.Remove(tempFile.Name())

	fmt.Printf("Создан временный файл: %s\n", tempFile.Name())

	// Записываем данные
	data := "Это временный лог-файл."
	_, err = tempFile.Write([]byte(data))
	if err != nil {
		fmt.Printf("Ошибка записи в файл: %v\n", err)
		return
	}

	// Закрываем файл перед чтением (важно!)
	defer tempFile.Close()

	// Читаем и выводим содержимое
	content, err := os.ReadFile(tempFile.Name())
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
		return
	}

	fmt.Printf("Содержимое файла: %s\n", string(content))
}
