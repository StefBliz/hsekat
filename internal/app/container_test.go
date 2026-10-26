package app

import (
	"testing"

	"hsekat/internal/domain"
)

func TestNewContainer(t *testing.T) {
	container, err := NewContainer()
	if err != nil {
		t.Fatal(err)
	}

	if container == nil || container.Fleet == nil {
		t.Fatal("контейнер должен создавать FleetService")
	}

	scooter, err := domain.NewScooter("SC-001", "Ninebot", 2, 8)
	if err != nil {
		t.Fatal(err)
	}

	if err := container.Fleet.AddTransport(scooter); err != nil {
		t.Fatalf("транспорт должен пройти техосмотр: %v", err)
	}

	count, _ := container.Fleet.Count()
	if count != 1 {
		t.Fatalf("ожидали 1 транспорт, получили %d", count)
	}
}

func TestContainerRejectsHighEnergy(t *testing.T) {
	container, err := NewContainer()
	if err != nil {
		t.Fatal(err)
	}

	scooter, err := domain.NewScooter("SC-002", "Ninebot", 15, 8)
	if err != nil {
		t.Fatal(err)
	}

	if err := container.Fleet.AddTransport(scooter); err == nil {
		t.Fatal("контейнер должен подключить техосмотр с лимитом энергии")
	}

	count, _ := container.Fleet.Count()
	if count != 0 {
		t.Fatal("отклонённый транспорт не должен попадать в парк")
	}
}
