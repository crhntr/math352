package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getItems(c *gin.Context) {
	type ItemRecord struct {
		ID         string             `json:"id"`
		Title      string             `json:"title"`
		Body       string             `json:"body"`
		Link       string             `json:"link"`
		Categories map[string]float64 `json:"categories"`
	}

	itemRecords := []ItemRecord{}

	// itemsMut.Lock()
	// defer itemsMut.Unlock()
	for i, item := range items {
		ir := ItemRecord{
			ID:         strconv.Itoa(i),
			Title:      item.Title(),
			Link:       fmt.Sprintf("/api/item/%d", i),
			Body:       item.Body(),
			Categories: map[string]float64{},
		}

		probabilities := classifier.ProbableCategoreies(strings.NewReader(
			item.Title() + " " + item.Body(),
		))

		for i, name := range classifier.CategoryNames() {
			ir.Categories[name] = probabilities[i]
		}

		itemRecords = append(itemRecords, ir)
	}
	c.JSON(200, gin.H{
		"data": itemRecords,
	})
}
