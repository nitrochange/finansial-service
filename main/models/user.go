package models

// Структура данных
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type User struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	SecondName  string `json:"second_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Balance     int    `json:"balance"`
	Address     string `json:"address"`
}
