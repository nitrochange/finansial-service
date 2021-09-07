package endpoints

import (
	"encoding/json"
	"finansial-service/main/dao"
	"finansial-service/main/models/requiests"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GetWriteDown(c *gin.Context) {

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	var request requiests.WriteOFFJSON
	json.Unmarshal(jsonData, request)
	var id = request.ID
	var amount = request.Amount
	fmt.Println(jsonData)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Incorrect JSON body. Cant deserialize it.")
		return
	}

	for _, a := range dao.Users {
		if a.ID == id {
			if a.Balance >= amount {
				a.Balance -= amount
				c.IndentedJSON(http.StatusOK, a)
				return
			} else {
				c.IndentedJSON(http.StatusNotImplemented, "User balance is less than requested amount, user balance cannot be < 0")
				return
			}
		}
	}

	c.IndentedJSON(http.StatusNotFound, "Requested user was not found")
}
