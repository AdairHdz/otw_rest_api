package entity

type PriceRate struct {
	EntityUUID
	CityID string
	City City
	EndingHour string
	KindOfService int
	Price        float64
	StartingHour string
	// WorkingDays  []int
}