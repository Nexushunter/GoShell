package extras

import (
	"errors"
	"strings"
)

// @author  Hunter Breathat
// @License BSD 3-Clause License Copyright (c) 2018, Hunter Breathat All rights reserved.
// @repo	{github,gitlab}.com/NexisHunter/GoShell/extras/

//----------------------------Counts Total Commands--------------------------\\
// Precondition: Must contain at least 1 command
// Post-condition: Returns the number of commands to be executed.... Thus
//		allowing for combo-line commands ie. cat...'&&'echo... etc
//---------------------------------------------------------------------------\\
func CountCmnds(command string) ([]string, int, error) {

	var size int
	var commands []string

	if command == "" {
		return nil, 0, errors.New("empty command inputted")
	}

	for i := 0; i < len(command); i++ {

		// Makes it easy for identifying a separate command
		separator := string(command[i] + command[i+1])
		if separator == "&&" {

			// Checks for the && and adds 1 to count since the shell
			// command cannot end with && etc
			size++
		}

	}
	commands = separateCmnds(command, size) //temp
	return commands, size, nil
}

func separateCmnds(command string, size int) (commands []string) {
	commands = strings.SplitN(command, "&&", size)
	return commands
}
