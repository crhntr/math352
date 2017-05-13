package main

import (
	"log"
	"sync"
	"time"

	. "github.com/crhntr/math352/internal"
)

var (
	lastRequestMut *sync.Mutex
	lastRequest    time.Time
)

func registerRequest() {
	lastRequestMut.Lock()
	defer lastRequestMut.Unlock()
	lastRequest = time.Now()
}

func startCleanupJob() {
	log.Println("starting cleanup Job")
	cleanupJobTicker := time.NewTicker(2 * time.Hour)
	if lastRequestMut == nil {
		lastRequestMut = &sync.Mutex{}
	}
	registerRequest()

	go func() {
		for range cleanupJobTicker.C {
			log.Println("attepting cleanup Job")
			if time.Since(lastRequest) > time.Hour {
				cleanupItems()
			}
		}
	}()
}

func cleanupItems() {
	itemsMut.Lock()
	defer itemsMut.Unlock()
	if len(items) > 0 {
		log.Printf("cleaning up items (%d items removed)", len(items))
		items = []Item{}
	}
}
