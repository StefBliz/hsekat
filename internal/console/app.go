package console

import (
	"bufio"
	"fmt"
	"hsekat/internal/fleet"
	"io"
	"strconv"
	"strings"
)

type ConsoleApp struct {
	fleet *fleet.FleetService
	in    *bufio.Reader
	out   io.Writer
}

func NewConsoleApp(f *fleet.FleetService, input io.Reader, output io.Writer) *ConsoleApp {
	return &ConsoleApp{
		fleet: f,
		in:    bufio.NewReader(input),
		out:   output,
	}
}

func (c *ConsoleApp) readLine() (string, error) {
	line, err := c.in.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}

	if err == io.EOF && line == "" {
		return "", io.EOF
	}

	return strings.TrimSpace(line), nil
}

func (c *ConsoleApp) readInt(prompt string) (int, error) {
	fmt.Fprint(c.out, prompt)
	value, err := c.readLine()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(value)
}

func (c *ConsoleApp) showMenu() {
	fmt.Fprintln(c.out, "\n========== ВЫШКАТ ==========")
	fmt.Fprintln(c.out, "1. Добавить транспорт")
	fmt.Fprintln(c.out, "2. Добавить вещь")
	fmt.Fprintln(c.out, "3. Количество транспорта и вещей")
	fmt.Fprintln(c.out, "4. Суммарное энергопотребление")
	fmt.Fprintln(c.out, "5. Транспорт для новичков")
	fmt.Fprintln(c.out, "6. Полный список имущества")
	fmt.Fprintln(c.out, "0. Выход")
}

func (c *ConsoleApp) Run() error {
	for {
		c.showMenu()
		choice, err := c.readInt("Выберите действие: ")
		if err == io.EOF {
			return nil
		}
		if err != nil {
			fmt.Fprintln(c.out, "Введите корректный номер команды.")
			continue
		}

		switch choice {
		case 1:
			if err := c.addTransport(); err != nil {
				fmt.Fprintln(c.out, "Ошибка:", err)
			}
		case 2:
			if err := c.addItem(); err != nil {
				fmt.Fprintln(c.out, "Ошибка:", err)
			}
		case 3:
			transport, items := c.fleet.Count()
			fmt.Fprintf(c.out, "Транспорт: %d, вещи: %d\n", transport, items)
		case 4:
			fmt.Fprintf(c.out, "Энергопотребление: %.2f кВт·ч/сутки\n",
				c.fleet.TotalDailyEnergy())
		case 5:
			fmt.Fprintln(c.out, "Транспорт для новичков:")
			for _, t := range c.fleet.ForBeginners() {
				fmt.Fprintf(c.out, "%s — %s\n", t.InventoryNumber(), t.Name())
			}
		case 6:
			fmt.Fprintln(c.out, "Всё имущество:")
			for _, item := range c.fleet.AllItems() {
				fmt.Fprintf(c.out, "%s — %s\n", item.InventoryNumber(), item.Name())
			}
		case 0:
			fmt.Fprintln(c.out, "До свидания!")
			return nil
		default:
			fmt.Fprintln(c.out, "Неизвестная команда.")
		}
	}
}
