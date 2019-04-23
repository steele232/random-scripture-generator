// // This is the javascript that will be running constantly in the background

// // HOW TO CONSOLE LOG TO BACKGROUND PAGE
// function handleClick() {
//     var bkg = chrome.extension.getBackgroundPage();
//     bkg.console.log('fooy-gooy');
// }

// // Recieve messages from the content script
// chrome.runtime.onMessage.addListener((message) => {

//     // Send messages to the content script
//     alert(message);

//     chrome.tabs.getSelected(null, (tab) => {
//         chrome.tabs.sendMessage(tab.id, "From Background");
//     });
// });

// // Local storage
// window.onload = () => {
//     chrome.storage.local.get(["variable"], (items) => {
//         // Create a variable if one does not exist
//         if (!items["variable"]) {
//             variable = true;
//         }

//         // Reassign or manipulate data
//         else {
//             variable = !variable;
//         }
//     });
// };