package app

import (
	"hsekat/internal/fleet"
	"hsekat/internal/inspection"
)

// хранит основные компоненты приложения
type Container struct {
	Fleet *fleet.FleetService
}

// создаёт зависимости и связывает их
func NewContainer() (*Container, error) {
	checker, err := inspection.NewStandardServiceCenter(10.0)
	if err != nil {
		return nil, err
	}
	fleetService := fleet.NewFleetService(checker)

	return &Container{
		Fleet: fleetService,
	}, nil
}
