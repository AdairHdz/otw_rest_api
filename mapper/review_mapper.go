package mapper

import (
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/response"
)

func NewReview() response.ReviewWithEvidence{
	return response.ReviewWithEvidence{
		Evidences: make([]response.Evidence, 0, 10),
	}
}

func CreateReviewWithEvidenceAsResponse(review entity.Review) response.ReviewWithEvidence {
	r := NewReview()

	r.ID = review.ID
	r.Title = review.Title
	r.DateOfReview = review.DateOfReview
	r.Details = review.Details
	r.Score = review.Score
	r.ServiceRequester = review.ServiceRequester.User.Names + " " + review.ServiceRequester.User.Lastname
	
	for _, evidenceItem := range review.Evidences {
		evidence := response.Evidence{
			ID: evidenceItem.ID,
			FileName: evidenceItem.FileName,
			FileExtension: evidenceItem.FileExtension,
		}
		r.Evidences = append(r.Evidences, evidence)
	}

	return r
}

func CreateReviewWithRequesterIDAsResponse(review entity.Review) (response.ReviewWithRequesterID) {
	r := response.ReviewWithRequesterID{
		ID: review.ID,
		Title: review.Title,
		DateOfReview: review.DateOfReview,
		Details: review.Details,
		Score: review.Score,
		ServiceRequesterID: review.ServiceRequesterID,
		Evidences: make([]response.Evidence, 0),
	}

	for _, evidenceItem := range review.Evidences {
		evidence := response.Evidence{
			ID: evidenceItem.ID,
			FileName: evidenceItem.FileName,
			FileExtension: evidenceItem.FileExtension,
		}
		r.Evidences = append(r.Evidences, evidence)
	}

	return r

}

