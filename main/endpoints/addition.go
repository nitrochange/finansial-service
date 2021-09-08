package endpoints

import (
	"encoding/json"
	"finansial-service/main/dao"
	"finansial-service/main/models"
	"finansial-service/main/models/requiests"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func MakeAddition(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Incorrect JSON body. Cant deserialize it.")
		return
	}

	var request requiests.WriteOFFJSON
	json.Unmarshal(jsonData, &request)

	var id string

	if Transaction {
		id = TransactionUserReceiverId
	} else {
		id = request.ID
	}

	amount, err := strconv.Atoi(request.Amount)

	if Transaction {
		amount, err = strconv.Atoi(Amount)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Cannot convert amount field to int value.")
	}

	var user = &models.User{}
	error := dao.DBConnect.Model(user).Where("? = ?", pg.Ident("id"), id).Select()
	if error != nil {
		log.Printf("Cant find user in DB")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cant find user in DB",
		})
		return
	}
	var userBalance, _ = strconv.Atoi(user.Balance)
	user.Balance = strconv.Itoa(userBalance + amount)

	_, err = dao.DBConnect.Model(user).Set("balance = ?balance").Where("id = ?id").Update()

	if error != nil {
		log.Printf("Can not insert new user balance into database")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not insert new user balance into database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Addition completed",
	})
	return

}
