package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	. "github.com/crhntr/math352/pubmed"
	"github.com/gin-gonic/gin"
)

var (
	lastFetchedMut = &sync.Mutex{}
	lastFetched    time.Time
)

func fetch(c *gin.Context) {
	if fetched {
		func() {
			lastFetchedMut.Lock()
			defer lastFetchedMut.Unlock()
			since := time.Since(lastFetched)
			if since < 5*time.Minute {
				c.JSON(http.StatusLocked, gin.H{
					"error": fmt.Sprintf("fetch disabled for a while (it will be unlocked in %f minutes)", since.Minutes()),
				})
			}
		}()

		// if forced := c.Query("force"); forced != "true" {
		// 	c.JSON(http.StatusLocked, gin.H{
		// 		"error": "fetch disabled for a while (after an hour of no use it will be available again)",
		// 	})
		// 	return
		// } else {
		// 	log.Print("FORCED fetch started")
		// 	cleanupItems()
		// }
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

	func() {
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
		time.Sleep(5 * time.Second)
	}()
	c.Status(http.StatusAccepted)
}
