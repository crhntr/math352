package main

import (
	"flag"
	"net/http"
	"sync"

	. "github.com/crhntr/bayesian"
	. "github.com/crhntr/math352/internal"
	"github.com/gin-gonic/gin"
)

const staticDirectoryPath = "static/"

var (
	verbose    bool
	classifier *Classifier
)

var (
	itemsMut *sync.Mutex
	items    []Item

	indexMut *sync.Mutex
	index    int = 0
	fetched      = false
)

func init() {
	verbose = *flag.Bool("verbose", true, "verbose logging")
}

func main() {
	itemsMut = &sync.Mutex{}
	indexMut = &sync.Mutex{}

	startCleanupJob()

	var err error
	classifier = NewClassifier("relevant", "irrelevent")
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.StaticFile("/", staticDirectoryPath+"index.html")
	router.StaticFS("/src/", http.Dir(staticDirectoryPath))
	// router.GET("/api/feed/:id/items", setDB, getFeedItemsHandler) // with optional params: year, month day

	router.GET("/act/item/next", nextItem)
	router.POST("/act/item/fetch", fetch)

	router.GET("/api/class", func(c *gin.Context) {
		c.JSON(200, gin.H{"classes": classifier.Classes})
	})
	router.GET("/api/class/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, classifier.WordsByClass(Class(name)))
	})
	router.GET("/api/class/:name/items", getClass)

	router.PATCH("/api/item/:id/classes", classify)
	router.GET("/api/item/:id", getItem)
	router.GET("/api/item", getItems)

	router.Run()
}
