package main

const ScripturesPrefix = "scriptures"

const langParam = "lang=eng"

const BookOfMormon = "bofm"
const DCTestament = "dc-testament"
const PearlGPrice = "pgp"
const NewTestament = "nt"
const OldTestament = "ot"

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
	numChaptersMap["ps"] = 4
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
