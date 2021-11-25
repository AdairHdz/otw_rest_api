package response

type User struct {
	ID 			 string `json:"id"`
	SpecificUserID string `json:"specificUserId"`
	Names        string `json:"names"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
	UserType     int    `json:"userType"`
	StateID      string `json:"stateId"`
}