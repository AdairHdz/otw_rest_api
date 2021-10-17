package entity

type User struct {
	EntityUUID
	StateID string
	State State
	Account Account
	Names string
	Lastname string
	Scores []Score
}