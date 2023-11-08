package main

import (
	"github.com/gocolly/colly/v2"
)

func getWordDefinition(word string) string {
	c := colly.NewCollector()

	var definition string

	c.OnHTML("div[class=def-content]", func(e *colly.HTMLElement) {
		definition = e.Text
	})

	c.Visit("https://www.merriam-webster.com/dictionary/" + word)

	return definition
}