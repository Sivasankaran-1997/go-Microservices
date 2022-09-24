package app

import "buy/controller"

func Routers() {
	r.POST("/buy/buyTicket", controller.BuyTicket)
	r.GET("/buy/getUserbuy", controller.GetUserBuy)
	r.DELETE("/buy/deleteUserbuy", controller.DeleteUserBuy)
	r.GET("/buy/getAllUserTickets", controller.GetAllUserTickets)
}
