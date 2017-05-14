package main

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jbrukh/bayesian"
)

func patchItemClasses(c *gin.Context) {
	type Data struct {
		Classes []string `json:"classes"`
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{
			"error": "invalid id: " + err.Error(),
		})
		return
	}
	itemsMut.Lock()
	if id > len(items) || id < 0 {
		c.JSON(404, gin.H{
			"error": "item with id not found: " + err.Error(),
		})
	}
	defer itemsMut.Unlock()

	data := Data{}
	if err = c.BindJSON(&data); err != nil {
		c.JSON(404, gin.H{
			"error": "invalid id: " + err.Error(),
		})
		return
	}
	classifierMutex.Lock()
	defer classifierMutex.Unlock()
	classified++
	for _, class := range data.Classes {
		classifier.Learn(append(Tokenize(items[id].Title()), Tokenize(items[id].Body())...), bayesian.Class(class))
	}
}
