package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{
			"error": "invalid id: " + err.Error(),
		})
		return
	}

	itemsMut.Lock()
	defer itemsMut.Unlock()
	c.JSON(200, gin.H{
		"data": items[id],
	})
}
