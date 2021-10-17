package entity

type Review struct {
	EntityUUID
	Title string
	DateOfReview string
	Details string
	//Evidence
	Score int
	ServiceProviderID string
	ServiceRequesterID string
}