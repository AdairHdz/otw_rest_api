package mapper

import (
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/response"
)

func CreateServiceProvidersAsResponse(price_rate entity.PriceRate) response.ServiceProvider {
	r := response.ServiceProvider {
		ID: price_rate.ServiceProvider.ID,
		Names: price_rate.ServiceProvider.User.Names,
		Lastname: price_rate.ServiceProvider.User.Lastname,
		AverageScore: price_rate.ServiceProvider.User.Score.AverageScore,
		PriceRate: price_rate.Price,
		BusinessName: price_rate.ServiceProvider.BusinessName,
	}
	return r
}

