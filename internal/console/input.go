package console

import (
	"fmt"
	"strconv"

	"hsekat/internal/domain"
)

// считывает дробное число
func (c *ConsoleApp) readFloat(prompt string) (float64, error) {
	fmt.Fprint(c.out, prompt)
	value, err := c.readLine()
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(value, 64)
}

// считывает строку
func (c *ConsoleApp) readText(prompt string) (string, error) {
	fmt.Fprint(c.out, prompt)
	return c.readLine()
}

// добавляет транспорт через FleetService
func (c *ConsoleApp) addTransport() error {
	fmt.Fprintln(c.out, "\nВид транспорта:")
	fmt.Fprintln(c.out, "1. Электросамокат")
	fmt.Fprintln(c.out, "2. Велосипед")
	fmt.Fprintln(c.out, "3. Электровелосипед")
	kind, err := c.readInt("Выберите вид: ")
	if err != nil {
		return err
	}
	if kind < 1 || kind > 3 {
		return fmt.Errorf("неизвестный вид транспорта")
	}

	id, err := c.readText("Инвентарный номер: ")
	if err != nil {
		return err
	}
	name, err := c.readText("Название: ")
	if err != nil {
		return err
	}
	ease, err := c.readInt("Простота управления (1-10): ")
	if err != nil {
		return err
	}
	var transport domain.Transport

	switch kind {
	case 1, 3:
		energy, err := c.readFloat("Энергопотребление (кВт·ч/сутки): ")
		if err != nil {
			return err
		}
		if kind == 1 {
			transport, err = domain.NewScooter(id, name, energy, ease)
		} else {
			transport, err = domain.NewElectricBike(id, name, energy, ease)
		}

		if err != nil {
			return err
		}

	case 2:
		transport, err = domain.NewBicycle(id, name, ease)
		if err != nil {
			return err
		}
	}
	if err := c.fleet.AddTransport(transport); err != nil {
		return err
	}
	fmt.Fprintln(c.out, "Транспорт успешно принят в парк")
	return nil
}

// добавляет вещь на баланс
func (c *ConsoleApp) addItem() error {
	fmt.Fprintln(c.out, "\nВид имущества:")
	fmt.Fprintln(c.out, "1. Шлем")
	fmt.Fprintln(c.out, "2. Зарядная станция")

	kind, err := c.readInt("Выберите вид: ")
	if err != nil {
		return err
	}
	if kind != 1 && kind != 2 {
		return fmt.Errorf("неизвестный вид имущества")
	}
	id, err := c.readText("Инвентарный номер: ")
	if err != nil {
		return err
	}
	name, err := c.readText("Название: ")
	if err != nil {
		return err
	}
	var item domain.InventoryItem
	switch kind {
	case 1:
		item, err = domain.NewHelmet(id, name)

	case 2:
		energy, readErr := c.readFloat("Энергопотребление (кВт·ч/сутки): ")
		if readErr != nil {
			return readErr
		}
		item, err = domain.NewChargingStation(id, name, energy)
	}
	if err != nil {
		return err
	}
	if err := c.fleet.AddItem(item); err != nil {
		return err
	}
	fmt.Fprintln(c.out, "Имущество успешно добавлено")
	return nil
}
