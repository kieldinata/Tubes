package sample

import (
	"app/global"
	"app/ui"
	"fmt"
)

func Menu() {
	var pilih int
	for pilih != 5 {
		ui.ClearScrn()
		pilih = 0
		fmt.Println("-------------------------------")
		fmt.Println("SAMPLE MENU")
		fmt.Println("-------------------------------")
		fmt.Println("1. Tunjukkan Semua Sample")
		fmt.Println("2. Tambahkan Sample")
		fmt.Println("3. Edit Sample")
		fmt.Println("4. Hapus Sample")
		fmt.Println("5. Back")
		fmt.Print("Pilih dengan angka (1/2/3/4): ")
		fmt.Scanf("%d", &pilih)

		switch pilih {
		case 1:
			ui.ClearScrn()
			show()
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
	fmt.Println("List Sample Bad:")
	for i = 0; i < global.NB; i++ {
		fmt.Printf("%d. %-10s\n", i+1, global.Bad[i])
	}
	fmt.Println("\n\nList Sample Good:")
	for i = 0; i < global.NG; i++ {
		fmt.Printf("%d. %-10s\n", i+1, global.Good[i])
	}
	fmt.Println("\n\nList Sample Multiplier:")
	for i = 0; i < global.NM; i++ {
		fmt.Printf("%d. %-10s \n", i+1, global.Mult[i])
	}
	ui.Pause()
}

func insert() {
	var tipe int
	var input string

	fmt.Scanln()
	fmt.Println("Pilih tipe sample:")
	fmt.Println("1. Bad")
	fmt.Println("2. Good")
	fmt.Println("3. Multiplier")
	fmt.Print("tipe pilihan (1/2/3): ")
	fmt.Scan(&tipe)
	ui.ClearScrn()
	fmt.Println("Masukkan sample anda diakhiri dengan # sebagai sentinel:")
	for input != "#" {
		fmt.Scan(&input)
		if input != "#" {
			if tipe == 1 {
				global.Bad[global.NB] = input
				global.NB++
			} else if tipe == 2 {
				global.Good[global.NG] = input
				global.NG++
			} else if tipe == 3 {
				global.Mult[global.NM] = input
				global.NM++
			}
		}
	}
	ui.Pause()
}

func edit() {
	var tipe int

	fmt.Scanln()
	fmt.Println("Pilih tipe sample yang ingin di hapus:")
	fmt.Println("1. Bad")
	fmt.Println("2. Good")
	fmt.Println("3. Multiplier")
	fmt.Print("tipe pilihan (1/2/3): ")
	fmt.Scan(&tipe)
	ui.ClearScrn()
	if tipe == 1 {
		editPerType(&global.Bad, global.NB)
	} else if tipe == 2 {
		editPerType(&global.Good, global.NG)
	} else if tipe == 3 {
		editPerType(&global.Mult, global.NM)
	} else {
		fmt.Print("Tipe tidak valid!")
		ui.Pause()
	}
}

func editPerType(x *global.TabSamples, nX int) {
	var edNum int

	fmt.Print("Pilih nomor yang ingin di edit: ")
	fmt.Scan(&edNum)
	if edNum-1 < 0 || edNum-1 >= nX {
		fmt.Print("Nomor tidak valid!")
	} else {
		fmt.Println("Masukkan sample  versi terbaru: ")
		fmt.Scan(&x[edNum-1])
		fmt.Print("Sample berhasil di ubah!")
	}
	ui.Pause()
}

func delete() {
	var tipe int

	fmt.Scanln()
	fmt.Println("Pilih tipe sample yang ingin di hapus:")
	fmt.Println("1. Bad")
	fmt.Println("2. Good")
	fmt.Println("3. Multiplier")
	fmt.Print("tipe pilihan (1/2/3): ")
	fmt.Scan(&tipe)
	ui.ClearScrn()
	if tipe == 1 {
		deletePerType(&global.Bad, &global.NB)
	} else if tipe == 2 {
		deletePerType(&global.Good, &global.NG)
	} else if tipe == 3 {
		deletePerType(&global.Mult, &global.NM)
	} else {
		fmt.Print("Tipe tidak valid.")
		ui.Pause()
	}
}

func deletePerType(x *global.TabSamples, nX *int) {
	var i int
	var delNum int

	for i = 0; i < *nX; i++ {
		fmt.Printf("%d. %-10s ", i+1, x[i])
		if (i+1)%10 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	fmt.Print("Pilih nomor yang ingin di hapus: ")
	fmt.Scan(&delNum)
	if delNum-1 < 0 || delNum-1 >= *nX {
		fmt.Print("Nomor tidak valid!")
	} else {
		for i = delNum - 1; i < *nX-1; i++ {
			x[i] = x[i+1]
		}
		*nX--
		fmt.Printf("Nomor %d berhasil di hapus", delNum)
	}
	ui.Pause()
}
