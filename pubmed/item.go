package pubmed

import (
	"errors"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	. "github.com/crhntr/math352/internal"
)

func (article Article) Title() string {
	return article.ArticleTitle
}
func (article Article) Body() string {
	return article.AbstractText
}

func (article Article) Published() time.Time {
	panic("not implemented")
}
func (article Article) Added() time.Time {
	panic("not implemented")
}

func (article Article) Validate() error {
	if len(article.ArticleTitle) > 0 {
		return errors.New("article missing title")
	}
	if len(article.AbstractText) > 0 {
		return errors.New("article missing abstract")
	}
	if !article.ID.Valid() {
		return errors.New("invalid article id")
	}
	if !article.FeedID.Valid() {
		return errors.New("invalid article feed_id")
	}
	return nil
}

func (article *Article) Create(db *mgo.Database) error {
	article.ID = bson.NewObjectId()
	if err := article.Validate(); err != nil {
		return err
	}
	return db.C(PubMedItemCollectionName).Insert(article)
}

func (article *Article) Save(db *mgo.Database) error {
	if err := article.Validate(); err != nil {
		return err
	}
	return db.C(PubMedItemCollectionName).UpdateId(article.ID, article)
}

func (article *Article) Delete(db *mgo.Database) error {
	return db.C(PubMedItemCollectionName).RemoveId(article.ID)
}

func (query Query) GetItems(db *mgo.Database, from, to time.Time) ([]Item, error) {
	articles := []Article{}
	err := db.C(PubMedItemCollectionName).Find(bson.M{
		"$and": []bson.M{
			bson.M{"created_at": bson.M{
				"$gte": from,
				"$lt":  to,
			}},
			bson.M{"feed_id": query.ID},
		},
	}).All(&articles)
	if err != nil {
		return []Item{}, err
	}

	items := []Item{}
	for _, art := range articles {
		items = append(items, art)
	}
	return items, nil
}
