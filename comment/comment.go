package comment

import (
	"app/global"
	"app/process"
	"app/ui"
	"fmt"
)

func Menu() {
	var pilih int
	for pilih != 5 {
		ui.ClearScrn()
		pilih = 0
		fmt.Println("-------------------------------")
		fmt.Println("COMMENT MENU")
		fmt.Println("-------------------------------")
		fmt.Println("1. List Semua Komentar")
		fmt.Println("2. Tambahkan Komentar")
		fmt.Println("3. Edit Komentar")
		fmt.Println("4. Hapus Komentar")
		fmt.Println("5. Back")
		fmt.Print("Pilih dengan angka (1/2/3/4): ")
		fmt.Scanf("%d", &pilih)

		switch pilih {
		case 1:
			ui.ClearScrn()
			show()
			ui.Pause()
		case 2:
			ui.ClearScrn()
			insert()
		case 3:
			ui.ClearScrn()
			edit()
		case 4:
			ui.ClearScrn()
			delete()
		}
	}
}

func show() {
	var i int

	fmt.Scanln()
	fmt.Println("List Semua Komentar:")
	for i = 0; i < global.NData; i++ {
		fmt.Printf("%d. %s\n", i+1, global.D[i].Komentar)
		global.D[i].Score = process.HitungScore(global.D[i].Komentar)
	}
}

func insert() {
	var i int
	var char rune
	var teks string

	fmt.Scanln()
	fmt.Println("Masukkan komentar diakhiri dengan enter:")
	for char != '\n' && i < global.MAXrune {
		fmt.Scanf("%c", &char)
		if char != '\n' && char != '\r' {
			teks += string(char)
			i++
		}
	}
	global.D[global.NData].Komentar = teks
	global.NData++
	fmt.Printf("Komentar ke-%d berhasil ditambahkan", global.NData)
	ui.Pause()
}

func edit() {
	var editNum, i int
	var char rune
	var teks string = ""

	show()
	fmt.Print("\nPilih nomor yang ingin di edit: ")
	fmt.Scan(&editNum)
	fmt.Scanln()
	fmt.Println("Tuliskan komentar setelah di edit: ")
	for char != '\n' && i < global.MAXrune {
		fmt.Scanf("%c", &char)
		if char != '\n' && char != '\r' {
			teks += string(char)
			i++
		}
	}
	global.D[editNum-1].Komentar = teks
}

func delete() {
	var delNUm int
	var i int

	show()
	fmt.Print("Pilih nomor yang ingin di hapus: ")
	fmt.Scan(&delNUm)

	for i = delNUm - 1; i < global.NData-1; i++ {
		global.D[i].Komentar = global.D[i+1].Komentar
		global.D[i].Score = global.D[i+1].Score
	}
	global.NData--
	fmt.Printf("Nomor %d telah berhasil di hapus.", delNUm)
	ui.Pause()
}
