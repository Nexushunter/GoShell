package cmnds

import (
	"os"
	"os/user"
	"strings"
)

// @author Hunter Breathat
// @License BSD 3-Clause License Copyright (c) 2018, Hunter Breathat All rights reserved.
// @repo	{github,gitlab}.com/NexisHunter/GoShell/cmnds/

//-----------------------Change Directory------------------------------------\\
// Change working directory
// Precondition: user wants to change current working directory
// Post-condition: Either changes working directory, or reports the error
//---------------------------------------------------------------------------\\
func CD(directory string) (fP string, err error) {
	currUser, _ := user.Current()
	root := "/"                                             // root directory
	home := "~"                                             // home directory
	quickAccess := []string{"Desktop", "Music", "Pictures", // Quick access to...
		"Videos", "Documents"}

	if directory == root {
		err := os.Chdir(root) // cd into root
		return updateDir(fP, err)

	} else if directory == home {
		err := os.Chdir(currUser.HomeDir) // cd $HOME
		return updateDir(fP, err)

	} else if len(directory) > 1 && strings.Contains(directory, "~") {
		directory = strings.Replace(directory, "~",
			currUser.HomeDir, -1) // Allows for ~ usage

		err := os.Chdir(directory)
		return updateDir(fP, err)

	} else {

	Outer:
		for i := 0; i < len(quickAccess); i++ {
			if strings.Contains(directory, quickAccess[i]) &&
				len(directory) == len(quickAccess[i]) {

				directory = currUser.HomeDir + "/" +
					quickAccess[i] // Allows for cd Desktop
				break Outer
			}
		}

		err := os.Chdir(directory)
		return updateDir(fP, err)
	}
}

// func CD helper method
func updateDir(fP string, err error) (string, error) {

	if err != nil {
		// Doesn't change working directory
		return fP, err     // Reports the error
	}
	
	fP, _ = os.Getwd() // Get the new working directory
	return fP, nil     // No error
}
