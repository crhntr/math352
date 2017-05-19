package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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
	// itemsMut.Lock()
	if id > len(items) || id < 0 {
		c.JSON(404, gin.H{
			"error": "item with id not found: " + err.Error(),
		})
	}
	// defer itemsMut.Unlock()

	data := Data{}
	if err = c.BindJSON(&data); err != nil {
		c.JSON(404, gin.H{
			"error": "invalid data: " + err.Error(),
		})
		return
	}
	fmt.Println(data.Classes)
	for _, class := range data.Classes {
		log.Println("here")
		classifier.Learn(class, strings.NewReader(
			items[id].Title()+" "+items[id].Body(),
		))
	}
}
