package fleet

import (
	"fmt"
	"hsekat/internal/domain"
	"testing"
)

// Подменный техосмотр для тестов
type fakeChecker struct {
	approved bool
}

func (f fakeChecker) Inspect(t domain.Transport) error {
	if !f.approved {
		return fmt.Errorf("техосмотр не пройден")
	}
	return nil
}

func TestAddTransport(t *testing.T) {
	park := NewFleetService(fakeChecker{approved: true})

	scooter, err := domain.NewScooter("SC-001", "Ninebot", 1.5, 8)
	if err != nil {
		t.Fatal(err)
	}

	if err := park.AddTransport(scooter); err != nil {
		t.Fatal(err)
	}

	transportCount, itemCount := park.Count()

	if transportCount != 1 || itemCount != 0 {
		t.Fatalf("ожидали 1 транспорт и 0 вещей, получили %d и %d",
			transportCount, itemCount)
	}
}

func TestDuplicateNumber(t *testing.T) {
	park := NewFleetService(fakeChecker{approved: true})

	first, _ := domain.NewBicycle("B-001", "Bike", 7)
	second, _ := domain.NewHelmet("B-001", "Helmet")

	if err := park.AddTransport(first); err != nil {
		t.Fatal(err)
	}

	if err := park.AddItem(second); err == nil {
		t.Fatal("ожидали ошибку при повторном номере")
	}
}

func TestInspectionRejected(t *testing.T) {
	park := NewFleetService(fakeChecker{approved: false})

	scooter, _ := domain.NewScooter("SC-002", "Ninebot", 2, 8)

	if err := park.AddTransport(scooter); err == nil {
		t.Fatal("транспорт должен был не пройти техосмотр")
	}

	count, _ := park.Count()
	if count != 0 {
		t.Fatal("транспорт не должен попадать в парк")
	}
}

func TestTotalEnergy(t *testing.T) {
	park := NewFleetService(fakeChecker{approved: true})

	scooter, _ := domain.NewScooter("SC-001", "Scooter", 1.5, 8)
	station, _ := domain.NewChargingStation("CH-001", "Station", 2.5)

	if err := park.AddTransport(scooter); err != nil {
		t.Fatal(err)
	}

	if err := park.AddItem(station); err != nil {
		t.Fatal(err)
	}

	if got := park.TotalDailyEnergy(); got != 4.0 {
		t.Fatalf("ожидали 4.0, получили %f", got)
	}
}

func TestForBeginners(t *testing.T) {
	park := NewFleetService(fakeChecker{approved: true})

	easy, _ := domain.NewBicycle("B-001", "Easy", 8)
	hard, _ := domain.NewBicycle("B-002", "Hard", 3)

	if err := park.AddTransport(easy); err != nil {
		t.Fatal(err)
	}

	if err := park.AddTransport(hard); err != nil {
		t.Fatal(err)
	}

	result := park.ForBeginners()

	if len(result) != 1 || result[0].InventoryNumber() != "B-001" {
		t.Fatal("неверный список транспорта для новичков")
	}
}

func TestAllItems(t *testing.T) {
	park := NewFleetService(fakeChecker{approved: true})

	scooter, err := domain.NewScooter("SC-001", "Ninebot", 1.5, 8)
	if err != nil {
		t.Fatal(err)
	}

	helmet, err := domain.NewHelmet("HE-001", "Helmet")
	if err != nil {
		t.Fatal(err)
	}

	if err := park.AddTransport(scooter); err != nil {
		t.Fatal(err)
	}

	if err := park.AddItem(helmet); err != nil {
		t.Fatal(err)
	}

	items := park.AllItems()

	if len(items) != 2 {
		t.Fatalf("ожидали 2 объекта, получили %d", len(items))
	}

	if items[0].InventoryNumber() != "HE-001" {
		t.Fatal("неверная сортировка имущества")
	}

	if items[1].InventoryNumber() != "SC-001" {
		t.Fatal("неверная сортировка имущества")
	}
}

func TestTransportCannotBeAddedAsItem(t *testing.T) {
	park := NewFleetService(fakeChecker{approved: true})

	scooter, err := domain.NewScooter("SC-001", "Ninebot", 1.5, 8)
	if err != nil {
		t.Fatal(err)
	}

	if err := park.AddItem(scooter); err == nil {
		t.Fatal("транспорт нельзя добавлять без техосмотра")
	}

	transportCount, itemCount := park.Count()

	if transportCount != 0 || itemCount != 0 {
		t.Fatal("отклонённый транспорт не должен попадать на баланс")
	}
}
