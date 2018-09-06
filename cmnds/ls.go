package cmnds

import (
	"os"
	"../extras"
	"io/ioutil"
	"fmt"
)

// @author 	Hunter Breathat
// @License BSD 3-Clause License Copyright (c) 2018, Hunter Breathat All rights reserved.
// @repo	{github,gitlab}.com/NexisHunter/GoShell/cmnds/

//------------------------------LS-------------------------------------------\\
// Lists the files in a given directory, if given, otherwise assumes current
// directory as the given dir
//---------------------------------------------------------------------------\\
var(
	err error
	dirInfo []os.FileInfo
	)
func Ls(dir string){

	switch dir {
	case "":
		dir,err = os.Getwd()
	}

	if err != nil {
		extras.PrintErr(err)
	}
	
	dirInfo,err = ioutil.ReadDir(dir)

	for _,f := range dirInfo {
		fmt.Println(f.Name())
	}
	
}
