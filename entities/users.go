package entities

type User struct {
	UUIDBaseModel
	Email        string `json:"email"`
	FirstName    string `json:"fname"`
	LastName     string `json:"lname"`
	PasswordHash string `json:"passwordhash"`
}
