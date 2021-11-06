package request

type RequestStatus struct {
	ServiceStatus int `validate:"oneof=1 2 3 4 5" json:"serviceStatus"`
}
