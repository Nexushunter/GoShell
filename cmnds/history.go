package cmnds

import (
	"os"
	"io/ioutil"
	"fmt"
	"io"
)

// @author  Hunter Breathat
// @License Copyright (R) 2018 Hunter Breathat
// @repo	{github,gitlab}.com/NexisHunter/GoShell/extras/

// Contains the history of the inputted commands if any skips empty \r\n | \n
// | \r input as it does not count as a command

// Returns the history as a string to be outputted
// Otherwise it is an empty string
const fName = "/.history_gs"

func LoadHistory() (string) {
	history, good := checkForHistory()
	if !good {
		SaveHistory("")
		return ""
	} else {
		return string(history)
	}
}

// Save the history to the .history file
func SaveHistory(command string) {

	// Open the .history file, or sreate it, and append the most recent command
	f, err := os.OpenFile(os.Getenv("HOME") + fName,os.O_APPEND|
		os.O_CREATE|os.O_WRONLY,0600)

	if err != nil {
		fmt.Print(err)
		return
	}
	defer f.Close()

	n,err := io.WriteString(f,command + "\n")
	if err != nil {
		fmt.Println(n,err)
		return
	}

}

// Checks for an empty history string
func emptyHistory(history string) bool {
	return len(history) == 0
}

// Displays history
func printHistory(){
	history := LoadHistory()
	fmt.Print(history)
}

// Checks to see if the .history file exists.
func checkForHistory() ([]byte,bool){
	hist, err := ioutil.ReadFile(os.Getenv("HOME") + fName)
	if err != nil {
		return hist,true
	} else {
		return nil, false
	}
}