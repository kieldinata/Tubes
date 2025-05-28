package ui

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScrn() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Pause() {
	fmt.Print("\n\nTekan enter untuk melanjutkan...")
	fmt.Scanln()
}