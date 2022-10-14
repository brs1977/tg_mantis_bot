package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func readHtmlFromFile(fileName string) (string, error) {

	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func TestParser(t *testing.T) {

	fileName := "test/test.html"
	text, err := readHtmlFromFile(fileName)

	if err != nil {
		log.Fatal(err)
	}
	var message = parseTable(text)
	var got = "alekseev РУСЗН задание(2) \nbabykina ИС МЦ ТСР задание(1) \nsavchenko РУСЗН задание(1) \n"

	if got != message {
		t.Errorf("message = ['%s']; want ['%s']", message, got)
	}
}
