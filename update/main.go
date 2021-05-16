package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
)

// script forked from https://github.com/victoriadrake/victoriadrake
func makeReadme(filename string) error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://andrebrandao.me/index.xml")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}
	// Get the freshest item
	blogItem := feed.Items[0]

	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}

	date := time.Now().Format("2 Jan 2006")

	hello := "### Hello! I’m André Brandão. 👋\n\nI’m a Software Engineer with a Bachelor's Degree in Computer Engineering from UNICAMP (Brazil) 🎓. Since then, I've been particularly interested in high-quality software, test-driven development and application performance."
	blog := "Check out my website **[andrebrandao.me](https://andrebrandao.me)** to find my latest article: **[" + blogItem.Title + "](" + blogItem.Link + ")**."
	updated := "<sub>Last updated on " + date + ".</sub>"
	data := fmt.Sprintf("%s\n\n%s\n\n%s\n", hello, blog, updated)

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func main() {

	makeReadme("../README.md")

}
