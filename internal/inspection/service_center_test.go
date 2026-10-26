package inspection

import (
	"testing"

	"hsekat/internal/domain"
)

func TestInspectionSuccess(t *testing.T) {
	checker, err := NewStandardServiceCenter(10)
	if err != nil {
		t.Fatal(err)
	}

	scooter, err := domain.NewScooter("SC-001", "Ninebot", 2, 8)
	if err != nil {
		t.Fatal(err)
	}

	if err := checker.Inspect(scooter); err != nil {
		t.Fatalf("транспорт должен пройти техосмотр: %v", err)
	}
}

func TestInspectionHighEnergy(t *testing.T) {
	checker, err := NewStandardServiceCenter(10)
	if err != nil {
		t.Fatal(err)
	}

	scooter, err := domain.NewScooter("SC-002", "Ninebot", 15, 8)
	if err != nil {
		t.Fatal(err)
	}

	if err := checker.Inspect(scooter); err == nil {
		t.Fatal("транспорт с высоким энергопотреблением должен быть отклонён")
	}
}

func TestInvalidInspectionLimit(t *testing.T) {
	_, err := NewStandardServiceCenter(-1)

	if err == nil {
		t.Fatal("отрицательный лимит должен вызывать ошибку")
	}
}
