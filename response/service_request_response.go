package response

type ServiceRequest struct{
	ID 					string							`json:"id"`
	Date 				string							`json:"date"`
	Cost 				float64							`json:"cost"`
	Description 		string							`json:"description"`
	HasBeenReviewed		bool							`json:"hasBeenReviewed"`
	KindOfService 		int								`json:"kindOfService"`
	Status 				int								`json:"status"`
	ServiceRequester 	ServiceRequesterInRequest		`json:"serviceRequester"`
	ServiceProvider 	ServiceProviderInRequest		`json:"serviceProvider"`
	DeliveryAddress		Address							`json:"deliveryAddress"`
}

type ServiceRequestWithCity struct{
	ID 					string							`json:"id"`
	Date 				string							`json:"date"`
	Cost 				float64							`json:"cost"`
	Description 		string							`json:"description"`
	HasBeenReviewed		bool							`json:"hasBeenReviewed"`
	KindOfService 		int								`json:"kindOfService"`
	Status 				int								`json:"status"`
	ServiceRequester 	ServiceRequesterInRequest		`json:"serviceRequester"`
	ServiceProvider 	ServiceProviderInRequest		`json:"serviceProvider"`
	DeliveryAddress		AddressWithCity					`json:"deliveryAddress"`
}