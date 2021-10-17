package entity

type ServiceProvider struct {
	EntityUUID
	UserID string
	User User
	BusinessPicture string
	BusinessName string
	ServiceRequests []ServiceRequest
	Reviews []Review
}