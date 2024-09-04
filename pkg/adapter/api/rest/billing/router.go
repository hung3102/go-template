package billing

import "github.com/gin-gonic/gin"

func Route(router *gin.Engine) {
	router.GET("/billings", getBillings)
	// router.GET("/billings/:id", getBillingById)
	// router.POST("/billing", postBilling)
}
