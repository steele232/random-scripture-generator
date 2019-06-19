package main

// Old Testament ("ot")
var OTBooks = []string{
	"gen",
	"ex",
	"lev",
	"num",
	"deut",
	"josh",
	"judg",
	"ruth",
	"1-sam",
	"2-sam",
	"1-kgs",
	"2-kgs",
	"1-chr",
	"2-chr",
	"ezra",
	"neh",
	"esth",
	"job",
	"ps",
	"prov",
	"eccl",
	"song",
	"isa", //isaiah
	"jer",
	"lam",
	"ezek",
	"dan",
	"hosea",
	"joel",
	"amos",
	"obad",
	"jonah",
	"micah",
	"nahum",
	"hab",
	"zeph",
	"hag",
	"zech",
	"mal",
}

// New Testament ("nt")
var NTBooks = []string{
	"matt",
	"mark",
	"luke",
	"john",
	"acts",
	"romans",
	"1-cor",
	"2-cor",
	"gal",
	"eph",
	"philip",
	"col",
	"1-thes",
	"2-thes",
	"1-tim",
	"2-tim",
	"titus",
	"philem",
	"heb",
	"james",
	"1-pet",
	"2-pet",
	"1-jn",
	"2-jn",
	"3-jn",
	"jude",
	"rev",
}

// Book of Mormon ("bofm")
// ?? skipped the introduction?
var BofMBooks = []string{

	"1-ne",
	"2-ne",
	"jacob",
	"enos",
	"jarom",
	"omni",
	"w-of-m",
	"mosiah",
	"alma",
	"hel",
	"3-ne",
	"4-ne",
	"morm",
	"ether",
	"moro",
}

// Doctrine and Covenants ("dc-testament")
// ?? skipped the introduction?
var DCBooks = []string{
	"dc", // example: https://www.churchofjesuschrist.org/scriptures/dc-testament/dc/1?lang=eng
	"od", // 2 official declarations. example: https://www.churchofjesuschrist.org/scriptures/dc-testament/od/1?lang=eng
}

// Pearl of Great Price ("pgp")
// ?? skipped the introduction?
var PGPBooks = []string{
	"moses",
	"abr",  // ignoring fac-1, fac-2, fac-3
	"js-m", // ?? this could be tricky with verses
	"js-h",
	"a-of-f",
}
