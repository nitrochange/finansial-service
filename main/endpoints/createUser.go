package endpoints

import (
	"encoding/json"
	"finansial-service/main/dao"
	"finansial-service/main/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(c *gin.Context) {

	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	var id = uuid.New().String()
	var user models.User
	json.Unmarshal(jsonData, &user)

	error := dao.DBConnect.Insert(&models.User{
		ID:          id,
		FirstName:   user.FirstName,
		SecondName:  user.SecondName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Balance:     "0",
		Address:     user.Address,
	})

	if error != nil {
		log.Printf("Can not insert user into database")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not insert user into database",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"userId": id,
	})

	return
}
