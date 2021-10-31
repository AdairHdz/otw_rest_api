package entity

const (
	SERVICE_PAYMENT = iota
	DRUG_PURCHASE
	GROCERY_SHOPPING
	DELIVERY
	OTHER
)

const (
	PENDING_OF_ACCEPTANCE = iota
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