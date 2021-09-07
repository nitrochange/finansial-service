package models

type Transaction struct {
	SenderUserID   string `json:"sender_user_id"`
	ReceiverUserID string `json:"receiver_user_id"`
	Amount         int    `json:"amount"`
}
