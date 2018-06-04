package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"log"
)

// @author Hunter Breathat
// @License MIT License
// @repo	NexisHunter/go/src/shell
/*	Objective: To create a go based shell for unix based oses
	Currently being tested on Debian based OSes ie. Ubuntu 18.04, Kali Linux
	Thus allowing for scripting using go as well as bash etc.*/

func main(){
	var command string
	var input = bufio.NewReader(os.Stdin)
	var currUser, _ = user.Current()

	for {
		fmt.Printf("%s@NexisOs> ", currUser.Username)

		command,_= input.ReadString('\n')
		command = strings.TrimSuffix(command,"\n")

		commands := strings.Split(command, " ")
		cmd,err := execute(commands)
		if cmd != nil {
			cmd.Wait()
			cmd.Kill()
		}
		if err != nil {
			//Display the errors
			fmt.Println(err)
			//Log the error
			log.Fatal(err)
		}

	}
}

func execute(command []string) (p *os.Process, err error) {

	if binary, err := exec.LookPath(command[0]); err == nil {
		//Set up ProcAttr
		var procAttr os.ProcAttr
		procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
		//Fork and execute command
		p, err := os.StartProcess(binary, command, &procAttr)
		if err == nil {
			return p, nil
		}
	}
	return nil, err
}



//child_pid = syscall.SYS_FORK
/*if child_pid  == 0 {
	execvp(command[0],command)
	fmt.Print("You shouldn't see me.....")
} else {
	waitpid(child_pid,&static_loc,syscall.WUNTRACED)
}*/