package main

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	// configDirPath = "tmpdata/"
	configDirPath = os.Getenv("HOME") + "/.config/falsename/"
	DataDirPath   = configDirPath
)

func init() {
	if _, err := os.Stat(configDirPath); os.IsNotExist(err) {
		os.MkdirAll(configDirPath, 0755)
	}

	dataDirPathFile := configDirPath + "data_dir_path.txt"
	if _, err := os.Stat(dataDirPathFile); err == nil {
		data, err := os.ReadFile(dataDirPathFile)
		if err == nil {
			DataDirPath = strings.TrimSpace(string(data))
			if !strings.HasSuffix(DataDirPath, "/") {
				DataDirPath += "/"
			}
		}
	}
}

func SetDataDirPath(path string) {
	DataDirPath = path
	os.WriteFile(configDirPath+"data_dir_path.txt", []byte(DataDirPath), 0644)
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
		commands = append(commands, Command{Name: name, Command: command})
	}

	if _, err := os.Stat(DataDirPath); os.IsNotExist(err) {
		os.MkdirAll(DataDirPath, 0755)
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

func GetAllCommands() []Command {
	f, _ := os.Open(DataDirPath + "aliases.yml")
	defer f.Close()

	var commands []Command
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
