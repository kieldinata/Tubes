package stats

import (
	"app/global"
	"app/process"
	"app/ui"
	"fmt"
)

func Menu() {
	var pilih int
	for pilih != 3 {
		ui.ClearScrn()
		ui.Logo()
		pilih = 0
		fmt.Println("-------------------------------")
		fmt.Println("STATISTIC MENU")
		fmt.Println("-------------------------------")
		fmt.Println("1. Tunjukkan Semua Statistik")
		fmt.Println("2. Cari Komentar")
		fmt.Println("3. Back")
		fmt.Print("Pilih dengan angka (1/2/3): ")
		fmt.Scanf("%d", &pilih)

		switch pilih {
		case 1:
			ui.ClearScrn()
			filters()
		case 2:
			ui.ClearScrn()
			//search()
		}
	}
}

// bagian sorting
func filters() {
	var pilih int
	for pilih != 5 {
		ui.ClearScrn()
		ui.Logo()
		pilih = 0
		fmt.Println("Tunjukkan Komentar berdasarkan:")
		fmt.Println("1. Terburuk (ascending)")
		fmt.Println("2. Terbaik (descending)")
		fmt.Println("3. A-Z (ascending)")
		fmt.Println("4. Z-A (descending)")
		fmt.Println("5. Back")
		fmt.Print("Pilih angka (1/2/3/4/5): ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			worst()
		case 2:
			best()
		case 3:
			az()
		case 4:
			za()
		}
	}

}
func show(A global.TabData, filterName string) {
	var i int
	var komentar string
	var score float64

	ui.ClearScrn()
	fmt.Scanln()
	ui.Logo()
	fmt.Println("List Komentar Berdasarkan Filter " + filterName + ":")
	if global.NData == 0 {
		fmt.Println("Belum ada komentar, silahkan isi terlebih dahulu.")
	} else {
		for i = 0; i < global.NData; i++ {
			komentar = A[i].Komentar
			score = A[i].Score
			fmt.Printf("%d. %s ", i+1, komentar)
			fmt.Printf(" [Score = %.2f]", score)
			if score < 0 {
				fmt.Println(" [KOMENTAR NEGATIF]")
			} else if score > 0.5 {
				fmt.Println(" [KOMENTAR POSITIF]")
			} else {
				fmt.Println(" [KOMENTAR NETRAL]")
			}
		}
	}
	ui.Pause()

}

func worst() {
	var W global.TabData = process.HitungScore((global.D))
	show(process.InsertionSortData(W, "asc", "Score"), "From Worst")
}

func best() {
	var B global.TabData = process.HitungScore((global.D))
	show(process.InsertionSortData(B, "dsc", "Score"), "From Best")
}

func az() {
	var AZ global.TabData = process.SelectionSortData(global.D, "asc", "Komentar")
	show(process.HitungScore(AZ), "From A to Z")
}

func za() {
	var ZA global.TabData = process.SelectionSortData(global.D, "dsc", "Komentar")
	show(process.HitungScore(ZA), "From Z to A")
}

/*func search() {
	var A global.TabData
	var key, check string
	var i, j int

	ui.ClearScrn()
	fmt.Scanln()
	fmt.Print("tulis Abjad yang ingin dicari: ")
	fmt.Scan(&key)
	for i = 0; i < global.NData; i ++{
		A[j] = global.D[i].Komentar
		check
		for
	}

}*/
