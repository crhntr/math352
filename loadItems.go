package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	. "github.com/crhntr/math352/pubmed"
	"github.com/gin-gonic/gin"
)

func fetch(c *gin.Context) {
	if fetched {
		return
	}
	fetched = true
	log.Print()
	daysParam := c.Query("days")
	if daysParam == "" {
		daysParam = "30"
	}
	days, err := strconv.Atoi(daysParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	queryString := ""
	queryBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		queryString = "happy [Title]"
	} else {
		queryString = string(queryBytes)
	}
	log.Println(queryString)
	q := Query{
		Query: queryString,
	}
	log.Printf("Query: %s", q)

	nDaysAgo := time.Now().Add(-1 * time.Duration(days) * 24 * time.Hour)
	log.Printf("nDaysAgo: %s", nDaysAgo)

	go func() {
		timeIter := time.Now()
		for timeIter.After(nDaysAgo) {
			nItems, err := q.FetchItemsForDay(timeIter)
			if err != nil {
				log.Print(err)
				continue
			}
			if len(nItems) > 0 {
				log.Println("adding items")
				log.Println(nItems)
				for _, itm := range nItems {
					itemsMut.Lock()
					items = append(items, itm)
					itemsMut.Unlock()
				}
			}

			timeIter = timeIter.Add(24 * (-1) * time.Hour)
		}

		log.Printf("len(item)= %d", len(items))
	}()
	c.Status(http.StatusAccepted)
}
