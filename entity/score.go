package entity

type Score struct {
	EntityUUID
	UserID string
	AverageScore float64
	MaxTotalPossible int
	ObtainedPoints int
}