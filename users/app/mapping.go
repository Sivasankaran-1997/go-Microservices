package app

import "users/controller"

func Routers() {
	r.POST("/users/create", controller.CreateUser)
	r.GET("/users/getuser", controller.GetUser)
	r.DELETE("/users/deleteuser", controller.DeleteUser)
	r.PATCH("/users/updateuser", controller.UpdateUser)
	r.GET("/users/getTrain", controller.GetTrain)
	r.POST("/users/userBuyTicket", controller.UserBuyTicket)
	r.GET("/users/userAllBuyTicket", controller.UserAllBuyTicket)
	r.DELETE("/users/cancelTicket", controller.UserCancelTicket)
}
