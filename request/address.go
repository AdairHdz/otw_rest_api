package request

import (
	"github.com/AdairHdz/OTW-Rest-API/entity"
	uuid "github.com/satori/go.uuid"
)

type Address struct {
	CityID string `validate:"required,uuid4" json:"cityId"`
	IndoorNumber string `validate:"omitempty,max=50,alphanum" json:"indoorNumber"`
	OutdoorNumber string `validate:"required,max=50,alphanum" json:"outdoorNumber"`
	Street string `validate:"required,max=100,alphanum" json:"street"`
	Suburb string `validate:"required,max=100,alphanum" json:"suburb"`
}

func (a *Address) ToEntity(serviceRequesterID string) (ad *entity.Address, err error) {
	ad = &entity.Address{
		EntityUUID: entity.EntityUUID{
			ID: uuid.NewV4().String(),
		},
		CityID: a.CityID,
		IndoorNumber: a.IndoorNumber,
		OutdoorNumber: a.OutdoorNumber,
		Street: a.Street,
		Suburb: a.Suburb,
		ServiceRequesterID: serviceRequesterID,
	}
	
	return
}
