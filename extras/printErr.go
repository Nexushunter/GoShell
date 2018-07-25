package extras

import "fmt"

// @author  Hunter Breathat
// @License Copyright (R) 2018 Hunter Breathat
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
