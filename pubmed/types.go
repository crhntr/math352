package pubmed

import (
	"encoding/xml"
)

type Query struct {
	Query string `json:"query_text" bson:"query_text"`
}

type ArticleSet struct {
	Name     xml.Name   `xml:"PubmedArticleSet"`
	Articles []*Article `xml:"PubmedArticle"`
}

type Article struct {
	Name          xml.Name `xml:"PubmedArticle" json:"-" bson:"-"`
	ArticleTitle  string   `xml:"MedlineCitation>Article>ArticleTitle" json:"title" bson:"title"`
	AbstractText  string   `xml:"MedlineCitation>Article>Abstract>AbstractText" json:"abstract" bson:"abstract-"`
	DatePublished Date     `xml:"MedlineCitation>Article>ArticleDate" json:"date_published" bson:"abstract-"`
	DateAdded     Date     `xml:"MedlineCitation>Article>PubmedData>History>PubMedPubDate" json:"date_added" bson:"abstract-"`
	done          bool
	readIndex     int
}

type ESearchResult struct {
	Count    int
	RetMax   int
	RetStart int
	QueryKey string
	WebEnv   string
	IDs      []string `xml:"IdList>Id"`
}
