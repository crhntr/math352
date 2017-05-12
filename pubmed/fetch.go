package pubmed

import (
	"encoding/xml"
	"errors"
	"fmt"
	"time"

	"net/http"
	"net/url"
	"strings"

	. "github.com/crhntr/math352/internal"
)

const dateFmt = "2006/01/02"

func (q Query) FetchItemsForDay(d time.Time) ([]Item, error) {
	day := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
	var (
		items          = []Item{}
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

	err = xml.NewDecoder(r.Body).Decode(&articlesSet)
	if err != nil {
		return items, err
	}

	for _, article := range articlesSet.Articles {
		article.FeedID = q.ID
		items = append(items, article)
	}

	return items, err
}
