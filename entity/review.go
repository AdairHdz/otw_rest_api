package entity

type Review struct {
	EntityUUID
	Title string
	DateOfReview string
	Details string
	Evidences []Evidence
	Score int
	ServiceProviderID string
	ServiceRequesterID string
	ServiceRequester ServiceRequester
	ServiceRequestID string
	ServiceRequest ServiceRequest
}