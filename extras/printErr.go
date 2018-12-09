package extras

import "fmt"

// @author  Hunter Breathat
// @License BSD 3-Clause License Copyright (c) 2018, Hunter Breathat All rights reserved.
// @repo	{github,gitlab}.com/NexisHunter/GoShell/extras/

//--------------Error Checking-----------------------------------------------\\
// Simply checks and Prints the error
// Precondition:   Need to check for Error
// Post-condition: Prints Error if found
//---------------------------------------------------------------------------\\
func PrintErr(err error) {
	if err != nil {
		fmt.Println(err) // Displays the error
		//log.Fatal(err)   								// Log the error
	}
}
