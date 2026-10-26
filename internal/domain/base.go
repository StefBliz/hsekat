package domain

import (
	"errors"
	"strings"
)

// общие хар-ки имущества
type baseItem struct {
	id   string
	name string
}

func newBaseItem(id, name string) (baseItem, error) {
	//избавляюсь от пробелов
	id = strings.TrimSpace(id)
	name = strings.TrimSpace(name)
	if id == "" {
		return baseItem{}, errors.New("инвентарный номер не должен быть пустым")
	}
	if name == "" {
		return baseItem{}, errors.New("название не должно быть пустым")
	}
	return baseItem{
		id:   id,
		name: name,
	}, nil
}

func (b baseItem) InventoryNumber() string {
	return b.id
}

func (b baseItem) Name() string {
	return b.name
}
