package extras

import (
	"strings"
)

// @author  Hunter Breathat
// @License BSD 3-Clause License Copyright (c) 2018, Hunter Breathat All rights reserved.
// @repo	{github,gitlab}.com/NexisHunter/GoShell/extras/

//-------------------------Custom Echo Parsing-------------------------------\\
// Parses Echo command based on how many args were present ie contained sudo
// Precondition: Echo was the command being passed
// Post-condition: Echo is parsed
//---------------------------------------------------------------------------\\
func ParseEcho(command string) (commands []string) {
	if strings.Contains(command, "sudo") {
		commands = strings.SplitN(command, " ", 3)
		return commands
	} else {
		commands = strings.SplitN(command, " ", 2)
		return commands
	}
}
