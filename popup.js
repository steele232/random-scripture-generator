
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
var __ScriptureRefs__

function loadJSON(callback) {

    var xobj = new XMLHttpRequest();
    xobj.overrideMimeType("application/json");
    xobj.open('GET', 'verses.json', true); // DONE // Replace 'my_data' with the path to your file
    xobj.onreadystatechange = function () {
          if (xobj.readyState == 4 && xobj.status == "200") {
            // Required use of an anonymous callback as .open will NOT return a value but simply returns undefined in asynchronous mode
            callback(xobj.responseText);
          }
    };
    xobj.send(null);  
}

function assignToScriptureRefs(str) {
    // Parse JSON string into object
    var __ScriptureRefs__ = JSON.parse(str);
}

function loadScriptureRefs() {
    loadJSON(
        assignToScriptureRefs
    )
}
document.addEventListener("load", loadScriptureRefs)

function logBool(b) {
    var bkg = chrome.extension.getBackgroundPage();
    bkg.console.log(b)
}
function generateNewScripture() {
    loadScriptureRefs()
    var bkg = chrome.extension.getBackgroundPage();
    bkg.console.log(__ScriptureRefs__)
    bkg.console.log( chrome.extension.isAllowedFileSchemeAccess(logBool) )
}
document.getElementById("newScripture").addEventListener("click",generateNewScripture);

