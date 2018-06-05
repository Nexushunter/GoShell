package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

// @author Hunter Breathat
// @License Copyright (R) 2018 Hunter Breathat
// @repo	NexisHunter/GoShell
/*	Objective: To create a go based shell for unix based oses
	Currently being tested on Debian based OSes ie. Ubuntu 18.04, Kali Linux
	Thus allowing for scripting using go as well as bash etc.*/

var (
	currUser, _ = user.Current()
)

func main() {
	var command string                    // To be user input
	var input = bufio.NewReader(os.Stdin) // Takes user input
	//var currUser, _ = user.Current()										// Getting the current user
	var newDir string   // New Directory
	var err error       // Generic error
	var cmd *os.Process // The process of the command

	for {
		var currDir string

		if currDir == "" && newDir == "" { // Checking if WD is empty
			currDir, _ = cd("~") // Default to $HOME
		} else {
			currDir = newDir // Else set to new
		}

		if strings.Contains(currDir, currUser.HomeDir) {
			currDir = strings.Replace(currDir, currUser.HomeDir, "~", -1)
		}
		fmt.Printf("%s@NexisOs: %s> ", currUser.Username,
			strings.Replace(currDir, currUser.HomeDir, "~", -1)) // Main line in terminal

		command, _ = input.ReadString('\n')         // Takes user input checking for \n
		command = strings.TrimSuffix(command, "\n") // Removes trailing \n from the user input
		commands := strings.Split(command, " ")     // Separates args from command/application
		if commands[0] == "cd" ||                   // Checking for sudo
			(commands[0] == "sudo" && commands[1] == "cd") {
			if commands[1] != "cd" {
				newDir, err = cd(commands[1])
				if err != nil {
					printErr(err)
				}

			} else {
				newDir, err = cd(commands[2])
				if err != nil {
					printErr(err)
				}
			}
		} else {
			cmd, err = execute(commands)

			if cmd != nil { // Active command/process
				cmd.Wait() // Wait for command to finish
				cmd.Kill() // Terminate the program after completion
			}

			if err != nil {
				printErr(err)  // Displays the error to the user
				log.Fatal(err) // Log the error
			}
		}
	}
}

//	Os agnostic approach to launching a process
func execute(command []string) (p *os.Process, err error) {
	if binary, err := exec.LookPath(command[0]); err == nil {
		var attributes os.ProcAttr                                    // Set up Process Attributes
		attributes.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr} // Setting IO and Error
		p, err := os.StartProcess(binary, command, &attributes)       // Fork and execute command
		if err == nil {
			return p, nil // Successful execution
		}
	}
	return nil, err // Fail and Error reporting
}

// Change working directory
// Precondition: user wants to change current working directory
// Postcondition: Either changes working directory, or reports the error
func cd(filePath string) (fP string, err error) {
	root := "/"                                            // root directory
	home := "~"                                            // home directory
	quickAcess := []string{"Desktop", "Music", "Pictures", // Quick access to ...
		"Videos", "Documents"}

	if filePath == root {
		err := os.Chdir(root) // cd into root
		return updateDir(fP, err)
	} else if filePath == home {
		err := os.Chdir(currUser.HomeDir) // cd $HOME
		return updateDir(fP, err)
	} else if len(filePath) > 1 && strings.Contains(filePath, "~") {
		filePath = strings.Replace(filePath, "~", currUser.HomeDir, -1) // Allows for ~ usage
		err := os.Chdir(filePath)
		return updateDir(fP, err)
	} else {
	Outer:
		for i := 0; i < len(quickAcess); i++ {
			if strings.Contains(filePath, quickAcess[i]) &&
				len(filePath) == len(quickAcess[i]) {
				filePath = currUser.HomeDir + "/" + quickAcess[i] // Allows for cd Desktop
				break Outer
			}
		}
		err := os.Chdir(filePath)
		return updateDir(fP, err)
	}
}

// func cd helper method
func updateDir(fP string, err error) (filePath string, nah error) {
	if err != nil {
		fP, _ := os.Getwd() // Doesn't change working directory
		return fP, err      // Reports the error
	} else {
		fP, _ := os.Getwd() // Get the new working directory
		return fP, nil      // No error
	}
}

// Simply Prints the error
func printErr(err error) {
	fmt.Println(err)
}
