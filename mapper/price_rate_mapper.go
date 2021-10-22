package mapper

import (
	"github.com/AdairHdz/OTW-Rest-API/entity"
)

type PriceRateWorkingDays struct{
	ID 				string			`json:"id"`
	StartingHour 	string			`json:"startingHour"`
	EndingHour 		string			`json:"endingHour"`
	Price        	float64			`json:"price"`
	KindOfService 	int				`json:"kindOfService"`
	City  			City			`json:"city"`
	WorkingDays  	[]int	`json:"workingDays"`
}

func NewPriceRates() PriceRateWorkingDays{
	return PriceRateWorkingDays{
		WorkingDays: make([]int, 0, 7),
	}
}

type WorkingDay struct{
	ID 			int 		`json:"id"`
}

type City struct{
	ID 			string		`json:"id"`
	Name		string 		`json:"name"`
}

func CreatePriceRateWorkingDaysAsResponse(priceRate entity.PriceRate) PriceRateWorkingDays {
	r := NewPriceRates()

	r.ID = priceRate.ID
	r.StartingHour = priceRate.StartingHour
	r.EndingHour = priceRate.EndingHour
	r.Price = priceRate.Price
	r.KindOfService = priceRate.KindOfService 		
	
	for _, WorkingDayItem := range priceRate.WorkingDays {
		workingDay := WorkingDay{
			ID: WorkingDayItem.ID,
		}
		r.WorkingDays = append(r.WorkingDays, workingDay.ID)
	}

	city := City{
		ID: priceRate.City.ID,
		Name: priceRate.City.Name,
	}  
	r.City = city

	return r
}

