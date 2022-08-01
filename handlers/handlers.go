package handlers

import (
	"awesomeProject/database"
	local_rabbit "awesomeProject/rabbit"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerInstance struct {
	queue *local_rabbit.LocalQueue
	db    *sql.DB
}

func InitHandlers(queue *local_rabbit.LocalQueue, DB *sql.DB) *HandlerInstance {
	return &HandlerInstance{
		queue: queue,
		db:    DB,
	}
}

func (hi *HandlerInstance) GetH(c *gin.Context) {
	tasks, err := database.ReturnTasks(hi.db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	}
	c.JSON(http.StatusOK, tasks)
}

func (hi *HandlerInstance) PostH(c *gin.Context) {
	var json database.Task
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.InsertTask(hi.db, json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Fail to add task to db: " + err.Error(),
		})
	}
	if err := hi.queue.AddToQueue(json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Fail to add task to queue: " + err.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{"status": "task was added"})
}
