package response

type Address struct{
	IndoorNumber 	string	`json:"indoorNumber"`
	OutdoorNumber 	string	`json:"outdoorNumber"`
	Street 			string	`json:"street"`
	Suburb 			string	`json:"suburb"`
}