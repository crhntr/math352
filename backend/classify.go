package main

import (
	"strconv"

	"github.com/gin-gonic/gin"

	. "github.com/crhntr/bayesian"
	. "github.com/crhntr/litsphere/internal"
)

func classify(c *gin.Context) {
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

	data := Data{}
	if err = c.BindJSON(&data); err != nil {
		c.JSON(404, gin.H{
			"error": "invalid id: " + err.Error(),
		})
		return
	}

	for _, class := range data.Classes {
		classifier.Learn(Tokenize(items[id].Title()), Class(class))
		classifier.Learn(Tokenize(items[id].Body()), Class(class))
	}
}
