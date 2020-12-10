#! /usr/bin/osascript
// https://developer.apple.com/library/archive/documentation/LanguagesUtilities/Conceptual/MacAutomationScriptingGuide/PromptforText.html#//apple_ref/doc/uid/TP40016239-CH80-SW1
function run(){
    var app = Application.currentApplication()
    app.includeStandardAdditions = true
    var response = app.displayDialog(
        "Enter your Yubikey passphrase",
        {
            defaultAnswer: "",
            withIcon: "stop",
            buttons: ["Cancel", "Save"],
            defaultButton: "Save",
            cancelButton: "Cancel",
            givingUpAfter: 120,
            hiddenAnswer: true
        })
    return response.textReturned
}
