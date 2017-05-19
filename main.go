package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/crhntr/math352/pubmed"
	"github.com/crhntr/naivegopher"
	"github.com/gin-gonic/gin"
)

const staticDirectoryPath = "static/"

var (
	verbose    bool
	classifier = naivegopher.NewClassifier()

	// itemsMut = &sync.Mutex{}
	items []*pubmed.Article

	// indexMut     = &sync.Mutex{}
	index   int = 0
	fetched     = false
)

func init() {
	verbose = *flag.Bool("verbose", true, "verbose logging")
}

func main() {
	// startCleanupJob()
	loadData(30)

	endLoadDataJob := make(chan struct{})
	startLoadDataJob(endLoadDataJob)
	// time.AfterFunc(30*time.Second, func() {
	// 	endLoadDataJob <- struct{}{}
	// })

	router := gin.Default()

	router.Use(
		func(c *gin.Context) {
			log.Println(c.Request.URL)
			log.Println(len(items))
			c.Next()
			log.Println(len(items))
		},
	)

	// router.Use(regesterRequestMiddleare)
	router.StaticFile("/", staticDirectoryPath+"index.html")
	router.StaticFS("/src/", http.Dir(staticDirectoryPath))
	// router.GET("/api/feed/:id/items", setDB, getFeedItemsHandler) // with optional params: year, month day

	// router.POST("/act/item/fetch", fetch)

	router.GET("/api/class", func(c *gin.Context) {
		c.JSON(200, gin.H{"classes": classifier.CategoryNames()})
	})

	router.PATCH("/api/item/:id/classes", patchItemClasses)
	router.GET("/api/item/:id", getItem)
	router.GET("/api/items", getItems)

	router.Run()
}
