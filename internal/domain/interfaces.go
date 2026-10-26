package domain

// контракт для объектов, которые можно поставить на баланс
type InventoryItem interface {
	InventoryNumber() string //номер инвентарный
	Name() string
}

// описывает объекты, кот потребляют электричество
type EnergyConsumer interface {
	DailyEnergyConsumption() float64 //суточный расход энергии
}

type Transport interface {
	InventoryItem
	EaseOfUse() int
}
