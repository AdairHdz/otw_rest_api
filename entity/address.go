package entity

type Address struct {
	EntityUUID
	CityID string
	City City
	IndoorNumber string
	OutdoorNumber string
	Street string
	Suburb string
	ServiceRequesterID string
}