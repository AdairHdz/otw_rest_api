package mapper

import (
	"github.com/AdairHdz/OTW-Rest-API/response"
	"github.com/AdairHdz/OTW-Rest-API/entity"
)

func CreateProviderAddAsResponse(provider *entity.ServiceProvider) response.User{
	r := response.User {
		ID: provider.ID,
		UserID: provider.User.ID,
		Names: provider.User.Names,
		LastName: provider.User.Lastname,
		EmailAddress: provider.User.Account.EmailAddress,
		UserType: provider.User.Account.UserType,
		StateID: provider.User.StateID,
	}
	return r
}

func CreateRequesterAddAsResponse(requester *entity.ServiceRequester) response.User{
	r := response.User {
		ID: requester.ID,
		UserID: requester.User.ID,
		Names: requester.User.Names,
		LastName: requester.User.Lastname,
		EmailAddress: requester.User.Account.EmailAddress,
		UserType: requester.User.Account.UserType,
		StateID: requester.User.StateID,
	}
	return r
}