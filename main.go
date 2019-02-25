package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {

	printDescriptiveStatistics()

	// m := bookNumChaptersMap()
	for _, bookCode := range BofMBooks {
		var url string
		maxVerseNum := bookNumChaptersMap()[bookCode]
		for i := 1; i <= maxVerseNum; i++ {
			url = buildURL(BookOfMormon, bookCode, i)
			fmt.Println("Url: ", url)
		}
	}

	// TODO do the same and collect all URLs for DC, NT, OT, PGP

	// TODO use GoQuery to parse for the scriptures and find the number of verses.
	// https://github.com/PuerkitoBio/goquery

	// actually print the URL first.
	// fmt.Println(url)

	// htmlBody := getHTMLfromURL(url)

	// fmt.Println(htmlBody[:100])

}

// buildURL uses package-global variables from apiConstraints.go
func buildURL(stdWorkCode, bookCode string, chapterInteger int) string {
	var url string

	//let's build a URL really fast
	url = "https://www.lds.org/" + ScripturesPrefix +
		"/" + stdWorkCode + "/" + bookCode + "/" + strconv.Itoa(chapterInteger) +
		"?" + langParam

	return url
}

func getHTMLfromURL(url string) string {

	// send the URL and print the results..
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}
