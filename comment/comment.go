package comment

import (
	"app/color"
	"app/global"
	"app/ui"
	"fmt"
)

func Menu() {
	var pilih int
	for pilih != 5 {
		ui.ClearScrn()
		pilih = 0
		ui.Logo()
		fmt.Println(color.Title + "<APP - COMMENT..........................>" + color.Reset)
		fmt.Println("")
		fmt.Println("1. List Semua Komentar")
		fmt.Println("2. Tambahkan Komentar")
		fmt.Println("3. Edit Komentar")
		fmt.Println("4. Hapus Komentar")
		fmt.Printf("5. %sBack%s\n", color.Red, color.Reset)
		fmt.Print("Pilih dengan angka (1/2/3/4): " + color.Input)
		fmt.Scan(&pilih)
		fmt.Print(color.Reset)

		switch pilih {
		case 1:
			ui.ClearScrn()
			ui.Logo()
			fmt.Println(color.Title + "<APP - COMMENT - SHOW...................>" + color.Reset)
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
	if global.NData != 0 {
		fmt.Println("List Semua Komentar:")
		for i = 0; i < global.NData; i++ {
			fmt.Printf("%d. %s%s%s\n", i+1, color.Show, global.D[i].Komentar, color.Reset)
		}
	} else {
		fmt.Println(color.Log + "Belum ada komentar, silahkan tambahkan komentar." + color.Reset)
	}
}

func insert() {
	var i int
	var char rune
	var teks string

	fmt.Scanln()
	ui.Logo()
	fmt.Println(color.Title + "<APP - COMMENT - ADD....................>" + color.Reset)
	fmt.Println("Masukkan komentar diakhiri dengan enter:" + color.Input)
	for char != '\n' && i < global.MAXrune {
		fmt.Scanf("%c", &char)
		if char != '\n' && char != '\r' {
			teks += string(char)
			i++
		}
	}
	fmt.Print(color.Reset)
	global.D[global.NData].Komentar = teks
	global.NData++
}

func edit() {
	var editNum, i int
	var char rune
	var teks string = ""

	ui.Logo()
	fmt.Println(color.Title + "<APP - COMMENT - EDIT...................>" + color.Reset)
	show()
	if global.NData > 0 {
		fmt.Print("\nPilih nomor yang ingin di edit: " + color.Input)
		fmt.Scan(&editNum)
		fmt.Print(color.Reset)
		fmt.Scanln()
		fmt.Println("Tuliskan komentar setelah di edit: " + color.Input)
		for char != '\n' && i < global.MAXrune {
			fmt.Scanf("%c", &char)
			if char != '\n' && char != '\r' {
				teks += string(char)
				i++
			}
		}
		fmt.Print(color.Reset)
		global.D[editNum-1].Komentar = teks
	} else {
		ui.Pause()
	}
}

func delete() {
	var delNUm int
	var i int

	ui.Logo()
	fmt.Println(color.Title + "<APP - COMMENT - DELETE.................>" + color.Reset)
	show()
	if global.NData > 0 {
		fmt.Print("\nPilih nomor yang ingin di hapus: " + color.Input)
		fmt.Scan(&delNUm)
		fmt.Print(color.Reset)

		for i = delNUm - 1; i < global.NData-1; i++ {
			global.D[i].Komentar = global.D[i+1].Komentar
			global.D[i].Score = global.D[i+1].Score
		}
		global.NData--
	} else {
		ui.Pause()
	}
}
