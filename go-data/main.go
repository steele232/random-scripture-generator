package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
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
	for _, bookCode := range DCBooks {
		var url string
		maxVerseNum := bookNumChaptersMap()[bookCode]
		for i := 1; i <= maxVerseNum; i++ {
			url = buildURL(DCTestament, bookCode, i)
			fmt.Println("Url: ", url)
		}
	}
	for _, bookCode := range NTBooks {
		var url string
		maxVerseNum := bookNumChaptersMap()[bookCode]
		for i := 1; i <= maxVerseNum; i++ {
			url = buildURL(NewTestament, bookCode, i)
			fmt.Println("Url: ", url)
		}
	}
	for _, bookCode := range OTBooks {
		var url string
		maxVerseNum := bookNumChaptersMap()[bookCode]
		for i := 1; i <= maxVerseNum; i++ {
			url = buildURL(OldTestament, bookCode, i)
			fmt.Println("Url: ", url)
		}
	}
	for _, bookCode := range PGPBooks {
		var url string
		maxVerseNum := bookNumChaptersMap()[bookCode]
		for i := 1; i <= maxVerseNum; i++ {
			url = buildURL(PearlGPrice, bookCode, i)
			fmt.Println("Url: ", url)
		}
	}

	// actually print the URL first.
	// fmt.Println(url)

	url := "https://www.lds.org/scriptures/pgp/a-of-f/1?lang=eng"
	htmlBody := getHTMLfromURL(url)

	fmt.Println(htmlBody[:300])

	// TODO use GoQuery to parse for the scriptures and find the number of verses.
	// https://github.com/PuerkitoBio/goquery
	bodyReader := strings.NewReader(htmlBody)
	document, err := goquery.NewDocumentFromReader(bodyReader)
	if err != nil {
		log.Println("Could not make a document from url: ", url, err)
	}

	fmt.Println("Print the cached results :) ::: ", len(document.Find("p .verse").Nodes))

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
