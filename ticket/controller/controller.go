package controller

import (
	//"fmt"

	"net/http"
	"strings"
	"ticket/domain"
	"ticket/services"
	"ticket/utils"

	"github.com/gin-gonic/gin"
)

func CreateTrain(c *gin.Context) {
	var newTicket domain.Ticket
	if err := c.ShouldBindJSON(&newTicket); err != nil {
		resterr := utils.BadRequest("Invalid JSON")
		c.JSON(resterr.Status, resterr)
		return
	}

	result, resterr := services.AddTrains(newTicket)

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	}
	c.JSON(http.StatusOK, result)

}

func GetTrain(c *gin.Context) {

	trainnoParam := c.Query("trainno")

	if strings.TrimSpace(trainnoParam) == "" {
		resterr := utils.BadRequest("Train No Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	result, resterr := services.GetTrain(trainnoParam)

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetAllTrain(c *gin.Context) {

	result, resterr := services.GetAllTrain()

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	}
	c.JSON(http.StatusOK, result)
}


func DeleteTrain(c *gin.Context) {
	trainnoParam := c.Query("trainno")

	if strings.TrimSpace(trainnoParam) == "" {
		resterr := utils.BadRequest("Train No Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	result, resterr := services.DeleteTrain(trainnoParam)

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func UpdateTrain(c *gin.Context) {
	trainnoParam := c.Query("trainno")

	var ticket domain.Ticket

	if err := c.ShouldBindJSON(&ticket); err != nil {
		resterr := utils.BadRequest("Invalid JSON")
		c.JSON(resterr.Status, resterr)
		return
	}

	if strings.TrimSpace(trainnoParam) == "" {
		resterr := utils.BadRequest("Train No Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	ticket.TrainNO = trainnoParam

	isParital := c.Request.Method == http.MethodPatch

	result, resterr := services.UpdateTrain(isParital, ticket, trainnoParam)

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	}
	c.JSON(http.StatusOK, result)
}
