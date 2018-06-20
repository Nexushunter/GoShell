package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/nexishunter/GoShell/cmnds"
	"github.com/nexishunter/GoShell/extras"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

// @author Hunter Breathat
// @License Copyright (R) 2018 Hunter Breathat
// @repo	{github,gitlab}.com/NexisHunter/GoShell

/*	Objective: To create a go based shell for unix based oses
	Currently being tested on Debian based OSes ie. Ubuntu 18.04, Kali Linux
	Thus allowing for scripting using go as well as bash etc.*/

var (
	currUser, _ = user.Current() // Current User
	newDir      string           // New Directory
)

//---------------------------------------------------------------------------\\
// The basic application of the shell
// Precondition: N/A
// Post-condition: Runs user commands/input if available
//---------------------------------------------------------------------------\\
func main() {
	var command string  // To be user input
	var err error       // Generic error
	var cmd *os.Process // The process of the command
	var commands []string
	var input = bufio.NewScanner(os.Stdin) // Takes user input

	for {
		var currDir string

		if currDir == "" && newDir == "" { // Checking if WD is empty
			currDir, _ = cmnds.CD("~") // Default to $HOME
		} else {
			currDir = newDir // Else set to new
		}

		if strings.Contains(currDir, currUser.HomeDir) {
			currDir = strings.Replace(currDir, currUser.HomeDir, "~", -1)
		}

		fmt.Printf("%s@NexisOs: %s> ", currUser.Username,
			strings.Replace(currDir, currUser.HomeDir, "~", -1)) // Main line in terminal

		input.Scan()           //stores user input
		command = input.Text() // Stores user input

		extras.Leave(command) // Check if command is exit or ^D

		if strings.Contains(command, "echo") {
			commands = extras.ParseEcho(command)
		} else {
			commands = strings.Split(command, " ") // Separates args from command/application
		}

		extras.Leave(commands[0]) // Check if command is exit or ^D
		command = commands[0]     // Set command to current command for easy checking

		if sudoReq(command) {

			command = commands[1]
			if command == "cd" {
				newDir, err = cmnds.CD(command)
				extras.PrintErr(err)
			} else if command == "echo" {
				cmnds.Echo(commands[2])
			} else {
				cmd, err = execute(commands)

				if cmd != nil { // Active command/process
					cmd.Wait() // Wait for command to finish
					cmd.Kill() // Terminate the program after completion
				}

				extras.PrintErr(err)
			}

		} else {

			if command == "cd" {
				newDir, err = cmnds.CD(commands[1])
				extras.PrintErr(err)
			} else if command == "echo" {
				cmnds.Echo(commands[1])
			} else {
				cmd, err = execute(commands)

				if cmd != nil { // Active command/process
					cmd.Wait() // Wait for command to finish
					cmd.Kill() // Terminate the program after completion
				}

				extras.PrintErr(err)
			}
		}
	}
}

//--------------------------Launch Process-----------------------------------\\
//	Os agnostic approach to launching a process
// Precondition: No currently running command
// Post-condition: Command is now runnning
//---------------------------------------------------------------------------\\
func execute(command []string) (p *os.Process, err error) {
	binary, err := exec.LookPath(command[0])
	if binary == "" {
		// Reports invalid commands/input
		return nil, errors.New(fmt.Sprintf("  %s : command not found",
			command[0]))
	}
	if binary, err := exec.LookPath(command[0]); err == nil {
		var attributes os.ProcAttr // Set up Process Attributes
		attributes.Files = []*os.File{
			os.Stdin, os.Stdout, os.Stderr} // Setting IO and Error
		p, err := os.StartProcess(binary,
			command, &attributes) // Fork and execute command
		if err == nil {
			return p, nil // Successful execution
		}
	}
	return nil, err // Fail and Error reporting
}

//---------------------------Check for sudo----------------------------------\\
// Checks for sudo in the command args
// Precondition: User must input something
// Post-condition: Verifies if sudo is present
//---------------------------------------------------------------------------\\
func sudoReq(command string) bool {
	if command == "sudo" {
		return true
	} else {
		return false
	}
}
