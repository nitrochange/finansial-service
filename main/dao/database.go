package dao

import (
	"finansial-service/main/models"
)

// Albums slice to seed record album data.
var Albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var Users = []models.User{
	{ID: "1", FirstName: "Nikita", SecondName: "Baranov", Email: "baranovnikita@mail.ru", PhoneNumber: "7999999999", Balance: 0, Address: "Moscow"},
	{ID: "2", FirstName: "Sasha", SecondName: "Lapshin", Email: "sasha@mail.ru", PhoneNumber: "71111111111", Balance: 0, Address: "Oslo"},
}

var Transactions = []models.Transaction{
	{SenderUserID: "1", ReceiverUserID: "2", Amount: 34},
}
