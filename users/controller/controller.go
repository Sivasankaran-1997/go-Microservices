package controller

import (
	//"fmt"

	"encoding/json"
	"net/http"
	"os"
	"users/domain"
	"users/services"
	"users/utils"

	"strings"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser domain.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		resterr := utils.BadRequest("Invalid JSON")
		c.JSON(resterr.Status, resterr)
		return
	}

	result, resterr := services.CreateUser(newUser)

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetUser(c *gin.Context) {
	emailParam := c.Query("useremail")
	passwordParam := c.Query("userpassword")

	if strings.TrimSpace(emailParam) == "" {
		resterr := utils.BadRequest("User Email Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	if strings.TrimSpace(passwordParam) == "" {
		resterr := utils.BadRequest("Password Email Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	result, resterr := services.GetUser(emailParam, passwordParam)

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetTrain(c *gin.Context) {
	emailParam := c.Query("useremail")
	passwordParam := c.Query("userpassword")

	if strings.TrimSpace(emailParam) == "" {
		resterr := utils.BadRequest("User Email Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	if strings.TrimSpace(passwordParam) == "" {
		resterr := utils.BadRequest("Password Email Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	_, resterr := services.GetUser(emailParam, passwordParam)

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	} else {
		getallTrainapi := os.Getenv("GETALLTRAINAPI")
		resp, _ := http.Get(getallTrainapi)

		if resp.StatusCode == http.StatusOK {
			var ticket []domain.Ticket
			json.NewDecoder(resp.Body).Decode(&ticket)
			c.JSON(http.StatusOK, ticket)
		} else {
			var e utils.Resterr
			json.NewDecoder(resp.Body).Decode(&e)
			c.JSON(resp.StatusCode, e)
		}

	}
}

func DeleteUser(c *gin.Context) {
	emailParam := c.Query("useremail")
	passwordParam := c.Query("userpassword")

	if strings.TrimSpace(emailParam) == "" {
		resterr := utils.BadRequest("User Email Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	if strings.TrimSpace(passwordParam) == "" {
		resterr := utils.BadRequest("Password Email Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	result, resterr := services.DeleteUser(emailParam, passwordParam)

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func UpdateUser(c *gin.Context) {
	emailParam := c.Query("useremail")
	passwordParam := c.Query("userpassword")

	var newUser domain.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		resterr := utils.BadRequest("Invalid JSON")
		c.JSON(resterr.Status, resterr)
		return
	}

	if strings.TrimSpace(emailParam) == "" {
		resterr := utils.BadRequest("User Email Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	if strings.TrimSpace(passwordParam) == "" {
		resterr := utils.BadRequest("Password Email Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	newUser.UserEmail = emailParam
	newUser.UserPassword = passwordParam
	isParital := c.Request.Method == http.MethodPatch

	result, resterr := services.UpdateUser(isParital, newUser, emailParam, passwordParam)

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func UserBuyTicket(c *gin.Context) {

	useremailparam := c.Query("useremail")
	userpasswordparam := c.Query("userpassword")

	var newBuy domain.Buy
	if err := c.ShouldBindJSON(&newBuy); err != nil {
		resterr := utils.BadRequest("Invalid JSON")
		c.JSON(resterr.Status, resterr)
		return
	}

	if strings.TrimSpace(useremailparam) == "" {
		resterr := utils.BadRequest("User Email Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	if strings.TrimSpace(userpasswordparam) == "" {
		resterr := utils.BadRequest("User PassWord Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	result, resterr := services.UserBuyTicket(newBuy, useremailparam, userpasswordparam)

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func UserAllBuyTicket(c *gin.Context) {
	emailParam := c.Query("useremail")
	passwordParam := c.Query("userpassword")

	if strings.TrimSpace(emailParam) == "" {
		resterr := utils.BadRequest("User Email Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	if strings.TrimSpace(passwordParam) == "" {
		resterr := utils.BadRequest("Password Email Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	result, resterr := services.GetUser(emailParam, passwordParam)

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	} else {
		getuserAllBuytickets := os.Getenv("USERALLBUYTICKET")
		resp, _ := http.Get(getuserAllBuytickets + result.UserID)

		if resp.StatusCode == http.StatusOK {
			var buy []domain.Buy
			json.NewDecoder(resp.Body).Decode(&buy)
			c.JSON(http.StatusOK, buy)
		} else {
			var e utils.Resterr
			json.NewDecoder(resp.Body).Decode(&e)
			c.JSON(resp.StatusCode, e)
		}

	}
}

func UserCancelTicket(c *gin.Context) {
	emailParam := c.Query("useremail")
	passwordParam := c.Query("userpassword")
	userticketParam := c.Query("userticket")

	if strings.TrimSpace(emailParam) == "" {
		resterr := utils.BadRequest("User Email Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	if strings.TrimSpace(passwordParam) == "" {
		resterr := utils.BadRequest("Password Email Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	if strings.TrimSpace(userticketParam) == "" {
		resterr := utils.BadRequest("User Ticket Required")
		c.JSON(resterr.Status, resterr)
		return
	}

	resterr := services.UserCancel(emailParam, passwordParam, userticketParam)

	if resterr != nil {
		c.JSON(resterr.Status, resterr)
		return
	}
	c.JSON(http.StatusOK, "Deleted")
}
