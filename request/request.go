package request

import (
	"github.com/AdairHdz/OTW-Rest-API/entity"
	uuid "github.com/satori/go.uuid"
)

type Request struct {
	Cost float64 `validate:"required,numeric" json:"cost"` 
	DeliveryAddressID string `validate:"required,uuid4" json:"deliveryAddressId"`
	Description string `validate:"required,max=255,alphanum" json:"description"`
	KindOfService int `validate:"oneof=1 2 3 4 5" json:"kindOfService"`
	ServiceRequesterID string `validate:"required,uuid4" json:"serviceRequesterId" `
	ServiceProviderID string  `validate:"required,uuid4" json:"serviceProviderId"`
}

func (r *Request) ToEntity() (sr *entity.ServiceRequest, err error) {
	sr = &entity.ServiceRequest{
		EntityUUID: entity.EntityUUID{
			ID: uuid.NewV4().String(),
		},
		Cost: r.Cost,
		DeliveryAddressID: r.DeliveryAddressID,
		Description: r.Description,
		KindOfService: r.KindOfService,
		ServiceRequesterID: r.ServiceRequesterID,
		ServiceProviderID: r.ServiceProviderID,
	}
	
	return
}