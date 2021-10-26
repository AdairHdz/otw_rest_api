package response 

type ReviewWithEvidence struct{
	ID 					string		`json:"id"`
	Title 				string 		`json:"title"`
	DateOfReview 		string 		`json:"dateOfReview"`
	Details 			string 		`json:"details"`
	Score 				int 		`json:"score"`
	ServiceRequester 	string 		`json:"requesterName"`
	Evidences 			[]Evidence	`json:"evidence"`
}