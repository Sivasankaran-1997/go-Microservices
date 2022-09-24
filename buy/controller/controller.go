package controller

import (
	//"fmt"

	"net/http"
	"strings"

	//"strings"
	"buy/domain"
	"buy/services"
	"buy/utils"

	"github.com/gin-gonic/gin"
)

func BuyTicket(c *gin.Context) {
	var newBuy domain.Buy
	if err := c.ShouldBindJSON(&newBuy); err != nil {
		resterr := utils.BadRequest("Invalid JSON")
		c.JSON(resterr.Status, resterr)
		return
	}

	result, resterr := services.BuyTicket(newBuy)

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	}
	c.JSON(http.StatusOK, result)

}

func GetUserBuy(c *gin.Context) {

	buyIDParam := c.Query("buyid")

	if strings.TrimSpace(buyIDParam) == "" {
		resterr := utils.BadRequest("Buyer ID  Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	result, resterr := services.GetUserBuy(buyIDParam)

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetAllUserTickets(c *gin.Context) {

	userid := c.Query("userid")

	if strings.TrimSpace(userid) == "" {
		resterr := utils.BadRequest("User ID  Required")
		c.JSON(resterr.Status, resterr)
		return
	}
	result, resterr := services.GetAllUserTickets(userid)

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUserBuy(c *gin.Context) {
	buyIDParam := c.Query("buyid")

	if strings.TrimSpace(buyIDParam) == "" {
		resterr := utils.BadRequest("Buyer ID  Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	result, resterr := services.DeleteUserBuy(buyIDParam)

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	}
	c.JSON(http.StatusOK, result)
}
