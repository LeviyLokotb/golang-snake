package debug_tools

import (
	"fmt"
	"os"
)

func AddToLog(message any) {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("Cannot open log.txt")
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintln(message))
	if err != nil {
		panic("Cannot write string to log.txt")
	}
}

func ClearLog() {
	file, err := os.Create("log.txt")
	if err != nil {
		panic("Cannot create log.txt")
	}
	defer file.Close()
	_, err = file.WriteString("")
	if err != nil {
		panic("Cannot write string to log.txt")
	}
}
