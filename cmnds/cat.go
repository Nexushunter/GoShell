package cmnds

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// @author Hunter Breathat
// @License Copyright (R) 2018 Hunter Breathat
// @repo	{github,gitlab}.com/NexisHunter/GoShell/cmnds/

//---------------------------Cat--------------------------------------------\\
// Displays contents of file to terminal
// Precondition: File must exist
// Post-condition: Displays info if file exists
//---------------------------------------------------------------------------\\
func Cat(fileName string) (err error) {
	//err = errors.New("cat : Function currently unavailable")
	var content string
	cwd, _ := os.Getwd()

	if !strings.HasPrefix(fileName, "/") {

		// Assuming current directory
		content, err = rf(cwd + fileName)

		if err == nil {
			fmt.Println(string(content))
			return err
		} else {
			return err
		}

	} else {

		content, err = rf(fileName)

		if err == nil {
			fmt.Println(string(content))
			return err
		} else {
			return err
		}
	}
}

//-----------------------I/O Portion-----------------------------------------\\
// The I/O portion of the cat command
// Precondition: cat must called
// Post-condition: I/O handled and returned
//---------------------------------------------------------------------------\\
func rf(fileName string) (string, error) {
	content, err := ioutil.ReadFile(fileName)
	return string(content), err
}
