package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jbrukh/bayesian"
)

func getItems(c *gin.Context) {
	type Category struct {
		Prob float64 `json:"prob"`
		Log  float64 `json:"log"`
	}
	type ItemRecord struct {
		ID             string                      `json:"id"`
		Title          string                      `json:"title"`
		Body           string                      `json:"body"`
		Link           string                      `json:"link"`
		LikelyCategory bayesian.Class              `json:"likely_category"`
		Categories     map[bayesian.Class]Category `json:"categories"`
	}

	itemRecords := []ItemRecord{}

	itemsMut.Lock()
	defer itemsMut.Unlock()
	classifierMutex.Lock()
	defer classifierMutex.Unlock()

	for i, item := range items {
		ir := ItemRecord{
			ID:             strconv.Itoa(i),
			Title:          item.Title(),
			Link:           fmt.Sprintf("/api/item/%d", i),
			Body:           item.Body(),
			LikelyCategory: "",
			Categories:     map[bayesian.Class]Category{},
		}

		if classified > 5 {
			lgscores, likely, _ := classifier.LogScores(Tokenize(item.Title() + " " + item.Body()))
			pbscores, likely, _ := classifier.ProbScores(Tokenize(item.Title() + " " + item.Body()))

			ir.LikelyCategory = classifier.Classes[likely]
			for i, class := range classifier.Classes {
				ir.Categories[class] = Category{
					Log:  lgscores[i],
					Prob: pbscores[i],
				}
			}
		} else {
			for _, dclass := range defaultClasses {
				ir.Categories[dclass] = Category{}
			}
		}

		itemRecords = append(itemRecords, ir)
	}
	c.JSON(200, gin.H{
		"data": itemRecords,
	})
}
