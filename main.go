package main

import (
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
		commands := GetAllCommands()
		selectedAlias := ChooseAlias(commands)
		if selectedAlias == -1 {
			return
		}
		if selectedAlias == -2 {
			fmt.Println("No aliases found. Use 'fn save' to save a command.")
			return
		}
		runCommand(commands[selectedAlias].Command)
		return
	}
	if flag.Arg(0) == "list" { // list
		commands := GetAllCommands()
		for _, cmd := range commands {
			fmt.Printf("%s: %s\n", cmd.Name, cmd.Command)
		}
		return
	}

	if flag.Arg(0) == "get" { // get
		cmd := GetCommand(flag.Arg(1))
		fmt.Println(cmd)
		return
	}

	if flag.Arg(0) == "shell" { // shell
		if flag.NArg() == 1 {
			shell := GetShell()
			fmt.Println(shell)
			return
		}

		if flag.NArg() == 2 {
			SetShell(flag.Arg(1))
			return
		}
	}

	if flag.Arg(0) == "help" { // help
		showHelp()
		return
	}

	if flag.Arg(0) == "delete" { // delete
		DeleteCommand(flag.Arg(1))
		return
	}

	if flag.Arg(0) == "save" {
		if flag.NArg() != 3 {
			fmt.Println("Usage: fn save <alias> <command>")
			return
		}

		SaveCommand(flag.Arg(1), flag.Arg(2))
		return
	}

	if flag.Arg(0) == "config" {
		if flag.NArg() == 1 {
			fmt.Println(DataDirPath)
			return
		}

		if flag.NArg() == 2 {
			DataDirPath = flag.Arg(1)
			SetDataDirPath(DataDirPath)
			return
		}
	}

	if flag.NArg() >= 1 { // run
		cmd := GetCommand(flag.Arg(0))
		if cmd == "--not found--" {
			fmt.Println("Alias not found. Use fn save to save a command.")
			return
		}

		remainingArgs := flag.Args()[1:]
		argStr := ""
		for _, arg := range remainingArgs {
			argStr += " " + arg
		}

		cmd += argStr
		runCommand(cmd)
		return
	}

	showHelp()
}

func showHelp() {
	configText := fmt.Sprintf(`
falsename - a simple cross-shell command aliaser

Usage:
	fn   -> choose an alias
	fn <alias> [args]   -> run an alias with optional arguments
	fn list   -> list all aliases
	fn get <alias>   -> get the command for an alias
	fn save <alias> <command>   -> save an alias
	fn delete <alias>   -> delete an alias
	fn shell   -> get configured shell (default /bin/sh)
	fn shell <shell>   -> set shell (not recommended unless you don't have /bin/sh)
	fn config   -> get the config directory
	fn config <dir>   -> set the config directory

The config directory is %s
`, DataDirPath)
	fmt.Println(configText)
}

func runCommand(cmdStr string) {
	if cmdStr == "--not found--" {
		fmt.Println("Alias not found. Use fn save to save a command.")
		return
	}
	cmd := exec.Command(GetShell(), "-c", cmdStr)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
