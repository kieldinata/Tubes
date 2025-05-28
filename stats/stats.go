package stats

import (
	"app/global"
	"app/process"
	"app/ui"
	"fmt"
)

func Menu() {
	var i int
	var komentar string
	var scr float64

	ui.ClearScrn()
	fmt.Scanln()
	fmt.Println("List Semua Komentar:")
	for i = 0; i < global.NData; i++ {
		komentar = global.D[i].Komentar
		scr = process.HitungScore(komentar)
		fmt.Printf("%d. ", i+1)
		fmt.Print(komentar)
		fmt.Printf(" [Score = %.2f]", scr)
		if scr < 0 {
			fmt.Println(" [KOMENTAR NEGATIF]")
		} else if scr > 0.5 {
			fmt.Println(" [KOMENTAR POSITIF]")
		} else {
			fmt.Println(" [KOMENTAR NETRAL]")
		}
	}
	ui.Pause()
}
