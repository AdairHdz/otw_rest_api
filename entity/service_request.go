package entity

const (
	_  = iota
	SERVICE_PAYMENT
	DRUG_PURCHASE
	GROCERY_SHOPPING
	DELIVERY
	OTHER
)

const (
	_ = iota
	PENDING_OF_ACCEPTANCE
	ACTIVE
	CONCLUDED
	CANCELED		
	REJECTED
)

type ServiceRequest struct {
	EntityUUID
	Cost float64
	Date string
	DeliveryAddress Address
	DeliveryAddressID string
	Description string
	HasBeenReviewed bool
	KindOfService int
	Status int
	ServiceRequesterID string
	ServiceRequester ServiceRequester
	ServiceProviderID string
	ServiceProvider ServiceProvider
}