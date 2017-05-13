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
	verbose bool

	classifier      *Classifier
	classifierMutex *sync.Mutex

	itemsMut *sync.Mutex
	items    []Item

	indexMut *sync.Mutex
	index    int = 0
	fetched      = false

	defaultClasses = []Class{"relevant", "irrelevent"}
)

func init() {
	verbose = *flag.Bool("verbose", true, "verbose logging")
}

func main() {
	itemsMut = &sync.Mutex{}
	indexMut = &sync.Mutex{}

	startCleanupJob()

	classifier = NewClassifier(defaultClasses...)
	classifierMutex = &sync.Mutex{}
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

	router.PATCH("/api/item/:id/classes", patchItemClasses)
	router.GET("/api/item/:id", getItem)
	router.GET("/api/items", getItems)

	router.Run()
}
