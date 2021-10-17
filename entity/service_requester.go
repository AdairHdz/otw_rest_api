package entity

type ServiceRequester struct {
	EntityUUID
	UserID string
	User User
	ServiceRequests []ServiceRequest
	Reviews []Review
}