package cmnds

import (
	"../extras"
	"os"
)
//---------------------------------------------------------------------------\\
// Gets and prints the current working directory
//
//
//---------------------------------------------------------------------------\\
func PWD() string{

	s,err := os.Getwd()

	// Prints the error if any
	if err != nil {
		extras.PrintErr(err)
	}

	return s
}