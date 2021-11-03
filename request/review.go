package request

import (
	"time"

	"github.com/AdairHdz/OTW-Rest-API/entity"
	uuid "github.com/satori/go.uuid"
)

type Review struct {
	Title              string           `json:"title" validate:"required,max=50"`
	Details            string           `json:"details" validate:"omitempty,max=150"`
	Score              int              `json:"score" validate:"min=1,max=5"`
	Evidence           []ReviewEvidence `json:"evidence"`
	ServiceRequesterID string           `json:"serviceRequesterId" validate:"required,uuid4"`
}

func (r Review) ToEntity(serviceProviderID string) (entityReview entity.Review) {
	t := time.Now()
	formattedTime := t.Format("01-02-2006")	

	entityReview = entity.Review{
		EntityUUID: entity.EntityUUID{
			ID: uuid.NewV4().String(),
		},
		Title: r.Title,
		DateOfReview: formattedTime,
		Details: r.Details,
		Score: r.Score,
		ServiceRequesterID: r.ServiceRequesterID,
		ServiceProviderID: serviceProviderID,		
	}
	return
}