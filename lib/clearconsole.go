package lib

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearConsole() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
