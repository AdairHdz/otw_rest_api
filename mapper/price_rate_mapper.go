package mapper

import (
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/response"
)

func NewPriceRates() response.PriceRateWorkingDays{
	return response.PriceRateWorkingDays{
		WorkingDays: make([]int, 0, 7),
	}
}

func NewPriceRate() response.PriceRate{
	return response.PriceRate{
		WorkingDays: make([]int, 0, 7),
	}
}

func CreatePriceRateWorkingDaysAsResponse(priceRate entity.PriceRate) response.PriceRateWorkingDays {
	r := NewPriceRates()

	r.ID = priceRate.ID
	r.StartingHour = priceRate.StartingHour
	r.EndingHour = priceRate.EndingHour
	r.Price = priceRate.Price
	r.KindOfService = priceRate.KindOfService 		
	
	for _, WorkingDayItem := range priceRate.WorkingDays {
		workingDay := response.WorkingDay{
			ID: WorkingDayItem.ID,
		}
		r.WorkingDays = append(r.WorkingDays, workingDay.ID)
	}

	city := response.City{
		ID: priceRate.City.ID,
		Name: priceRate.City.Name,
	}  
	r.City = city

	return r
}

func CreatePriceRateAddAsResponse(priceRate *entity.PriceRate) response.PriceRate {
	r := NewPriceRate()
	r.ID = priceRate.ID
	r.StartingHour = priceRate.StartingHour
	r.EndingHour = priceRate.EndingHour
	r.Price = priceRate.Price
	r.KindOfService = priceRate.KindOfService 
	r.CityId = priceRate.CityID		
	
	for _, WorkingDayItem := range priceRate.WorkingDays {
		workingDay := response.WorkingDay{
			ID: WorkingDayItem.ID,
		}
		r.WorkingDays = append(r.WorkingDays, workingDay.ID)
	}
	return r
}
