package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getItems(c *gin.Context) {
	type ItemRecord struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
		Link  string `json:"link"`
	}

	itemRecords := []ItemRecord{}

	itemsMut.Lock()
	defer itemsMut.Unlock()

	for i, item := range items {
		itemRecords = append(itemRecords, ItemRecord{
			ID:    strconv.Itoa(i),
			Title: item.Title(),
			Link:  fmt.Sprintf("/api/item/%d", i),
			Body:  item.Body(),
		})
	}
	c.JSON(200, gin.H{
		"data": itemRecords,
	})
}
