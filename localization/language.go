package localization

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

// Very common text
var TerminalClear = "\033[H\033[2J"
var Title = Red +
	`
██   ██ ██ ██████   ██████   ███████ ███   ██  ███████ ██    ██ ███████ 
██   ██ ██ ██    ██ ██    ██ ██      ████  ██  ██       ██  ██  ██      
███████ ██ ██    ██ ██    ██ ███████ ██ ██ ██  ███████   ████   ███████ 
██   ██ ██ ██    ██ ██    ██ ██      ██  ████  ██         ██    ██      
██   ██ ██ ██████   ██████   ███████ ██   ███  ███████    ██    ███████ 
							    Reborn v0.1
` + Reset
var TitleDash = Blue + "-----------------------------------------------------------------------" + Reset

// Prereqs
var EulaApproval = Red + "Will you use this tool for educational purposes? (Y/N): " + Reset
var EulaDenied = Red + "This tool cannot be used for nefarious purposes." + Reset
var MissingBore = Red + "Bore is missing; please redownload Hiddeneye." + Reset
var MissingPhp = Red + "Php is missing; please redownload Hiddeneye." + Reset
var BigOopsie = Red + "Multiple dependicies are missing; please redownload Hiddeneye."

// MainRunner
var PageList = Red + "{1} Badoo" + Reset
var PageChooser = Red + "Choose a site: " + Reset

// Redirectererer
var RedirectPage = Red + "Redirect Link: " + Reset

// OutputPage
var LocalhostUrl = Red + "Localhost URL: localhost:8000" + Reset
