package main

import (
	"log"
	"time"

	. "github.com/crhntr/math352/internal"
	"github.com/crhntr/math352/pubmed"
)

func startLoadDataJob() {
	log.Println("starting loadData Job")
	tkr := time.NewTicker(24 * time.Hour)

	tm := time.Now()
	tm = time.Date(tm.Year(), tm.Month(), tm.Day()+1, 2, 0, 0, 0, tm.Location())
	log.Printf("next update will be at %s", tm)
	time.Sleep(time.Until(tm))
	go func() {
		for range tkr.C {
			cleanupItems()
			log.Println("starting loadData Job")
			loadData(365)
			log.Println("done with loadData Job")
		}
	}()
}

func loadData(days int) {
	back := time.Now().Add(-1 * time.Duration(days) * 24 * time.Hour)

	q := pubmed.Query{
		Query: "talimogene laherparepvec [All Fields]",
	}

	tempItems := []Item{}

	timeIter := time.Now()
	for timeIter.After(back) {
		nItems, err := q.FetchItemsForDay(timeIter)
		if err != nil {
			log.Print(err)
			continue
		}
		if len(nItems) > 0 {
			log.Println("adding items")
			log.Println(nItems)

			for _, itm := range nItems {
				tempItems = append(tempItems, itm)
			}

		}
		timeIter = timeIter.Add(24 * (-1) * time.Hour)
	}

	itemsMut.Lock()
	items = tempItems
	itemsMut.Unlock()
}
