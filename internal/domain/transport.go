package domain

import (
	"errors"
	"math"
)

type Scooter struct {
	baseItem
	energyUsage
	easeOfUse int
}

func NewScooter(id, name string, energy float64, ease int) (*Scooter, error) {
	base, err := newBaseItem(id, name)
	if err != nil {
		return nil, err
	}
	if energy < 0 || math.IsNaN(energy) || math.IsInf(energy, 0) {
		return nil, errors.New("некорректное энергопотребление")
	}
	if ease < 1 || ease > 10 {
		return nil, errors.New("простота управления должна быть от 1 до 10")
	}

	return &Scooter{
		baseItem:    base,
		energyUsage: energyUsage{dailyEnergy: energy},
		easeOfUse:   ease,
	}, nil
}

func (s *Scooter) EaseOfUse() int {
	return s.easeOfUse
}

type Bicycle struct {
	baseItem
	easeOfUse int
}

func NewBicycle(id, name string, ease int) (*Bicycle, error) {
	base, err := newBaseItem(id, name)
	if err != nil {
		return nil, err
	}

	if ease < 1 || ease > 10 {
		return nil, errors.New("простота управления должна быть от 1 до 10")
	}

	return &Bicycle{
		baseItem:  base,
		easeOfUse: ease,
	}, nil
}

func (b *Bicycle) EaseOfUse() int {
	return b.easeOfUse
}

type ElectricBike struct {
	baseItem
	energyUsage
	easeOfUse int
}

func NewElectricBike(id, name string, energy float64, ease int) (*ElectricBike, error) {
	base, err := newBaseItem(id, name)
	if err != nil {
		return nil, err
	}

	if energy < 0 || math.IsNaN(energy) || math.IsInf(energy, 0) {
		return nil, errors.New("некорректное энергопотребление")
	}
	if ease < 1 || ease > 10 {
		return nil, errors.New("простота управления должна быть от 1 до 10")
	}

	return &ElectricBike{
		baseItem:    base,
		energyUsage: energyUsage{dailyEnergy: energy},
		easeOfUse:   ease,
	}, nil
}

func (e *ElectricBike) EaseOfUse() int {
	return e.easeOfUse
}
