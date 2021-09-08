package endpoints

import (
	"encoding/json"
	"finansial-service/main/models/requiests"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var Transaction = false
var TransactionUserSenderId string
var TransactionUserReceiverId string
var Amount string

func MakeTransaction(c *gin.Context) {

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Incorrect JSON body. Cant deserialize it.")
		return
	}

	var request requiests.Transaction
	json.Unmarshal(jsonData, &request)

	TransactionUserSenderId = request.SenderUserID
	TransactionUserReceiverId = request.ReceiverUserID
	Amount = request.Amount

	Transaction = true
	MakeAddition(c)
	GetWriteDown(c)
	Transaction = false

	TransactionUserReceiverId = ""
	TransactionUserSenderId = ""

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction completed",
	})
	return

}
