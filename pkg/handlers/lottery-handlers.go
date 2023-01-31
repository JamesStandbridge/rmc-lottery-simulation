package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type postData struct {
	Caracteristics []int `json:"caracteristics"`
}

func CreateLottery(c *gin.Context) {
	var body postData

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": body})
}
