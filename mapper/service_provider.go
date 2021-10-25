package mapper

import (
	"github.com/AdairHdz/OTW-Rest-API/entity"
)

type ServiceProvider struct{
	ID 				string		`json:"id"`
	Names 			string 		`json:"names"`
	Lastname 		string 		`json:"lastNames"`
	AverageScore 	float64 	`json:"averageScore"`
	PriceRate 		float64 	`json:"priceRate"`
}

func CreateServiceProvidersAsResponse(price_rate entity.PriceRate) ServiceProvider {
	r := ServiceProvider {
		ID: price_rate.ServiceProvider.ID,
		Names: price_rate.ServiceProvider.User.Names,
		Lastname: price_rate.ServiceProvider.User.Lastname,
		AverageScore: price_rate.ServiceProvider.User.Score.AverageScore,
		PriceRate: price_rate.Price,
	}
	return r
}

