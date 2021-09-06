package endpoints

import (
	"finansial-service/main/dao"
	"finansial-service/main/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	dao.Albums = append(dao.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
