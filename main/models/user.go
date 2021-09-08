package models

type User struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	SecondName  string `json:"second_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Balance     string `json:"balance"`
	Address     string `json:"address"`
}
