package pubmed

import (
	"encoding/xml"

	"gopkg.in/mgo.v2/bson"
)

type Query struct {
	ID    bson.ObjectId `json:"id" bson:"_id"`
	Query string        `json:"query_text" bson:"query_text"`
}

type ArticleSet struct {
	Name     xml.Name  `xml:"PubmedArticleSet"`
	Articles []Article `xml:"PubmedArticle"`
}

type Article struct {
	ID           bson.ObjectId `xml:"-" json:"id" bson:"_id"`
	FeedID       bson.ObjectId `xml:"-" json:"feed_id" bson:"feed_id"`
	Name         xml.Name      `xml:"PubmedArticle" json:"-" bson:"-"`
	ArticleTitle string        `xml:"MedlineCitation>Article>ArticleTitle" json:"title" bson:"title"`
	AbstractText string        `xml:"MedlineCitation>Article>Abstract>AbstractText" json:"abstract" bson:"abstract-"`
}

type ESearchResult struct {
	Count    int
	RetMax   int
	RetStart int
	QueryKey string
	WebEnv   string
	IDs      []string `xml:"IdList>Id"`
}
