package app

import "ticket/controller"

func Routers() {
	r.POST("/ticket/addTrain", controller.CreateTrain)
	r.GET("/ticket/getTrain", controller.GetTrain)
	r.GET("/ticket/getAllTrain", controller.GetAllTrain)
	r.DELETE("/ticket/deleteTrain", controller.DeleteTrain)
	r.POST("/ticket/updateTrain", controller.UpdateTrain)
	
}
