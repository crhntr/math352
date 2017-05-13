package main

import (
	"log"
	"sync"
	"time"

	. "github.com/crhntr/math352/internal"
	"github.com/gin-gonic/gin"
)

var (
	lastRequestMut *sync.Mutex
	lastRequest    time.Time
)

func regesterRequestMiddleare(c *gin.Context) {
	registerRequest()
	c.Next()
}

func registerRequest() {
	lastRequestMut.Lock()
	defer lastRequestMut.Unlock()
	lastRequest = time.Now()
}

func startCleanupJob() {
	log.Println("starting cleanup Job")
	cleanupJobTicker := time.NewTicker(30 * time.Minute)
	if lastRequestMut == nil {
		lastRequestMut = &sync.Mutex{}
	}
	registerRequest()

	go func() {
		for range cleanupJobTicker.C {
			log.Println("attepting cleanup Job")
			if time.Since(lastRequest) > 15*time.Minute {
				cleanupItems()
			}
		}
	}()
}

func cleanupItems() {
	itemsMut.Lock()
	defer itemsMut.Unlock()
	fetched = false

	if len(items) > 0 {
		log.Printf("cleaning up items (%d items removed)", len(items))
		items = []Item{}
	}
}
