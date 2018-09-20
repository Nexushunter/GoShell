package main

import (
	"bufio"
	"fmt"
	"./cmnds"
	"./extras"
	"os"
	"os/user"
	"strings"
	"os/exec"
	"errors"
)

// @author Hunter Breathat
// @License BSD 3-Clause License Copyright (c) 2018, Hunter Breathat All rights reserved.
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

		switch currDir {
			case "":
				switch newDir {
				case "":
					currDir, _ = cmnds.CD("~") // Default to $HOME
				}
			default:
				currDir = newDir 						// Otherwise new Dir
		}

		switch currDir {
			case currUser.HomeDir:
				currDir =  strings.Replace(currDir,currUser.HomeDir,"~",-1)
		}

		fmt.Printf("%s@NexisOs: %s> ", currUser.Username,
			strings.Replace(currDir, currUser.HomeDir, "~", -1)) // Main line in terminal

		scan := input.Scan()           //stores user input
		command = input.Text() // Stores user input

		eof := !scan && input.Err() == nil

		switch eof {		// Check if command is exit or ^D
		case true:
			extras.LeaveEOF()
		default:
			extras.Leave(command)
		}

		//cmnds.SaveHistory(command) // Save the command to the history file

		commands = strings.Split(command," ")

		switch commands[0] {
		case "echo":
			commands = extras.ParseEcho(command)
		default:
			commands = strings.Split(command," ")
		}

		command = commands[0]     // Set command to current command for easy checking

		switch command {
		case "sudo":
			switch commands[1] {
			case "cd":
				newDir, err = cmnds.CD(commands[2])
				extras.PrintErr(err)
			case "echo":
				cmnds.Echo(commands[2])
			default:
				cmd, err = execute(commands)

				if cmd != nil { // Active command/process
					cmd.Wait() // Wait for command to finish
					cmd.Kill() // Terminate the program after completion
				}

				extras.PrintErr(err)
			}
		case "cd":
			newDir, err = cmnds.CD(commands[1])
			extras.PrintErr(err)
		case "echo":
			cmnds.Echo(commands[1])
		default:
			cmd, err = execute(commands)

			if cmd != nil { // Active command/process
				cmd.Wait() // Wait for command to finish
				cmd.Kill() // Terminate the program after completion
			}

			switch err{
			case nil:
				break
			default:
				extras.PrintErr(err)
			}
		}

	}

}

//--------------------------Launch Process-----------------------------------\\
//	Os agnostic approach to launching a process
// 	Precondition: No currently running command
// 	Post-condition: Command is now running
//---------------------------------------------------------------------------\\
func execute(command []string) (p *os.Process, err error) {

	switch command[0]{
	case "":
		return nil,nil
	default:
		binary, err := exec.LookPath(command[0])

		switch binary{
		case "":
			// Reports invalid commands/input
			return nil, errors.New(fmt.Sprintf("  %s : command not found",
				command[0]))
		default:

			switch err{
			case nil:
				var attributes os.ProcAttr // Set up Process Attributes
				attributes.Files = []*os.File{
				os.Stdin, os.Stdout, os.Stderr} // Setting IO and Error
				p, err := os.StartProcess(binary,
				command, &attributes) // Fork and execute command

				switch err {
				case nil:
					return p, nil // Successful execution
				default:
					return nil, err // Fail and Error reporting
				}
			}
		}
	}
	return

}