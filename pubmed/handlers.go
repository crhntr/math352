package pubmed

import (
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func MW(db *mgo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		db = db.With(db.Session.Clone())
		defer db.Session.Close()
		c.Next()
	}
}
