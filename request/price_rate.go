package request

import (
	"time"

	"github.com/AdairHdz/OTW-Rest-API/entity"
	uuid "github.com/satori/go.uuid"
)

type PriceRate struct {
	StartingHour  string  `json:"startingHour"`
	EndingHour    string  `json:"endingHour"`
	Price         float64 `json:"price"`
	KindOfService int     `json:"kindOfService"`
	CityID        string  `json:"cityId"`
	WorkingDays   []int   `json:"workingDays"`
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
	}

	for _, w := range p.WorkingDays {
		pr.WorkingDays = append(pr.WorkingDays, entity.WorkingDay{
			ID: w,
		})
	}

	return
}