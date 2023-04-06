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
	var helpFlag bool

	// flag.BoolVar(&saveCommandFlag, "s", false, "Save the command")
	flag.BoolVar(&helpFlag, "h", false, "Show help")
	flag.Parse()

	if helpFlag {
		fmt.Println("falsename - a simple cross-shell command aliaser. For help: fn help")
		return
	}

	if flag.NArg() == 0 { // choose
		commands := data.GetAllCommands()
		selectedAlias := choose.ChooseAlias(commands)
		if selectedAlias == -1 {
			return
		}
		if selectedAlias == -2 {
			fmt.Println("No aliases found. Use -s to save a command.")
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

	if flag.Arg(0) == "shell" { // shell
		if flag.NArg() == 1 {
			shell := data.GetShell()
			fmt.Println(shell)
			return
		}

		if flag.NArg() == 2 {
			data.SetShell(flag.Arg(1))
			return
		}
	}

	if flag.Arg(0) == "help" { // help
		showHelp()
		return
	}

	if flag.Arg(0) == "delete" { // delete
		data.DeleteCommand(flag.Arg(1))
		return
	}

	if flag.Arg(0) == "save" {
		if flag.NArg() != 3 {
			fmt.Println("Usage: fn -s <alias> <command>")
			return
		}

		data.SaveCommand(flag.Arg(1), flag.Arg(2))
		return
	}

	if flag.NArg() == 1 { // run
		cmd := data.GetCommand(flag.Arg(0))
		runCommand(cmd)
		if cmd == "--not found--" {
			fmt.Println("Alias not found. Use -s to save a command.")
			return
		}
		return
	}

	showHelp()
}

func showHelp() {
	configText := fmt.Sprintf(`
falsename - a simple cross-shell command aliaser

Usage:
	fn   -> choose an alias
	fn <alias>   -> run an alias
	fn list   -> list all aliases
	fn get <alias>   -> get the command for an alias
	fn save <alias> <command>   -> save an alias
	fn shell <shell> -> set shell
	fn shell -> get configured shell
	fn delete <alias>   -> delete an alias

The config directory is %s
`, data.DataDirPath)
	fmt.Println(configText)
}

func runCommand(cmdStr string) {
	cmd := exec.Command(data.GetShell(), "-c", cmdStr)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
