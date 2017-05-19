package pubmed

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const dateFmt = "2006/01/02"

func (q Query) FetchItemsForDay(d time.Time) ([]*Article, error) {
	day := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
	var (
		items          = []*Article{}
		eSearchResults = ESearchResult{}
		query          = q.Query
	)

	if len(query) < 1 {
		return items, errors.New("empty query")
	}
	if strings.Contains(query, "[Publication Date]") {
		return items, errors.New("Publication Date field should be set by litsphere")
	}
	query = strings.TrimLeft(query, "(")
	query = strings.TrimRight(query, ")")

	query = fmt.Sprintf(`((%q [Publication Date] : %q[Publication Date]) AND %s)`,
		day.Format(dateFmt), day.Format(dateFmt), query)
	query = url.QueryEscape(query)

	r, err := http.Get(pubMedMaxbaseURL + "esearch.fcgi?db=pubmed&term=" + query + "&usehistory=y")
	if err != nil {
		return items, err
	}

	dec := xml.NewDecoder(r.Body)
	err = dec.Decode(&eSearchResults)
	if err != nil {
		return items, err
	}
	if len(eSearchResults.IDs) == 0 {
		return items, err
	}

	idString := "&id=" + strings.Join(eSearchResults.IDs, ",")

	articlesSet := ArticleSet{}
	r, err = http.Get(pubMedMaxbaseURL + "efetch.fcgi?db=pubmed&rettype=abstract&retmode=xml" + idString)
	if err != nil {
		return items, err
	}

	// r, err := os.Open("static/testing.xml")
	// if err != nil {
	// 	return items, err
	// }
	// defer r.Close()
	err = xml.NewDecoder(r.Body).Decode(&articlesSet)
	if err != nil {
		return items, err
	}

	for _, article := range articlesSet.Articles {
		pub := article.Published()
		if pub.Year() == day.Year() && pub.Month() == day.Month() && pub.Day() == day.Day() {
			items = append(items, article)
		}
	}

	return items, err
}
