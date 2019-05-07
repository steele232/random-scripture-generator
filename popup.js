
// fooey
var counter = 0
function handleClick() {
    var bkg = chrome.extension.getBackgroundPage();
    counter++
    bkg.console.log('fooy-gooy :: ' + counter)
}
document.getElementById("fooey").addEventListener("click",handleClick);

// Test the JQuery with #appendTest
function appendTest() {
    $("body").append("Test!!");
}
document.getElementById("appendTest").addEventListener("click", appendTest);

/*================================*/
// Load in the Scripture References
/*================================*/
var __ScriptureRefs__ = JSON.parse(__verses__);

// function loadJSON(callback) {

//     var xobj = new XMLHttpRequest();
//     xobj.overrideMimeType("application/json");
//     xobj.open('GET', 'verses.json', true); // DONE // Replace 'my_data' with the path to your file
//     xobj.onreadystatechange = function () {
//           if (xobj.readyState == 4 && xobj.status == "200") {
//             // Required use of an anonymous callback as .open will NOT return a value but simply returns undefined in asynchronous mode
//             callback(xobj.responseText);
//           }
//     };
//     xobj.send(null);  
// }

// function assignToScriptureRefs(str) {
//     // Parse JSON string into object
//     var __ScriptureRefs__ = JSON.parse(str);
// }

// function loadScriptureRefs() {
//     loadJSON(
//         assignToScriptureRefs
//     )
// }
// document.addEventListener("load", loadScriptureRefs)

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
    chrome.tabs.create({url:verseRefObj.url+"#p8"})

}
document.getElementById("newScripture").addEventListener("click",generateNewScripture);

