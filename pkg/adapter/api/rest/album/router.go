package album

import "github.com/gin-gonic/gin"

func Route(router *gin.Engine) {
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
}
