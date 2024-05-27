package main

import (
	"HiddenEyeGoEdition/localization"
	"fmt"
	"os"
)

func main() {
	prereqs()
	mainRunner()
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func prereqs() {
	fmt.Println(localization.TerminalClear)
	fmt.Println(localization.Title)
	fmt.Println(localization.TitleDash)
	fmt.Println()
	fmt.Println(localization.EulaApproval)
	var EulaPerms string
	fmt.Scanln(&EulaPerms)
	fmt.Println(localization.TerminalClear)

	switch EulaPerms {
	case "Y", "y":
		checkExecutables()
	case "N", "n":
		fmt.Println(localization.EulaDenied)
		os.Exit(0)
	default:
		prereqs() // Recursively call prereqs() for invalid input
	}
}

func checkExecutables() {
	boreExists := fileExists("src/bore.exe")
	phpExists := fileExists("src/php/php.exe")

	switch {
	case boreExists && phpExists:
		// Both executables exist, nothing to do here
	case !boreExists && phpExists:
		fmt.Println(localization.MissingBore)
	case boreExists && !phpExists:
		fmt.Println(localization.MissingPhp)
	case !boreExists && !phpExists:
		fmt.Println(localization.BigOopsie)
		os.Exit(0)
	}
}

func mainRunner() {
	fmt.Println(localization.Title)
	fmt.Println(localization.TitleDash)
	fmt.Println()
	fmt.Println(localization.PageList)
}
