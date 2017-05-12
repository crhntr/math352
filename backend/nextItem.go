package main

import "github.com/gin-gonic/gin"

func nextItem(c *gin.Context) {
	itemsMut.Lock()
	defer itemsMut.Unlock()
	indexMut.Lock()
	defer indexMut.Unlock()

	if index < 0 || index >= len(items) {
		c.JSON(404, gin.H{"error": "invalid id"})
		return
	}

	c.JSON(200, gin.H{"data": items[index], "id": index, "type": "pubmed"})
	index++
	if index > len(items) {
		index = 0
	}
}
