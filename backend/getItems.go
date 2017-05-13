package main

import (
	"fmt"
	"strconv"

	. "github.com/crhntr/bayesian"
	. "github.com/crhntr/math352/internal"
	"github.com/gin-gonic/gin"
)

func getItems(c *gin.Context) {
	type ItemRecord struct {
		ID             string            `json:"id"`
		Title          string            `json:"title"`
		Body           string            `json:"body"`
		Link           string            `json:"link"`
		LikelyCategory Class             `json:"likely_category"`
		Categories     map[Class]float64 `json:"categories"`
	}

	itemRecords := []ItemRecord{}

	itemsMut.Lock()
	defer itemsMut.Unlock()
	classifierMutex.Lock()
	defer classifierMutex.Unlock()

	for i, item := range items {
		ir := ItemRecord{
			ID:         strconv.Itoa(i),
			Title:      item.Title(),
			Link:       fmt.Sprintf("/api/item/%d", i),
			Body:       item.Body(),
			Categories: map[Class]float64{},
		}

		scores, likely, _ := classifier.LogScores(Tokenize(item.Title() + " " + item.Body()))

		ir.LikelyCategory = classifier.Classes[likely]
		for i, class := range classifier.Classes {
			ir.Categories[class] = scores[i]
		}

		itemRecords = append(itemRecords, ir)
	}
	c.JSON(200, gin.H{
		"data": itemRecords,
	})
}
