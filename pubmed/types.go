package pubmed

import (
	"encoding/xml"
)

type Query struct {
	Query string `json:"query_text" bson:"query_text"`
}

type ArticleSet struct {
	Name     xml.Name  `xml:"PubmedArticleSet"`
	Articles []Article `xml:"PubmedArticle"`
}

type Article struct {
	Name         xml.Name `xml:"PubmedArticle" json:"-" bson:"-"`
	ArticleTitle string   `xml:"MedlineCitation>Article>ArticleTitle" json:"title" bson:"title"`
	AbstractText string   `xml:"MedlineCitation>Article>Abstract>AbstractText" json:"abstract" bson:"abstract-"`
}

type ESearchResult struct {
	Count    int
	RetMax   int
	RetStart int
	QueryKey string
	WebEnv   string
	IDs      []string `xml:"IdList>Id"`
}
