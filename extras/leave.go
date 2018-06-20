package extras

import (
	"os"
)

// @author  Hunter Breathat
// @License Copyright (R) 2018 Hunter Breathat
// @repo	{github,gitlab}.com/NexisHunter/GoShell/extras/

//-------------Exit ---------------------------------------------------------\\
// Simple exit strategy
// Precondition: Check if exiting program
// Post-condition: Terminates program if necessary
//---------------------------------------------------------------------------\\
func Leave(command string) {
	exit := "exit"
	kill := string(0x4)
	if command == exit || command == kill {
		os.Exit(0)
	}
}
