package endpoints

import (
	"encoding/json"
	"finansial-service/main/dao"
	"finansial-service/main/models/requiests"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func MakeAddition(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Incorrect JSON body. Cant deserialize it.")
		return
	}

	var request requiests.WriteOFFJSON
	json.Unmarshal(jsonData, &request)

	var id = request.ID
	amount, err := strconv.Atoi(request.Amount)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Cannot convert amount field to int value.")
	}

	var utils = dao.Utils{}
	var user = utils.GetUser(id)
	balance, err := strconv.Atoi(user.Balance)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Can not parse balance from database")
		log.Fatal("Can not parse balance from database")
		return
	}
	user.Balance = strconv.Itoa(balance + amount)
	//utils.UpdateUser()

	c.IndentedJSON(http.StatusOK, user)

	//for i:= 0; i < len(dao.Users); i++ {
	//	temp := &dao.Users[i]
	//	if temp.ID == id {
	//			temp.Balance += amount
	//			c.IndentedJSON(http.StatusOK, temp)
	//			return
	//	}
	//}

	c.IndentedJSON(http.StatusNotFound, "Requested user was not found")
}
