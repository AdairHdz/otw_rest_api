package mapper

import (
	"github.com/AdairHdz/OTW-Rest-API/entity"
	"github.com/AdairHdz/OTW-Rest-API/response"
)

func CreateAddressesAsResponse(address entity.Address) response.Address{
	r := response.Address {
		ID: address.ID,
		IndoorNumber: address.IndoorNumber,
		OutdoorNumber: address.OutdoorNumber,
		Street: address.Street,
		Suburb: address.Suburb,
	}
	return r
}
