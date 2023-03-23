package data

import (
	"falsename/types"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	DataDirPath = "~/.config/.falsename/"
)

func init() {
	if _, err := os.Stat(DataDirPath); os.IsNotExist(err) {
		os.MkdirAll(DataDirPath, 0755)
	}
}

func SaveCommand(name, command string) {
	commands := GetAllCommands()
	commands = append(commands, types.Command{Name: name, Command: command})

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
