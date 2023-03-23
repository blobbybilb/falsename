package main

import (
	"falsename/choose"
	"falsename/data"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	var saveCommandFlag bool

	flag.BoolVar(&saveCommandFlag, "s", false, "Save the command")
	flag.Parse()

	if flag.NArg() == 0 { // choose
		commands := data.GetAllCommands()
		selectedAlias := choose.ChooseAlias(commands)
		if selectedAlias == -1 {
			return
		}
		runCommand(commands[selectedAlias].Command)
		return
	}
	if flag.Arg(0) == "list" { // list
		commands := data.GetAllCommands()
		for _, cmd := range commands {
			fmt.Printf("%s: %s\n", cmd.Name, cmd.Command)
		}
		return
	}

	if flag.Arg(0) == "get" { // get
		cmd := data.GetCommand(flag.Arg(1))
		fmt.Println(cmd)
		return
	}

	if flag.Arg(0) == "help" { // help
		printUsageAndExit()
		return
	}

	if flag.NArg() == 1 { // run
		cmd := data.GetCommand(flag.Arg(0))
		runCommand(cmd)
		return
	}

	if saveCommandFlag { // save
		data.SaveCommand(flag.Arg(0), flag.Arg(1))
		return
	}

	printUsageAndExit()
}

func printUsageAndExit() {
	fmt.Printf(`
falsename - a simple cross-shell command aliaser

Usage:
  fn   -> choose an alias
  fn <alias>   -> run an alias
  fn list   -> list all aliases
  fn get <alias>   -> get the command for an alias
  fn -s <alias> <command>   -> save an alias

The config directory is %s
	
	`, data.DataDirPath)
	os.Exit(0)
}

func runCommand(cmdStr string) {
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
