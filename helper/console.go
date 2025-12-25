package helper

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var Scanner = bufio.NewScanner(os.Stdin)

func ScanInput(promptString string) string {
	fmt.Print(promptString)
	Scanner.Scan()
	return Scanner.Text()
}

func PressEnter() {
	fmt.Print("Press [Enter] to continue...")
	if _, err := fmt.Scanln(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

func ClearScreen() {
	// Clear screen for Unix systems
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		// Clear screen for Windows
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("Error clearing screen:", err)
	}
}
