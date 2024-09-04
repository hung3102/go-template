package billing

import (
	"context"
	"gcim/example/pkg/domain/entity"
	billingService "gcim/example/pkg/domain/service/billing"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getBillings responds with the list of all billings as JSON.
func getBillings(c *gin.Context) {
	billings, err := billingService.GetBillings(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	result := []Billing{}
	for _, v := range billings {
		result = append(result, convertBilling(c, v))
	}

	c.IndentedJSON(http.StatusOK, result)
}

func convertBilling(_ context.Context, billing entity.Billing) Billing {
	return Billing{
		Id:   billing.Id,
		Cost: billing.Cost,
	}
}
