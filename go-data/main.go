package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ScriptureRef struct {
	Name    string `json:"humanName"`
	Url     string `json:"url"`
	Content string `json:"content"`
}

type JsonOutput struct {
	Verses []ScriptureRef `json:"verses"`
}

func main() {

	printDescriptiveStatistics()

	allChapterRefs := collectAllChapterURLs()
	fmt.Println("Collected all Chapter URLs :: ", len(allChapterRefs))

	allVerseRefs := make([]ScriptureRef, 0)
	for idx, chapterRef := range allChapterRefs {

		// if idx > 10 {
		// 	break
		// }

		htmlBody := getHTMLfromURL(chapterRef.Url)
		bodyReader := strings.NewReader(htmlBody)
		document, err := goquery.NewDocumentFromReader(bodyReader)
		if err != nil {
			log.Println("Could not make a document from url: ", chapterRef.Url, err)
		}

		numVerses := len(document.Find(".verse").Nodes)
		fmt.Println("URL : ", chapterRef.Url, " ==> ", numVerses, " :: @Chapter-IDX ", idx)

		for i := 1; i <= numVerses; i++ {

			// Get ready to use GoQuery
			verseNum := i
			verseSelector := "#p" + strconv.Itoa(verseNum)
			verseRef := document.Find(verseSelector)

			// Remove the Footnotes
			verseRef.Find(".marker").Each(func(idx int, elem *goquery.Selection) {
				elem.Empty()
			})

			// Save the ScriptureRef object
			verseContent := verseRef.Text()
			newName := chapterRef.Name + ":" + strconv.Itoa(verseNum)
			newURL := insertVerseNumIntoURL(chapterRef.Url, i)
			newRef := ScriptureRef{
				Name:    newName,
				Url:     newURL,
				Content: verseContent,
			}
			allVerseRefs = append(allVerseRefs, newRef)
		}

	}
	fmt.Println("Number of Verses: ", len(allVerseRefs))
	fmt.Println("SAMPLE VERSES:: ", allVerseRefs[:30])

	/* Write to a file */

	// open the file
	f, err := os.Create("verses.json")
	if err != nil {
		log.Println("Could not write to file verses.json: ", err)
	}
	defer f.Close()

	// make a writer from the file
	w := bufio.NewWriter(f)

	// write it to a file.
	enc := json.NewEncoder(w)
	enc.Encode(allVerseRefs)

	w.Flush()

}

func collectAllChapterURLs() []ScriptureRef {

	allChapterRefs := make([]ScriptureRef, 0)

	for _, bookCode := range BofMBooks {
		maxChapterNum := bookNumChaptersMap()[bookCode]
		for i := 1; i <= maxChapterNum; i++ {
			chapterRef := ScriptureRef{
				Name: humanizedBookAndChapterName(bookCode, i), // (e.g. D&C 10)
				Url:  buildURL(BookOfMormon, bookCode, i),
			}
			allChapterRefs = append(allChapterRefs, chapterRef)
		}
	}
	for _, bookCode := range DCBooks {
		maxChapterNum := bookNumChaptersMap()[bookCode]
		for i := 1; i <= maxChapterNum; i++ {
			chapterRef := ScriptureRef{
				Name: humanizedBookAndChapterName(bookCode, i), // (e.g. D&C 10)
				Url:  buildURL(DCTestament, bookCode, i),
			}
			allChapterRefs = append(allChapterRefs, chapterRef)
		}
	}
	for _, bookCode := range NTBooks {
		maxChapterNum := bookNumChaptersMap()[bookCode]
		for i := 1; i <= maxChapterNum; i++ {
			chapterRef := ScriptureRef{
				Name: humanizedBookAndChapterName(bookCode, i), // (e.g. D&C 10)
				Url:  buildURL(NewTestament, bookCode, i),
			}
			allChapterRefs = append(allChapterRefs, chapterRef)
		}
	}
	for _, bookCode := range OTBooks {
		maxChapterNum := bookNumChaptersMap()[bookCode]
		for i := 1; i <= maxChapterNum; i++ {
			chapterRef := ScriptureRef{
				Name: humanizedBookAndChapterName(bookCode, i), // (e.g. D&C 10)
				Url:  buildURL(OldTestament, bookCode, i),
			}
			allChapterRefs = append(allChapterRefs, chapterRef)
		}
	}
	for _, bookCode := range PGPBooks {
		maxChapterNum := bookNumChaptersMap()[bookCode]
		for i := 1; i <= maxChapterNum; i++ {
			chapterRef := ScriptureRef{
				Name: humanizedBookAndChapterName(bookCode, i), // (e.g. D&C 10)
				Url:  buildURL(PearlGPrice, bookCode, i),
			}
			allChapterRefs = append(allChapterRefs, chapterRef)
		}
	}

	return allChapterRefs
}

func insertVerseNumIntoURL(url string, verseNum int) string {

	sArr := strings.Split(url, "?")
	return sArr[0] + "." + strconv.Itoa(verseNum) + "?" + sArr[1] + "#p" + strconv.Itoa(verseNum)
}

// buildURL uses package-global variables from apiConstraints.go
func buildURL(stdWorkCode, bookCode string, chapterInteger int) string {
	var url string

	//let's build a URL really fast
	url = "https://www.churchofjesuschrist.org/" + ScripturesPrefix +
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
