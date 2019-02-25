package main

import "fmt"

func printDescriptiveStatistics() {

	m := bookNumChaptersMap()
	sumBooks := 0
	sumChapters := 0
	for _, v := range m {
		sumBooks++
		sumChapters += v
		// fmt.Println(k, "=", v)
	}
	fmt.Println("Total num of Books:", sumBooks)
	fmt.Println("Total num of Chapters:", sumChapters)

}
