package stats

import (
	"app/color"
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
		fmt.Println(color.Title + "<STATISTIC MENU.....................>" + color.Reset)
		fmt.Println("")
		fmt.Println("1. Tunjukkan Semua Statistik")
		fmt.Println("2. Cari Komentar")
		fmt.Printf("3. %sBack%s\n", color.Red, color.Reset)
		fmt.Print("Pilih dengan angka (1/2/3): " + color.Input)
		fmt.Scan(&pilih)
		fmt.Print(color.Reset)
		process.HitungScore(&global.D)

		switch pilih {
		case 1:
			ui.ClearScrn()
			filters()
		case 2:
			ui.ClearScrn()
			search()
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
		fmt.Printf("5. %sBack%s\n", color.Red, color.Reset)
		fmt.Print("Pilih angka (1/2/3/4/5): " + color.Input)
		fmt.Scan(&pilih)
		fmt.Print(color.Reset)
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
func show(A global.TabData, filterName string, n int) {
	var i int
	var komentar string
	var score float64

	ui.ClearScrn()
	fmt.Scanln()
	ui.Logo()
	fmt.Println("List Komentar Berdasarkan Filter " + filterName + ":")
	if n == 0 {
		fmt.Println("Belum ada komentar, silahkan isi terlebih dahulu.")
	} else {
		for i = 0; i < n; i++ {
			komentar = color.LightBlue + A[i].Komentar + color.Reset
			score = A[i].Score
			fmt.Printf("%d. %s\n", i+1, komentar)
			fmt.Printf("[Score = %.2f] ", score)
			switch A[i].Category {
			case "[SANGAT POSITIF]":
				fmt.Println(color.Green + A[i].Category + color.Reset)
			case "[CUKUP POSITIF]":
				fmt.Println(color.LimeGreen + A[i].Category + color.Reset)
			case "[NETRAL]":
				fmt.Println(color.Yellow + A[i].Category + color.Reset)
			case "[CUKUP NEGATIF]":
				fmt.Println(color.Orange + A[i].Category + color.Reset)
			case "[SANGAT NEGATIF]":
				fmt.Println(color.DarkRed + A[i].Category + color.Reset)
			}
		}

	}
	ui.Pause()

}

func worst() {
	var W global.TabData = global.D
	show(process.InsertionSortData(W, "asc", "Score"), "From Worst", global.NData)
}

func best() {
	var B global.TabData = global.D
	show(process.InsertionSortData(B, "dsc", "Score"), "From Best", global.NData)
}

func az() {
	var AZ global.TabData = process.SelectionSortData(global.D, "asc", "Komentar")
	show(AZ, "From A to Z", global.NData)
}

func za() {
	var ZA global.TabData = process.SelectionSortData(global.D, "dsc", "Komentar")
	show(ZA, "From Z to A", global.NData)
}

// bagian searching
func search() {
	var pilih int
	for pilih != 3 {
		ui.ClearScrn()
		ui.Logo()
		pilih = 0
		fmt.Println("Cari Komentar berdasarkan:")
		fmt.Println("1. Keyword")
		fmt.Println("2. Kategori")
		fmt.Println("3." + color.Red + " Back" + color.Reset)
		fmt.Print("Pilih angka (1/2/3): " + color.Input)
		fmt.Scan(&pilih)
		fmt.Print(color.Reset)
		switch pilih {
		case 1:
			keywordSearch()
		case 2:
			categorySearch()
		}
	}

}

func keywordSearch() {
	var result global.TabData
	var keyword string

	ui.ClearScrn()
	ui.Logo()
	fmt.Scanln()
	fmt.Print("Masukkan kata kunci pencarian: " + color.Input)
	fmt.Scan(&keyword)
	fmt.Print(color.Reset)
	result = global.D
	result = process.SequentialSearch(result, keyword)
	if len(result) == 0 || global.NResult == 0 {
		fmt.Scanln()
		fmt.Println(color.Log + "Komentar dengan kata kunci tersebut tidak ditemukan." + color.Reset)
		ui.Pause()
	} else {
		show(result, "Hasil Pencarian untuk '"+color.Underline+color.Bold+keyword+color.Reset+"'", global.NResult)
	}
}

func categorySearch() {
	var category int
	var A global.TabData

	for category != 6 {
		ui.ClearScrn()
		ui.Logo()
		fmt.Println("Masukkan kategori yang ingin dicari: ")
		fmt.Println("1. SANGAT POSITIF")
		fmt.Println("2. CUKUP POSITIF")
		fmt.Println("3. NETRAL")
		fmt.Println("4. CUKUP NEGATIF")
		fmt.Println("5. SANGAT NEGATIF")
		fmt.Println("6." + color.Red + " Back" + color.Reset)
		fmt.Print("Pilih angka (1/2/3/4/5): " + color.Input)
		fmt.Scan(&category)
		fmt.Print(color.Reset)

		switch category {
		case 1:
			A = process.BinarySearch(global.D, "[SANGAT POSITIF]")
			if global.NResult != 0 {
				show(A, "[SANGAT POSITIF]", global.NResult)
			} else {
				fmt.Scanln()
				fmt.Println(color.Log + "Komentar dengan kategori tersebut tidak ditemukan." + color.Reset)
				ui.Pause()
			}

		case 2:
			A = process.BinarySearch(global.D, "[CUKUP POSITIF]")
			if global.NResult != 0 {
				show(A, "[CUKUP POSITIF]", global.NResult)
			} else {
				fmt.Scanln()
				fmt.Println(color.Log + "Komentar dengan kategori tersebut tidak ditemukan." + color.Reset)
				ui.Pause()
			}

		case 3:
			A = process.BinarySearch(global.D, "[NETRAL]")
			if global.NResult != 0 {
				show(A, "[NETRAL]", global.NResult)
			} else {
				fmt.Scanln()
				fmt.Println(color.Log + "Komentar dengan kategori tersebut tidak ditemukan." + color.Reset)
				ui.Pause()
			}

		case 4:
			A = process.BinarySearch(global.D, "[CUKUP NEGATIF]")
			if global.NResult != 0 {
				show(A, "[CUKUP NEGATIF]", global.NResult)
			} else {
				fmt.Scanln()
				fmt.Println(color.Log + "Komentar dengan kategori tersebut tidak ditemukan." + color.Reset)
				ui.Pause()
			}

		case 5:
			A = process.BinarySearch(global.D, "[SANGAT NEGATIF]")
			if global.NResult != 0 {
				show(A, "[SANGAT NEGATIF]", global.NResult)
			} else {
				fmt.Scanln()
				fmt.Println(color.Log + "Komentar dengan kategori tersebut tidak ditemukan." + color.Reset)
				ui.Pause()
			}
		}
	}

}
