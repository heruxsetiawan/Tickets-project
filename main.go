package main

import (
	"log"
	"tickets-project/controller"
	"tickets-project/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//main
	models.InitializeDB()

	router := gin.Default()
	router.Use(controller.LoggerMiddleware())

	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	authenticated := router.Group("/")
	authenticated.Use(controller.JWTMiddleware())
	{
		authenticated.GET("/tickets", controller.GetTickets)
		authenticated.POST("/tickets", controller.CreateTicket)
		authenticated.PUT("/tickets/:id", controller.UpdateTicket)
		authenticated.DELETE("/tickets/:id", controller.DeleteTicket)

		authenticated.GET("/tickets/:ticketID/tasks", controller.GetTasks)
		authenticated.POST("/tickets/tasks", controller.CreateTask)
		authenticated.PUT("/tasks/:id", controller.UpdateTask)
		authenticated.DELETE("/tasks/:id", controller.DeleteTask)
	}

	log.Fatal(router.Run(":8080"))
}
