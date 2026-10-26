package console

import (
	"bytes"
	"hsekat/internal/fleet"
	"hsekat/internal/inspection"
	"strings"
	"testing"
)

func newTestApp(input string) (*ConsoleApp, *bytes.Buffer) {
	checker, _ := inspection.NewStandardServiceCenter(10)
	park := fleet.NewFleetService(checker)
	output := &bytes.Buffer{}
	program := NewConsoleApp(park, strings.NewReader(input), output)
	return program, output
}

func TestConsoleAddTransport(t *testing.T) {
	input := "1\n1\nSC-001\nNinebot\n8\n1.5\n3\n0\n"
	program, output := newTestApp(input)
	if err := program.Run(); err != nil {
		t.Fatal(err)
	}

	result := output.String()
	if !strings.Contains(result, "Транспорт успешно принят") {
		t.Fatal("транспорт не был добавлен")
	}
	if !strings.Contains(result, "Транспорт: 1, вещи: 0") {
		t.Fatal("неверный подсчёт транспорта")
	}
}

func TestConsoleAddItem(t *testing.T) {
	input := "2\n1\nHE-001\nHelmet\n3\n0\n"
	program, output := newTestApp(input)
	if err := program.Run(); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output.String(), "Транспорт: 0, вещи: 1") {
		t.Fatal("вещь не была добавлена")
	}
}

func TestConsoleInvalidCommand(t *testing.T) {
	program, output := newTestApp("99\n0\n")
	if err := program.Run(); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output.String(), "Неизвестная команда") {
		t.Fatal("неверная обработка неизвестной команды")
	}
}

func TestConsoleDuplicateNumber(t *testing.T) {
	input := "2\n1\nHE-001\nHelmet\n2\n1\nHE-001\nAnother Helmet\n3\n0\n"
	program, output := newTestApp(input)
	if err := program.Run(); err != nil {
		t.Fatal(err)
	}

	result := output.String()
	if !strings.Contains(result, "уже используется") {
		t.Fatal("дубликат номера должен вызывать ошибку")
	}

	if !strings.Contains(result, "Транспорт: 0, вещи: 1") {
		t.Fatal("дубликат не должен увеличивать количество имущества")
	}
}
