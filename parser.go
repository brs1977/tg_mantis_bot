package main

import (
	"strings"

	"golang.org/x/net/html"
)

func parseTable(text string) string {

	tkn := html.NewTokenizer(strings.NewReader(text))

	var isTd bool
	var isB bool
	var bText string
	var message string = ""
	var n int

	for {

		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return message

		case tt == html.StartTagToken:

			t := tkn.Token()
			isB = t.Data == "b"
			isTd = t.Data == "td"

		case tt == html.TextToken:

			t := tkn.Token()

			if isB {
				bText = t.Data
			}

			if bText == "Открытые:" && isTd {
				message += t.Data + " "
				n++
			}

			if bText == "Открытые:" && isTd && n%3 == 0 {
				message += "\n"
			}

			isTd = false
		}
	}
}
