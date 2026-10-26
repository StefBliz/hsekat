package fleet

import (
	"errors"
	"fmt"
	"sort"

	"hsekat/internal/domain"
	"hsekat/internal/inspection"
)

type FleetService struct {
	items      map[string]domain.InventoryItem
	checker    inspection.ServiceCenter
	transports map[string]domain.Transport
}

func NewFleetService(checker inspection.ServiceCenter) *FleetService {
	return &FleetService{
		items:      make(map[string]domain.InventoryItem),
		transports: make(map[string]domain.Transport),
		checker:    checker,
	}
}

// принимает транспорт после техосмотра
func (f *FleetService) AddTransport(t domain.Transport) error {
	if t == nil {
		return errors.New("транспорт не указан")
	}

	id := t.InventoryNumber()
	if id == "" {
		return errors.New("пустой инвентарный номер")
	}

	if _, exists := f.items[id]; exists {
		return fmt.Errorf("номер %s уже используется", id)
	}

	if f.checker == nil {
		return errors.New("сервис техосмотра не настроен")
	}

	if err := f.checker.Inspect(t); err != nil {
		return fmt.Errorf("техосмотр не пройден: %w", err)
	}

	f.items[id] = t
	f.transports[id] = t
	return nil
}

// добавляет вещь на баланс
func (f *FleetService) AddItem(item domain.InventoryItem) error {
	if item == nil {
		return errors.New("вещь не указана")
	}

	if _, isTransport := item.(domain.Transport); isTransport {
		return errors.New("транспорт необходимо добавлять через AddTransport")
	}

	id := item.InventoryNumber()
	if id == "" {
		return errors.New("пустой инвентарный номер")
	}

	if _, exists := f.items[id]; exists {
		return fmt.Errorf("номер %s уже используется", id)
	}

	f.items[id] = item
	return nil
}

// Count число транспорта и вещей
func (f *FleetService) Count() (int, int) {
	transportCount := len(f.transports)
	itemCount := len(f.items) - transportCount
	return transportCount, itemCount
}

// TotalDailyEnergy считает суточное энергопотребление
func (f *FleetService) TotalDailyEnergy() float64 {
	total := 0.0

	for _, item := range f.items {
		if consumer, ok := item.(domain.EnergyConsumer); ok {
			total += consumer.DailyEnergyConsumption()
		}
	}

	return total
}

// возвращает транспорт для новичков
func (f *FleetService) ForBeginners() []domain.Transport {
	result := make([]domain.Transport, 0)

	for _, transport := range f.transports {
		if transport.EaseOfUse() >= 6 {
			result = append(result, transport)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].InventoryNumber() < result[j].InventoryNumber()
	})

	return result
}

// возвращает всё имущество по номерам
func (f *FleetService) AllItems() []domain.InventoryItem {
	result := make([]domain.InventoryItem, 0, len(f.items))

	for _, item := range f.items {
		result = append(result, item)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].InventoryNumber() < result[j].InventoryNumber()
	})

	return result
}
