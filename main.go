package main

import (
	"delshared/helpers"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	args, err := helpers.GetOptions(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// Elevate Privileges to Administrator level
	if !helpers.AmIAdmin() {
		helpers.ElevateAsAdmin()
	}

	utility := "C:\\Middleware\\user_projects\\epmsystem1\\Planning\\planning1\\DeleteSharedDescendants.cmd"
	passFile := fmt.Sprintf("%q", "C:\\epmpi\\enc_pass")

	logFile := fmt.Sprintf("%s\\epmpi_%s.log", currentDir, time.Now().Format("20060102_150405"))
	fmt.Printf("\n\nLog file is located: %s\n\n", logFile)
	logFile = fmt.Sprintf("%q", logFile)

	var app, element string

	app = strings.TrimSpace(args[0])
	element = strings.TrimSpace(args[1])

	helpers.StartDelete(utility, app, element, passFile, logFile)

}
