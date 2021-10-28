package response

type Address struct{
	ID 				string	`json:"id"`
	IndoorNumber 	string	`json:"indoorNumber"`
	OutdoorNumber 	string	`json:"outdoorNumber"`
	Street 			string	`json:"street"`
	Suburb 			string	`json:"suburb"`
}

type AddressWithCity struct{
	ID 				string	`json:"id"`
	IndoorNumber 	string	`json:"indoorNumber"`
	OutdoorNumber 	string	`json:"outdoorNumber"`
	Street 			string	`json:"street"`
	Suburb 			string	`json:"suburb"`
	City 			City 	`json:"city"`
}