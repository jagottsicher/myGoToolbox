package myGoToolbox

import (
	"os"
	"os/exec"
	"runtime"
)

// Clearscreen empties the terminal screen independently from the underlying operation system
func Clearscreen() {
	if runtime.GOOS != "windows" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
