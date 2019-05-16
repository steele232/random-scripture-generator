
/*================================*/
// Load in the Scripture References
/*================================*/
var __ScriptureRefs__ = JSON.parse(__verses__);

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
    // loadScriptureRefs()
    var bkg = chrome.extension.getBackgroundPage();
    let numVerses = __ScriptureRefs__.data.length
    bkg.console.log("Number of Verses :: " + numVerses)
    var randVerseIdx = getRandomInt(numVerses)
    bkg.console.log("Show the random verse IDX!! :: " + randVerseIdx)
    let verseRefObj = __ScriptureRefs__.data[randVerseIdx]
    linkElm = document.getElementById("verse-name")
    linkElm.innerText = verseRefObj.humanName
    linkElm.href = verseRefObj.url

    // let's try opening it in a new tab...
    // chrome.tabs.create({url:verseRefObj.url+"#p8"})

}
document.getElementById("newScripture").addEventListener("click",generateNewScripture);

