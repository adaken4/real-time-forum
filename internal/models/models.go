package models

// User struct for parsing request data
type User struct {
	Nickname        string `json:"nickname"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Age             int    `json:"age"`
	Gender          string `json:"gender"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}
