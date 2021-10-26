package response

type PriceRateWorkingDays struct{
	ID 				string			`json:"id"`
	StartingHour 	string			`json:"startingHour"`
	EndingHour 		string			`json:"endingHour"`
	Price        	float64			`json:"price"`
	KindOfService 	int				`json:"kindOfService"`
	City  			City			`json:"city"`
	WorkingDays  	[]int	`json:"workingDays"`
}