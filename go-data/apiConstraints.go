package main

import (
	"strconv"
)

const ScripturesPrefix = "scriptures"

const langParam = "lang=eng"

const BookOfMormon = "bofm"
const DCTestament = "dc-testament"
const PearlGPrice = "pgp"
const NewTestament = "nt"
const OldTestament = "ot"

func humanizedBookNameMap() map[string]string {
	bookNameMap := make(map[string]string)

	// Old Testament ("ot")
	bookNameMap["gen"] = "Genesis"
	bookNameMap["ex"] = "Exodus"
	bookNameMap["lev"] = "Leviticus"
	bookNameMap["num"] = "Numbers"
	bookNameMap["deut"] = "Deuteronomy"
	bookNameMap["josh"] = "Joshua"
	bookNameMap["judg"] = "Judges"
	bookNameMap["ruth"] = "Ruth"
	bookNameMap["1-sam"] = "1 Samuel"
	bookNameMap["2-sam"] = "2 Samuel"
	bookNameMap["1-kgs"] = "1 Kings"
	bookNameMap["2-kgs"] = "2 Kings"
	bookNameMap["1-chr"] = "1 Chronicles"
	bookNameMap["2-chr"] = "2 Chronicles"
	bookNameMap["ezra"] = "Ezra"
	bookNameMap["neh"] = "Nehemiah"
	bookNameMap["esth"] = "Esther"
	bookNameMap["job"] = "Job"
	bookNameMap["ps"] = "Psalms"
	bookNameMap["prov"] = "Proverbs"
	bookNameMap["eccl"] = "Ecclesiastes"
	bookNameMap["song"] = "Song of Solomon"
	bookNameMap["isa"] = "Isaiah"
	bookNameMap["jer"] = "Jeremiah"
	bookNameMap["lam"] = "Lamentations"
	bookNameMap["ezek"] = "Ezekiel"
	bookNameMap["dan"] = "Daniel"
	bookNameMap["hosea"] = "Hosea"
	bookNameMap["joel"] = "Joel"
	bookNameMap["amos"] = "Amos"
	bookNameMap["obad"] = "Obadiah"
	bookNameMap["jonah"] = "Jonah"
	bookNameMap["micah"] = "Micah"
	bookNameMap["nahum"] = "Nahum"
	bookNameMap["hab"] = "Habakkuk"
	bookNameMap["zeph"] = "Zephaniah"
	bookNameMap["hag"] = "Haggai"
	bookNameMap["zech"] = "Zechariah"
	bookNameMap["mal"] = "Malachi"

	// New Testament ("nt")
	bookNameMap["matt"] = "Matthew"
	bookNameMap["mark"] = "Mark"
	bookNameMap["luke"] = "Luke"
	bookNameMap["john"] = "John"
	bookNameMap["acts"] = "Acts"
	bookNameMap["romans"] = "Romans"
	bookNameMap["1-cor"] = "1 Corinthians"
	bookNameMap["2-cor"] = "2 Corinthians"
	bookNameMap["gal"] = "Galatians"
	bookNameMap["eph"] = "Ephesians"
	bookNameMap["philip"] = "Philippians"
	bookNameMap["col"] = "Collosians"
	bookNameMap["1-thes"] = "1 Thessalonians"
	bookNameMap["2-thes"] = "2 Thessalonians"
	bookNameMap["1-tim"] = "1 Timothy"
	bookNameMap["2-tim"] = "2 Timothy"
	bookNameMap["titus"] = "Titus"
	bookNameMap["philem"] = "Philemon"
	bookNameMap["heb"] = "Hebrews"
	bookNameMap["james"] = "James"
	bookNameMap["1-pet"] = "1 Peter"
	bookNameMap["2-pet"] = "2 Peter"
	bookNameMap["1-jn"] = "1 John"
	bookNameMap["2-jn"] = "2 John"
	bookNameMap["3-jn"] = "3 John"
	bookNameMap["jude"] = "Jude"
	bookNameMap["rev"] = "Revelation"

	// Book of Mormon ("bofm")
	// ?? skipped the introduction?
	bookNameMap["1-ne"] = "1 Nephi"
	bookNameMap["2-ne"] = "2 Nephi"
	bookNameMap["jacob"] = "Jacob"
	bookNameMap["enos"] = "Enos"
	bookNameMap["jarom"] = "Jarom"
	bookNameMap["omni"] = "Omni"
	bookNameMap["w-of-m"] = "Words of Mormon"
	bookNameMap["mosiah"] = "Mosiah"
	bookNameMap["alma"] = "Alma"
	bookNameMap["hel"] = "Helaman"
	bookNameMap["3-ne"] = "3 Nephi"
	bookNameMap["4-ne"] = "4 Nephi"
	bookNameMap["morm"] = "Mormon"
	bookNameMap["ether"] = "Ether"
	bookNameMap["moro"] = "Moroni"

	// Doctrine and Covenants ("dc-testament")
	// ?? skipped the introduction?
	bookNameMap["dc"] = "D&C"
	bookNameMap["od"] = "Official Declaration"

	// Pearl of Great Price ("pgp")
	// ?? skipped the introduction?
	bookNameMap["moses"] = "Moses"
	bookNameMap["abr"] = "Abraham"      // ignoring fac-1, fac-2, fac-3
	bookNameMap["js-m"] = "JST-Matthew" // ?? this could be tricky with verses
	bookNameMap["js-h"] = "Joseph Smith History"
	bookNameMap["a-of-f"] = "Articles of Faith"

	return bookNameMap
}

func humanizedBookAndChapterName(bookCode string, chapterNum int) string {
	bookMap := humanizedBookNameMap()
	return bookMap[bookCode] + " " + strconv.Itoa(chapterNum)
}

func bookNumChaptersMap() map[string]int {

	numChaptersMap := make(map[string]int)

	// Old Testament ("ot")
	numChaptersMap["gen"] = 50
	numChaptersMap["ex"] = 40
	numChaptersMap["lev"] = 27
	numChaptersMap["num"] = 36
	numChaptersMap["deut"] = 34
	numChaptersMap["josh"] = 24
	numChaptersMap["judg"] = 21
	numChaptersMap["ruth"] = 4
	numChaptersMap["1-sam"] = 31
	numChaptersMap["2-sam"] = 24
	numChaptersMap["1-kgs"] = 22
	numChaptersMap["2-kgs"] = 25
	numChaptersMap["1-chr"] = 29
	numChaptersMap["2-chr"] = 36
	numChaptersMap["ezra"] = 10
	numChaptersMap["neh"] = 13
	numChaptersMap["esth"] = 10
	numChaptersMap["job"] = 42
	numChaptersMap["ps"] = 150
	numChaptersMap["prov"] = 31
	numChaptersMap["eccl"] = 12
	numChaptersMap["song"] = 8
	numChaptersMap["isa"] = 66 //isaiah
	numChaptersMap["jer"] = 52
	numChaptersMap["lam"] = 5
	numChaptersMap["ezek"] = 48
	numChaptersMap["dan"] = 12
	numChaptersMap["hosea"] = 14
	numChaptersMap["joel"] = 3
	numChaptersMap["amos"] = 9
	numChaptersMap["obad"] = 1
	numChaptersMap["jonah"] = 4
	numChaptersMap["micah"] = 7
	numChaptersMap["nahum"] = 3
	numChaptersMap["hab"] = 3
	numChaptersMap["zeph"] = 3
	numChaptersMap["hag"] = 2
	numChaptersMap["zech"] = 14
	numChaptersMap["mal"] = 4

	// New Testament ("nt")
	numChaptersMap["matt"] = 28
	numChaptersMap["mark"] = 16
	numChaptersMap["luke"] = 24
	numChaptersMap["john"] = 21
	numChaptersMap["acts"] = 28
	numChaptersMap["romans"] = 16
	numChaptersMap["1-cor"] = 16
	numChaptersMap["2-cor"] = 13
	numChaptersMap["gal"] = 6
	numChaptersMap["eph"] = 6
	numChaptersMap["philip"] = 4
	numChaptersMap["col"] = 4
	numChaptersMap["1-thes"] = 5
	numChaptersMap["2-thes"] = 3
	numChaptersMap["1-tim"] = 6
	numChaptersMap["2-tim"] = 4
	numChaptersMap["titus"] = 3
	numChaptersMap["philem"] = 1
	numChaptersMap["heb"] = 13
	numChaptersMap["james"] = 5
	numChaptersMap["1-pet"] = 5
	numChaptersMap["2-pet"] = 3
	numChaptersMap["1-jn"] = 5
	numChaptersMap["2-jn"] = 1
	numChaptersMap["3-jn"] = 1
	numChaptersMap["jude"] = 1
	numChaptersMap["rev"] = 22

	// Book of Mormon ("bofm")
	// ?? skipped the introduction?
	numChaptersMap["1-ne"] = 22
	numChaptersMap["2-ne"] = 33
	numChaptersMap["jacob"] = 7
	numChaptersMap["enos"] = 1
	numChaptersMap["jarom"] = 1
	numChaptersMap["omni"] = 1
	numChaptersMap["w-of-m"] = 1
	numChaptersMap["mosiah"] = 29
	numChaptersMap["alma"] = 63
	numChaptersMap["hel"] = 16
	numChaptersMap["3-ne"] = 30
	numChaptersMap["4-ne"] = 1
	numChaptersMap["morm"] = 9
	numChaptersMap["ether"] = 15
	numChaptersMap["moro"] = 10

	// Doctrine and Covenants ("dc-testament")
	// ?? skipped the introduction?
	numChaptersMap["dc"] = 138 // example: https://www.lds.org/scriptures/dc-testament/dc/1?lang=eng
	numChaptersMap["od"] = 2   // 2 official declarations. example: https://www.lds.org/scriptures/dc-testament/od/1?lang=eng

	// Pearl of Great Price ("pgp")
	// ?? skipped the introduction?
	numChaptersMap["moses"] = 8
	numChaptersMap["abr"] = 5  // ignoring fac-1, fac-2, fac-3
	numChaptersMap["js-m"] = 1 // ?? this could be tricky with verses
	numChaptersMap["js-h"] = 1
	numChaptersMap["a-of-f"] = 1

	return numChaptersMap
}
