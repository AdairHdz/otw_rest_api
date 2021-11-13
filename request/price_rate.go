package request

import (
	"time"

	"github.com/AdairHdz/OTW-Rest-API/entity"
	uuid "github.com/satori/go.uuid"
)

type PriceRate struct {
	StartingHour  string  `validate:"required" json:"startingHour"`
	EndingHour    string  `validate:"required" json:"endingHour"`
	Price         float64 `validate:"numeric,gt=0" json:"price"`
	KindOfService int     `validate:"oneof=1 2 3 4 5" json:"kindOfService"`
	CityID        string  `validate:"required,uuid4" json:"cityId"`
	WorkingDays   []int   `validate:"required"json:"workingDays"`
}

func (p *PriceRate) ToEntity(serviceProviderID string) (pr entity.PriceRate, err error) {
	_, err = time.Parse("15:04", p.StartingHour)
	if err != nil {		
		return
	}

	_, err = time.Parse("15:04", p.EndingHour)
	if err != nil {		
		return
	}

	pr = entity.PriceRate {
		EntityUUID: entity.EntityUUID{
			ID: uuid.NewV4().String(),			
		},
		CityID: p.CityID,
		EndingHour: p.EndingHour,
		StartingHour: p.StartingHour,
		ServiceProviderID: serviceProviderID,
		WorkingDays: []entity.WorkingDay{},
		KindOfService: p.KindOfService,
		Price: p.Price,
	}

	for _, w := range p.WorkingDays {
		pr.WorkingDays = append(pr.WorkingDays, entity.WorkingDay{
			ID: w,
		})
	}

	return
}