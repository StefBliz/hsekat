package domain

import (
	"errors"
	"math"
)

type Helmet struct {
	baseItem
}

func NewHelmet(id, name string) (*Helmet, error) {
	base, err := newBaseItem(id, name)
	if err != nil {
		return nil, err
	}

	return &Helmet{
		baseItem: base,
	}, nil
}

// зарядная станция
type ChargingStation struct {
	baseItem
	energyUsage
}

func NewChargingStation(id, name string, energy float64) (*ChargingStation, error) {
	base, err := newBaseItem(id, name)
	if err != nil {
		return nil, err
	}

	if energy < 0 || math.IsNaN(energy) || math.IsInf(energy, 0) {
		return nil, errors.New("некорректное энергопотребление")
	}

	return &ChargingStation{
		baseItem:    base,
		energyUsage: energyUsage{dailyEnergy: energy},
	}, nil
}
