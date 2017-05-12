package main

import (
	"strconv"

	. "github.com/crhntr/math352/internal"
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

func getClass(c *gin.Context) {
	className := c.Param("name")
	if className == "" {
		c.JSON(404, gin.H{
			"error": "class name cannot be empty",
		})
		return
	}
	threshold, err := strconv.ParseFloat(c.Param("threshold"), 64)
	if err != nil {
		threshold = 0.5
	}

	classifiedItems := []Item{}

	for _, item := range items {
		scores, _, _ := classifier.ProbScores(append(Tokenize(item.Title()), Tokenize(item.Body())...))

		for i, cl := range classifier.Classes {
			if className == string(cl) && scores[i] > threshold {
				classifiedItems = append(classifiedItems, item)
			}
		}
	}

	c.JSON(200, gin.H{
		"items": classifiedItems,
	})
}
