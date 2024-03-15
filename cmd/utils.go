package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func clearTerminal() {
	// Check the operating system
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// For Unix-like systems, use ANSI escape codes
		fmt.Print("\033[H\033[2J")
	}
}

// Also clears the terminal buffer
func PressEnterToContinue() {
	fmt.Print("Press Enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	clearTerminal()
}
