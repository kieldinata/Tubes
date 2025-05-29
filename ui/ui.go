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

func Logo() {
	fmt.Println(" _____       _____        ___      ")
	fmt.Println("/  __ \\     /  ___|      / _ \\     ")
	fmt.Println("| /  \\/     \\ `--.      / /_\\ \\    ")
	fmt.Println("| |          `--. \\     |  _  |    ")
	fmt.Println("| \\__/\\  _  /\\__/ /  _  | | | |  _ ")
	fmt.Println("\\_____/ (_) \\____/  (_) \\_| |_/ (_)")
	fmt.Println()
}
