package extras

import (
	"os"
	"fmt"
)

// @author  Hunter Breathat
// @License BSD 3-Clause License Copyright (c) 2018, Hunter Breathat All rights reserved.
// @repo	{github,gitlab}.com/NexisHunter/GoShell/extras/

//-------------Exit ---------------------------------------------------------\\
// Simple exit strategy
// Precondition: Check if exiting program
// Post-condition: Terminates program if necessary
//---------------------------------------------------------------------------\\
func Leave(command string) {
	exit := "exit"
	switch command {
	case exit:
		os.Exit(0)
	}
}

func LeaveEOF(){
	fmt.Println("Exiting....")
	Leave("exit")
}
