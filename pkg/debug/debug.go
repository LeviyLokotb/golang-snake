package debug_tools

import (
	"errors"
	"fmt"
	"os"
)

func AddToLog(messages ...any) {
	_, log := pathToLog()

	file, err := os.OpenFile(log, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintln(messages...))
	if err != nil {
		panic(err)
	}
}

func GetHomeDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "_", err
	}
	if home == "" {
		return "_", errors.New("Invalid home directory")
	}
	return home, nil
}

func pathToLog() (dir, path string) {
	home, err := GetHomeDir()
	if err != nil {
		panic(err)
	}
	dir = home + "/" + ".snake"
	path = dir + "/" + "log.txt"
	return dir, path
}

func ClearLog() {
	dir, path := pathToLog()

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString("")
	if err != nil {
		panic(err)
	}
}
