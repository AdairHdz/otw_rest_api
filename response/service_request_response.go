package response

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


type ServiceRequestAdd struct{
	ID 					string							`json:"id"`
	Cost 				float64							`json:"cost"`
	DeliveryAddressId	string 							`json:"deliveryAddressId"`
	Description 		string							`json:"description"`
	KindOfService 		int								`json:"kindOfService"`
	ServiceProviderId 	string 							`json:"serviceProviderId"`
	ServiceRequesterId 	string							`json:"serviceRequesterId"`
}

type ServiceRequestDetails struct{
	ID 					string							`json:"id"`
	Date 				string							`json:"date"`
	KindOfService 		int								`json:"kindOfService"`
	Status 				int								`json:"status"`
}