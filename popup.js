
/*================================*/
// Load in the Scripture References
/*================================*/
var __ScriptureRefs__ = __verses__
var volume

/*================================*/
// Get a Random Number!
/*================================*/
function getRandomInt(max) {
    return Math.floor(Math.random() * Math.floor(max));
}
  
// console.log(getRandomInt(3));
// // expected output: 0, 1 or 2

function logBool(b) {
    var bkg = chrome.extension.getBackgroundPage();
    bkg.console.log(b)
}
function generateNewScripture() {
    
    var bkg = chrome.extension.getBackgroundPage();

    let numVerses = __ScriptureRefs__.data.length
    bkg.console.log("Number of Verses :: " + numVerses)
    var randVerseIdx = getRandomInt(numVerses)
    bkg.console.log("Show the random verse IDX!! :: " + randVerseIdx)
    let verseRefObj = __ScriptureRefs__.data[randVerseIdx]
    linkElm = document.getElementById("verse-name")
    linkElm.innerText = verseRefObj.humanName
    linkElm.href = verseRefObj.url
    contentElm = document.getElementById("verse-content")
    contentElm.innerHTML = verseRefObj.content
    bookVolume = document.getElementById("book-volume")
    bookVolume.innerText = determineBook(verseRefObj.url)

    // SAVE the newly generated scripture as the lastScripture
    chrome.storage.local.set({"lastScripture": randVerseIdx}, (items) => {

        var bkg = chrome.extension.getBackgroundPage();
        bkg.console.log("Saved as LASTSCRIPTURE:: " + randVerseIdx)

    });

}
document.getElementById("newScripture").addEventListener("click",generateNewScripture);

/* Load Scriptures */ 
function loadInScripture(scriptureIdx) {
    
    let verseRefObj = __ScriptureRefs__.data[scriptureIdx]
    linkElm = document.getElementById("verse-name")
    linkElm.innerText = verseRefObj.humanName
    linkElm.href = verseRefObj.url
    contentElm = document.getElementById("verse-content")
    contentElm.innerHTML = verseRefObj.content
    bookVolume = document.getElementById("book-volume")
    bookVolume.innerText = determineBook(verseRefObj.url)
}

function determineBook(url) {
    console.log(url)
    if(url.includes("/bofm/")) {
        return "Book of Mormon"
    } else if(url.includes("/dc-testament/")) {
        return "Doctrine and Covenants"
    } else if(url.includes("/pgp/")) {
        return "Pearl of Great Price"
    } else if (url.includes("/nt/")) {
        return "New Testament"
    } else if (url.includes("/ot/")) {
        return "Old Testament"
    }
}

// Local storage && Load in scripture on Startup
var LASTSCRIPTURE = 0;
window.onload = () => {
    chrome.storage.local.get(["lastScripture"], (items) => {
        // Create a variable if one does not exist
        if (!items["lastScripture"]) {
            LASTSCRIPTURE = 0;
        }

        // Reassign or manipulate data
        else {
            LASTSCRIPTURE = items.lastScripture
        }
        loadInScripture(LASTSCRIPTURE)

        var bkg = chrome.extension.getBackgroundPage();
        bkg.console.log("Loaded in LASTSCRIPTURE:: " + LASTSCRIPTURE)

    });
};
