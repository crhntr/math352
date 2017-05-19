package pubmed

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

func (date *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type dateXML struct {
		Year  int
		Month int
		Day   int
	}
	var (
		err  error
		dxml dateXML
	)

	if err = d.DecodeElement(&dxml, &start); err != nil {
		return err
	}

	*date = Date{time.Date(dxml.Year, time.Month(dxml.Month), dxml.Day, 0, 0, 0, 0, time.UTC)}
	return nil
}

func (date Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type dateXML struct {
		Year  string
		Month string
		Day   string
	}
	return e.EncodeElement(&dateXML{
		Year:  fmt.Sprintf("%0.4d", date.Year()),
		Month: fmt.Sprintf("%0.2d", int(date.Month())),
		Day:   fmt.Sprintf("%0.2d", date.Day()),
	}, start)
}
