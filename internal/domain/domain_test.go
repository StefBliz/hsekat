package domain

import (
	"math"
	"testing"
)

func TestInvalidScooterEnergy(t *testing.T) {
	values := []float64{-1, math.NaN(), math.Inf(1)}

	for _, energy := range values {
		_, err := NewScooter("SC-001", "Ninebot", energy, 8)

		if err == nil {
			t.Fatalf("ожидали ошибку для энергопотребления %v", energy)
		}
	}
}

func TestInvalidElectricBikeEnergy(t *testing.T) {
	values := []float64{-1, math.NaN(), math.Inf(1)}

	for _, energy := range values {
		_, err := NewElectricBike("EB-001", "Electric Bike", energy, 7)

		if err == nil {
			t.Fatalf("ожидали ошибку для энергопотребления %v", energy)
		}
	}
}

func TestInvalidChargingStationEnergy(t *testing.T) {
	values := []float64{-1, math.NaN(), math.Inf(1)}

	for _, energy := range values {
		_, err := NewChargingStation("CH-001", "Station", energy)

		if err == nil {
			t.Fatalf("ожидали ошибку для энергопотребления %v", energy)
		}
	}
}

func TestInvalidEaseOfUse(t *testing.T) {
	_, err := NewBicycle("B-001", "Bike", 11)

	if err == nil {
		t.Fatal("простота управления больше 10 должна вызывать ошибку")
	}
}

func TestEmptyInventoryNumber(t *testing.T) {
	_, err := NewHelmet("", "Helmet")

	if err == nil {
		t.Fatal("пустой инвентарный номер должен вызывать ошибку")
	}
}

func TestScooterProperties(t *testing.T) {
	scooter, err := NewScooter("SC-001", "Ninebot", 1.5, 8)
	if err != nil {
		t.Fatal(err)
	}

	if scooter.InventoryNumber() != "SC-001" {
		t.Fatal("неверный инвентарный номер")
	}

	if scooter.Name() != "Ninebot" {
		t.Fatal("неверное название")
	}

	if scooter.EaseOfUse() != 8 {
		t.Fatal("неверная простота управления")
	}

	if scooter.DailyEnergyConsumption() != 1.5 {
		t.Fatal("неверное энергопотребление")
	}
}

func TestBicycleProperties(t *testing.T) {
	bicycle, err := NewBicycle("B-001", "City Bike", 7)
	if err != nil {
		t.Fatal(err)
	}

	if bicycle.InventoryNumber() != "B-001" {
		t.Fatal("неверный инвентарный номер")
	}

	if bicycle.Name() != "City Bike" {
		t.Fatal("неверное название")
	}

	if bicycle.EaseOfUse() != 7 {
		t.Fatal("неверная простота управления")
	}
}

func TestElectricBikeProperties(t *testing.T) {
	bike, err := NewElectricBike("EB-001", "Electric Bike", 2.5, 6)
	if err != nil {
		t.Fatal(err)
	}

	if bike.EaseOfUse() != 6 {
		t.Fatal("неверная простота управления")
	}

	if bike.DailyEnergyConsumption() != 2.5 {
		t.Fatal("неверное энергопотребление")
	}
}

func TestHelmetProperties(t *testing.T) {
	helmet, err := NewHelmet("HE-001", "Helmet")
	if err != nil {
		t.Fatal(err)
	}

	if helmet.InventoryNumber() != "HE-001" {
		t.Fatal("неверный инвентарный номер")
	}

	if helmet.Name() != "Helmet" {
		t.Fatal("неверное название")
	}
}
