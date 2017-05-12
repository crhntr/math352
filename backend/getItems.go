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
		ID          string            `json:"id"`
		Title       string            `json:"title"`
		Body        string            `json:"body"`
		Link        string            `json:"link"`
		LikelyClass Class             `json:"likely_class"`
		Classes     map[Class]float64 `json:"classes"`
	}

	itemRecords := []ItemRecord{}

	itemsMut.Lock()
	defer itemsMut.Unlock()
	classifierMutex.Lock()
	defer classifierMutex.Unlock()

	for i, item := range items {
		ir := ItemRecord{
			ID:      strconv.Itoa(i),
			Title:   item.Title(),
			Link:    fmt.Sprintf("/api/item/%d", i),
			Body:    item.Body(),
			Classes: map[Class]float64{},
		}

		if classifier.Learned() > 0 {
			scores, likely, _ := classifier.LogScores(Tokenize(item.Title() + " " + item.Body()))

			ir.LikelyClass = classifier.Classes[likely]
			for i, class := range classifier.Classes {
				ir.Classes[class] = scores[i]
			}
		}

		itemRecords = append(itemRecords, ir)
	}
	c.JSON(200, gin.H{
		"data": itemRecords,
	})
}
