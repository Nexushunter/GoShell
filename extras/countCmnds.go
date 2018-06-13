package extras

import (
	"errors"
	"strings"
)

// @author  Hunter Breathat
// @License Copyright (R) 2018 Hunter Breathat
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
		return nil, 0, errors.New("invalid amount of commands")
	}

	for i := 0; i < len(command); i++ {

		if command[i] == '&' {

			if command[i+1] == '&' {

				// Checks for the && and adds 1 to count since the shell
				// command cannot end with && etc
				size++
			}

		}

	}
	commands = separateCmnds(command, size) //temp
	return commands, size, nil
}

func separateCmnds(command string, size int) (commands []string) {
	commands = strings.Split(command, "&&")
	return commands
}
