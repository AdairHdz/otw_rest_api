package mapper

import (
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/response"
)

func CreateRequestsAsResponse(request entity.ServiceRequest) response.ServiceRequestWithCity {
	r := response.ServiceRequestWithCity {
		ID: request.ID,
		Date: request.Date,
		Cost: request.Cost,
		Description: request.Description,
		HasBeenReviewed: request.HasBeenReviewed,
		KindOfService: request.KindOfService,
		Status: request.Status,
	}

	address := response.AddressWithCity{
		ID: request.DeliveryAddress.ID,
		IndoorNumber: request.DeliveryAddress.IndoorNumber,
		OutdoorNumber: request.DeliveryAddress.OutdoorNumber,
		Street: request.DeliveryAddress.Street,
		Suburb: request.DeliveryAddress.Suburb,
		City: response.City {
			ID: request.DeliveryAddress.City.ID,
			Name: request.DeliveryAddress.City.Name,
		},
	}  
	r.DeliveryAddress = address

	provider := response.ServiceProviderInRequest{
		ID: request.ServiceProvider.ID,
		Names: request.ServiceProvider.User.Names,
		Lastname: request.ServiceProvider.User.Lastname,
		BusinessName: request.ServiceProvider.BusinessName,
	}
	r.ServiceProvider = provider

	requester := response.ServiceRequesterInRequest{
		ID: request.ServiceRequester.ID,
		Names: request.ServiceRequester.User.Names,
		Lastname: request.ServiceRequester.User.Lastname,
	}
	r.ServiceRequester = requester

	return r
}

func CreateRequestsDetailsAsResponse(request entity.ServiceRequest) response.ServiceRequestDetails {
	r := response.ServiceRequestDetails {
		ID: request.ID,
		Date: request.Date,
		KindOfService: request.KindOfService,
		Status: request.Status,
	}

	return r
}
