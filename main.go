package main

import (
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
		fmt.Println("-------------------------------")
		fmt.Println("MAIN MENU")
		fmt.Println("-------------------------------")
		fmt.Println("1. Menu Komentar")
		fmt.Println("2. Menu Sample")
		fmt.Println("3. Statistik")
		fmt.Println("4. Exit")
		fmt.Print("Pilih dengan angka (1/2/3/4): ")
		fmt.Scanf("%d", &pilih)

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
