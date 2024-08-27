package restapi

import (
	"gcim/example/pkg/adapter/api/rest/album"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	album.Route(router)
}
