package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Movie struct {
	Title string
	Info  map[string]string
}

func main() {
	c := colly.NewCollector()

	file, err := os.Create("movie_data.csv")
	if err != nil {
		log.Fatalf("Could not create CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var movieInfo Movie
	movieInfo.Info = make(map[string]string)

	c.OnHTML("h1.unset > rt-text", func(e *colly.HTMLElement) {
		title := e.Text
		movieInfo.Title = title
		fmt.Println("Title:", title)
		writer.Write([]string{"Title", title}) // Write title with key "Title" to CSV
	})

	c.OnHTML(".media-info dl", func(e *colly.HTMLElement) {
		e.ForEach(".category-wrap", func(_ int, el *colly.HTMLElement) {
			key := el.ChildText("dt")
			value := el.ChildText("dd")
			writer.Write([]string{key, value})
			fmt.Printf("%s: %s\n", key, value)
			movieInfo.Info[key] = value
		})
	})

	c.OnHTML(".synopsis-wrap", func(e *colly.HTMLElement) {
		key := e.ChildText(".key")
		synopsis := e.ChildText("rt-text:not(.key)")
		writer.Write([]string{key, synopsis})
		fmt.Printf("%s: %s\n", key, synopsis)
		movieInfo.Info[key] = synopsis
	})

	movieURL := "https://www.rottentomatoes.com/m/the_good_the_bad_and_the_ugly"
	err = c.Visit(movieURL)
	if err != nil {
		log.Fatalf("Error visiting URL: %v", err)
	}

	fmt.Println("Scraping finished and data saved to movie_data.csv")
}
