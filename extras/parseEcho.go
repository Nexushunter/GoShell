package extras

import "strings"

// @author  Hunter Breathat
// @License Copyright (R) 2018 Hunter Breathat
// @repo	NexisHunter/GoShell/extras

//-------------------------Custom Echo Parsing-------------------------------\\
// Parses Echo command based on how many args were present ie contained sudo
// Precondition: Echo was the command being passed
// Post-condition: Echo is parsed
//---------------------------------------------------------------------------\\
func ParseEcho(command string) []string {
	if strings.Contains(command, "sudo") {
		return strings.SplitN(command, " ", 2)
	} else {
		return strings.SplitN(command, " ", 1)
	}
}
