package main

import (
	"awesomeProject/database"
	"awesomeProject/handlers"
	local_rabbit "awesomeProject/rabbit"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db, _ := database.CreateDBConnection(
		"localhost",
		"5400",
		"misha",
		"12345",
		"test",
	)

	defer db.Close()
	localQueue, err := local_rabbit.InitQueue("amqp://guest:guest@localhost:5672/", "TestQueue")
	if err != nil {
		log.Fatal("Fail to create queue: ", err)
	}
	if err := database.CreateTable(db); err != nil {
		log.Println("Fail to create tasks table: ", err)
	}
	hi := handlers.InitHandlers(localQueue, db)

	r := gin.Default()
	r.GET("/task", hi.GetH)

	r.POST("/task", hi.PostH)

	log.Fatal(r.Run(":8001")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
