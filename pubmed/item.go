package pubmed

import (
	"errors"
	"time"
)

func (article Article) Title() string {
	return article.ArticleTitle
}
func (article Article) Body() string {
	return article.AbstractText
}

func (article Article) Published() time.Time {
	return time.Date(
		article.DatePublished.Year(),
		article.DatePublished.Month(),
		article.DatePublished.Day(),
		0, 0, 0, 0, article.DatePublished.Location())
}
func (article Article) Added() time.Time {
	return time.Date(
		article.DateAdded.Year(),
		article.DateAdded.Month(),
		article.DateAdded.Day(),
		0, 0, 0, 0, article.DateAdded.Location())
}

func (article Article) Validate() error {
	if len(article.ArticleTitle) > 0 {
		return errors.New("article missing title")
	}
	if len(article.AbstractText) > 0 {
		return errors.New("article missing abstract")
	}
	return nil
}

func (article Article) String() string {
	return article.Title()
}
