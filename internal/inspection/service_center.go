package inspection

import (
	"errors"
	"math"

	"hsekat/internal/domain"
)

// ServiceCenter задаёт контракт техосмотра
type ServiceCenter interface {
	Inspect(t domain.Transport) error
}

// StandardServiceCenter выполняет проверку транспорта
type StandardServiceCenter struct {
	maxEnergy float64
}

// NewStandardServiceCenter создаёт сервис техосмотра

func NewStandardServiceCenter(maxEnergy float64) (*StandardServiceCenter, error) {
	if maxEnergy < 0 || math.IsNaN(maxEnergy) || math.IsInf(maxEnergy, 0) {
		return nil, errors.New("некорректный лимит энергопотребления")
	}

	return &StandardServiceCenter{
		maxEnergy: maxEnergy,
	}, nil
}

// Inspect проверяет транспорт перед приёмкой
func (s *StandardServiceCenter) Inspect(t domain.Transport) error {
	if t == nil {
		return errors.New("транспорт не указан")
	}

	if t.EaseOfUse() < 1 || t.EaseOfUse() > 10 {
		return errors.New("некорректная простота управления")
	}

	if electric, ok := t.(domain.EnergyConsumer); ok {
		if electric.DailyEnergyConsumption() > s.maxEnergy {
			return errors.New("превышено допустимое энергопотребление")
		}
	}

	return nil
}
