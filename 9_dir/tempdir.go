// tempdir.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Создаём временную директорию
	tempDir, err := os.MkdirTemp(".", "app-cache-*")
	if err != nil {
		fmt.Printf("Ошибка создания временной директории: %v\n", err)
		return
	}
	// Удаляем всю директорию при завершении
	//defer os.RemoveAll(tempDir)

	fmt.Printf("Создана временная директория: %s\n", tempDir)

	// Путь к файлу внутри временной директории
	jsonPath := filepath.Join(tempDir, "data.json")

	// Записываем JSON-данные
	jsonData := `{"status": "ok"}`
	err = os.WriteFile(jsonPath, []byte(jsonData), 0644)
	if err != nil {
		fmt.Printf("Ошибка записи файла: %v\n", err)
		return
	}

	// Читаем и выводим содержимое
	content, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
		return
	}

	fmt.Printf("Содержимое data.json: %s\n", string(content))
}
