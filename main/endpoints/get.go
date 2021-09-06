package endpoints

import (
	"finansial-service/main/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dao.Albums)
}

