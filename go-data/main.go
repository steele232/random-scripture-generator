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

type Pair struct {
	Left  string
	Right int
}

func main() {

	printDescriptiveStatistics()

	allURLs := make([]string, 0)

	// m := bookNumChaptersMap()
	for _, bookCode := range BofMBooks {
		var url string
		maxChapterNum := bookNumChaptersMap()[bookCode]
		for i := 1; i <= maxChapterNum; i++ {
			url = buildURL(BookOfMormon, bookCode, i)
			allURLs = append(allURLs, url)
			// fmt.Println("Url: ", url)
		}
	}
	for _, bookCode := range DCBooks {
		var url string
		maxChapterNum := bookNumChaptersMap()[bookCode]
		for i := 1; i <= maxChapterNum; i++ {
			url = buildURL(DCTestament, bookCode, i)
			allURLs = append(allURLs, url)
			// fmt.Println("Url: ", url)
		}
	}
	for _, bookCode := range NTBooks {
		var url string
		maxChapterNum := bookNumChaptersMap()[bookCode]
		for i := 1; i <= maxChapterNum; i++ {
			url = buildURL(NewTestament, bookCode, i)
			allURLs = append(allURLs, url)
			// fmt.Println("Url: ", url)
		}
	}
	for _, bookCode := range OTBooks {
		var url string
		maxChapterNum := bookNumChaptersMap()[bookCode]
		for i := 1; i <= maxChapterNum; i++ {
			url = buildURL(OldTestament, bookCode, i)
			allURLs = append(allURLs, url)
			// fmt.Println("Url: ", url)
		}
	}
	for _, bookCode := range PGPBooks {
		var url string
		maxChapterNum := bookNumChaptersMap()[bookCode]
		for i := 1; i <= maxChapterNum; i++ {
			url = buildURL(PearlGPrice, bookCode, i)
			allURLs = append(allURLs, url)
			// fmt.Println("Url: ", url)
		}
	}

	fmt.Println("Collected all Chapter URLs :: ", len(allURLs))

	pairs := make([]Pair, 0)
	for _, url := range allURLs[:20] {

		htmlBody := getHTMLfromURL(url)
		bodyReader := strings.NewReader(htmlBody)
		document, err := goquery.NewDocumentFromReader(bodyReader)
		if err != nil {
			log.Println("Could not make a document from url: ", url, err)
		}

		numVerses := len(document.Find("p .verse").Nodes)
		fmt.Println("URL : ", url, " ==> ", numVerses)

		for i := 1; i <= numVerses; i++ {
			newURL := insertVerseNumIntoURL(url, i)
			pairs = append(pairs, Pair{Left: newURL, Right: i})
		}

	}
	fmt.Println("**NumPairs: ", len(pairs))
	fmt.Println("SAMPLE PAIRS:: ", pairs[:30])

}

func insertVerseNumIntoURL(url string, verseNum int) string {

	sArr := strings.Split(url, "?")
	return sArr[0] + "." + strconv.Itoa(verseNum) + "?" + sArr[1]
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
