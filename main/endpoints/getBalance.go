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
)

func GetBalance(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Incorrect JSON body. Cant deserialize it.")
		return
	}

	var request requiests.GetBalance
	json.Unmarshal(jsonData, &request)

	var id = request.ID

	var user = &models.User{}
	err2 := dao.DBConnect.Model(user).Where("? = ?", pg.Ident("id"), id).Select()
	if err2 != nil {
		log.Printf("Cant find user in DB")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cant find user in DB",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         id,
		"firstName":  user.FirstName,
		"secondName": user.SecondName,
		"balance":    user.Balance,
	})
	return

}
