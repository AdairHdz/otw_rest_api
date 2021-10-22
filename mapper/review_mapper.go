package mapper

import (
	"github.com/AdairHdz/OTW-Rest-API/entity"
)

type ReviewWithEvidence struct{
	ID 					string		`json:"id"`
	Title 				string 		`json:"title"`
	DateOfReview 		string 		`json:"dateOfReview"`
	Details 			string 		`json:"details"`
	Score 				int 		`json:"score"`
	ServiceRequester 	string 		`json:"requesterName"`
	Evidences 			[]Evidence	`json:"evidence"`
}

func NewReview() ReviewWithEvidence{
	return ReviewWithEvidence{
		Evidences: make([]Evidence, 0, 10),
	}
}

type Evidence struct{
	ID 				string 		`json:"id"`
	FileName 		string 		`json:"fileName"`
	FileExtension 	string 		`json:"fileExtension"`
}

func CreateReviewWithEvidenceAsResponse(review entity.Review) ReviewWithEvidence {
	r := NewReview()

	r.ID = review.ID
	r.Title = review.Title
	r.DateOfReview = review.DateOfReview
	r.Details = review.Details
	r.Score = review.Score
	r.ServiceRequester = review.ServiceRequester.User.Names + " " + review.ServiceRequester.User.Lastname
	
	for _, evidenceItem := range review.Evidences {
		evidence := Evidence{
			ID: evidenceItem.ID,
			FileName: evidenceItem.FileName,
			FileExtension: evidenceItem.FileExtension,
		}
		r.Evidences = append(r.Evidences, evidence)
	}

	return r
}

