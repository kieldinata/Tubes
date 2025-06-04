package main

import (
	"app/color"
	"app/comment"
	"app/sample"
	"app/stats"
	"app/ui"
	"fmt"
)

func main() {
	var pilih int
	for pilih != 4 {
		pilih = 0
		ui.ClearScrn()
		ui.Logo()
		fmt.Println(color.Title + "<MAIN MENU..........................>" + color.Reset)
		fmt.Println("")
		fmt.Println("1. Menu Komentar")
		fmt.Println("2. Menu Sample")
		fmt.Println("3. Statistik")
		fmt.Printf("4. %sExit%s\n", color.Red, color.Reset)
		fmt.Print("Pilih dengan angka (1/2/3/4): " + color.Input)
		fmt.Scan(&pilih)
		fmt.Print(color.Reset)

		switch pilih {
		case 1:
			comment.Menu()
		case 2:
			sample.Menu()
		case 3:
			stats.Menu()
		}
	}
}
