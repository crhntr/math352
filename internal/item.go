package internal

import (
	"io"
	"time"
)

type Item interface {
	Title() string
	Body() string
	Published() time.Time
	Added() time.Time
	io.Reader
}
