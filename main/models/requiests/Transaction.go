package requiests

type Transaction struct {
	SenderUserID   string `json:"sender_user_id"`
	ReceiverUserID string `json:"receiver_user_id"`
	Amount         string `json:"amount"`
}
