package main

import (
	"fmt"
	"hsekat/internal/app"
	"hsekat/internal/console"
	"os"
)

func main() {
	container, err := app.NewContainer()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка запуска:", err)
		os.Exit(1)
	}
	program := console.NewConsoleApp(
		container.Fleet,
		os.Stdin,
		os.Stdout,
	)
	if err := program.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка приложения:", err)
		os.Exit(1)
	}
}
