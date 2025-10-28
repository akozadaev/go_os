package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Задайте команду, которую вы хотите выполнить
	cmd := exec.Command("ls", "-l")
	wd, _ := os.Getwd()
	fmt.Println("Текущая рабочая директория:", wd)
	
	cmd.Dir = "/home/akozadaev/projects/akozadaev/go/ibs/go_os"
	// Выполните команду
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Ошибка выполнения команды:", err)
		return
	}

	// Выведите результат выполнения команды
	fmt.Println(string(output))
}
