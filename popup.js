
// newScripture
var counter = 0
function handleClick() {
    var bkg = chrome.extension.getBackgroundPage();
    counter++
    bkg.console.log('fooy-gooy :: ' + counter)
}
document.getElementById("newScripture").addEventListener("click",handleClick);

// Test the JQuery with #appendTest
function appendTest() {
    $("body").append("Test!!");
}
document.getElementById("appendTest").addEventListener("click", appendTest);
