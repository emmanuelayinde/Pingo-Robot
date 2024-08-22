package controllers

import (
	"net/http"

	"github.com/emmanuelayinde/pingo-robot/config"
	"github.com/emmanuelayinde/pingo-robot/models"
	"github.com/gin-gonic/gin"
)

func CreateMonitor(c *gin.Context) {
	var monitor models.Monitor
	if err := c.ShouldBindJSON(&monitor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&monitor)
	c.JSON(http.StatusOK, monitor)
}

func GetMonitor(c *gin.Context) {
	id := c.Param("id")
	var monitor models.Monitor
	if err := config.DB.First(&monitor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monitor not found!"})
		return
	}
	c.JSON(http.StatusOK, monitor)
}
