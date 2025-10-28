package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Создаем канал для получения сигналов
	sigChan := make(chan os.Signal, 1)

	// Настраиваем обработку сигналов SIGINT (Ctrl+C) и SIGTERM (сигнал завершения)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Ожидаем получения сигнала
	fmt.Println("Приложение запущено. Для завершения нажмите Ctrl+C или отправьте SIGTERM.")
	<-sigChan

	fmt.Println("Получен сигнал завершения. Завершаем приложение корректно.")

	// Ваш код для завершения приложения
	// Например, закрытие файлов, соединений с базой данных и т. д.

	// Завершаем приложение с кодом 0 (успешное завершение)
	os.Exit(0)
}
