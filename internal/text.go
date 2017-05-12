package internal

import (
	"strconv"
	"strings"
)

func Tokenize(raw string) []string {
	tokens := []string{}
	repl := strings.NewReplacer(".", "", ",", "", ":", "", ";", "", "<p>", "", "</p>", "")
	str := repl.Replace(raw)
	for _, wd := range strings.Split(str, " ") {
		wd = strings.ToLower(wd)
		skip := false

		wd = strings.TrimSpace(wd)

		if len(wd) < 2 {
			skip = true
		} else if _, err := strconv.Atoi(wd); err == nil {
			skip = true
		} else {
			_, found := skipWords[wd]
			if found {
				skip = true
				skipWords[wd]++
			}
		}

		wd = strings.Replace(wd, "(", "", -1)
		wd = strings.Replace(wd, ")", "", -1)

		if !skip {
			tokens = append(tokens, wd)
		}
	}

	return tokens
}

var skipWords = map[string]int{
	"the":        0,
	"them":       0,
	"this":       0,
	"i":          0,
	"a":          0,
	"all":        0,
	"or":         0,
	"it":         0,
	"was":        0,
	"an":         0,
	"and":        0,
	"are":        0,
	"as":         0,
	"which":      0,
	"be":         0,
	"been":       0,
	"but":        0,
	"by":         0,
	"may":        0,
	"for":        0,
	"did":        0,
	"we":         0,
	"methods":    0,
	"conclusion": 0,
	"about":      0,
	"has":        0,
	"from":       0,
	"have":       0,
	"in":         0,
	"more":       0,
	"there":      0,
	"our":        0,
	"to":         0,
	"of":         0,
	"that":       0,
	"these":      0,
	"being":      0,
	"on":         0,
	"than":       0,
	"with":       0,
	"is":         0,
	"its":        0,
	"<p>":        0,
	"</p>":       0,
	"do":         0,
	"not":        0,
	"also":       0,
	"one":        0,
	"two":        0,
	"three":      0,
	"four":       0,
	"five":       0,
	"six":        0,
	"seven":      0,
	"eight":      0,
	"nine":       0,
	"ten":        0,
	"eleven":     0,
	"twelve":     0,
	"thirteen":   0,
	"fourtieen":  0,
	"fifteen":    0,
	"sixteen":    0,
	"seventeen":  0,
	"eighteen":   0,
	"nineteen":   0,
	"most":       0,
	"now":        0,
	"who":        0,
	"each":       0,
	"at":         0,
	"shows":      0,
	"can":        0,
	"amount":     0,
	"were":       0,
	"after":      0,
	"into":       0,
	"even":       0,
	"often":      0,
	"those":      0,
	"both":       0,
}
