package cmnds

import (
	"errors"
	"fmt"
	"github.com/nexishunter/GoShell/extras"
	"os"
	"strings"
)

// @author Hunter Breathat
// @License Copyright (R) 2018 Hunter Breathat
// @repo	{github,gitlab}.com/NexisHunter/GoShell/cmnds/

//-----------------------------------Echo------------------------------------\\
// Prints input or Environment variable or allows for appending to files
// Precondition: Not applicable
// Post-condition: Either appends to file or prints input/variable
//---------------------------------------------------------------------------\\
func Echo(command string) {
	err := errors.New("echo : Function not valid/available yet")

	if strings.Contains(command, "$") {
		s := strings.TrimPrefix(command, "$")
		fmt.Println(os.Getenv(s))

	} else if strings.Contains(command, ">>") {
		parts := strings.Split(command, ">>")
		input := parts[0] // Whats is being appended
		fName := parts[1] // The file
		fName = strings.Replace(fName, "\"", "", -1)
		// ----> Prevents errors until fileIO is added
		fmt.Println(input + " " + fName)
		extras.PrintErr(err)

	} else {

		fmt.Println(command)
	}
}
