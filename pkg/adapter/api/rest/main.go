package restapi

import (
	"template/example/pkg/adapter/api/rest/album"
	"template/example/pkg/adapter/api/rest/billing"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	album.Route(router)
	billing.Route(router)
}
