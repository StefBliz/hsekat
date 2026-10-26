package domain

// суточное энергопотребление устройства
type energyUsage struct {
	dailyEnergy float64
}

func (e energyUsage) DailyEnergyConsumption() float64 {
	return e.dailyEnergy
}
