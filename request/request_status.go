package request

type RequestStatus struct {
	ServiceStatus int `validate:"oneof=0 1 2 3 4" json:"serviceStatus"`
}
