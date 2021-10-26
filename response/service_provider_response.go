package response

type ServiceProvider struct{
	ID 				string		`json:"id"`
	Names 			string 		`json:"names"`
	Lastname 		string 		`json:"lastNames"`
	AverageScore 	float64 	`json:"averageScore"`
	PriceRate 		float64 	`json:"priceRate"`
	BusinessName	string 		`json:"businessName"`
}