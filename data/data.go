package data

import (
	"falsename/types"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	DataDirPath = os.Getenv("HOME") + "/.config/.falsename/"
	// DataDirPath = "tmpdata/"
)

func init() {
	if _, err := os.Stat(DataDirPath); os.IsNotExist(err) {
		os.MkdirAll(DataDirPath, 0755)
	}
}

func SaveCommand(name, command string) {
	commands := GetAllCommands()

	var foundExisting bool

	for i, cmd := range commands {
		if cmd.Name == name {
			commands[i].Command = command
			foundExisting = true
			break
		}
	}

	if !foundExisting {
		commands = append(commands, types.Command{Name: name, Command: command})
	}

	f, _ := os.Create(DataDirPath + "aliases.yml")
	defer f.Close()

	enc := yaml.NewEncoder(f)
	enc.Encode(commands)
}

func GetCommand(name string) string {
	commands := GetAllCommands()
	for _, cmd := range commands {
		if cmd.Name == name {
			return cmd.Command
		}
	}
	return "--not found--"
}

func GetAllCommands() []types.Command {
	f, _ := os.Open(DataDirPath + "aliases.yml")
	defer f.Close()

	var commands []types.Command
	dec := yaml.NewDecoder(f)
	dec.Decode(&commands)

	return commands
}

func GetShell() string {
	f, err := os.Open(DataDirPath + "shell.yml")
	if err != nil {
		return "/bin/sh"
	}

	defer f.Close()

	var shell string
	dec := yaml.NewDecoder(f)
	dec.Decode(&shell)

	return shell
}

func SetShell(shell string) {
	f, _ := os.Create(DataDirPath + "shell.yml")
	defer f.Close()

	enc := yaml.NewEncoder(f)
	enc.Encode(shell)
}

func DeleteCommand(name string) {
	commands := GetAllCommands()

	for i, cmd := range commands {
		if cmd.Name == name {
			commands = append(commands[:i], commands[i+1:]...)
			break
		}
	}

	f, _ := os.Create(DataDirPath + "aliases.yml")
	defer f.Close()

	enc := yaml.NewEncoder(f)
	enc.Encode(commands)
}
