package pubmed

import (
	"errors"
	"io"
	"time"
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
	return nil
}

func (article *Article) Read(p []byte) (n int, err error) {
	if article.done {
		return 0, io.EOF
	}
	read := []byte(article.ArticleTitle + article.AbstractText)[article.readIndex:]
	n = copy(p, read)
	article.readIndex += n
	return
}

func (article *Article) ResetReader() {
	article.done = false
	article.readIndex = 0
}
