package sample

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
		fmt.Println(color.Title + "<APP - SAMPLE...........................>\n" + color.Reset)
		fmt.Println("1. Tunjukkan Semua Sample")
		fmt.Println("2. Tambahkan Sample")
		fmt.Println("3. Edit Sample")
		fmt.Println("4. Hapus Sample")
		fmt.Printf("5. %sBack%s\n", color.Red, color.Reset)
		fmt.Print("Pilih dengan angka (1/2/3/4): " + color.Input)
		fmt.Scan(&pilih)
		fmt.Print(color.Reset)

		switch pilih {
		case 1:
			ui.ClearScrn()
			ui.Logo()
			show()
		case 2:
			ui.ClearScrn()
			ui.Logo()
			insert()
		case 3:
			ui.ClearScrn()
			ui.Logo()
			edit()
		case 4:
			ui.ClearScrn()
			ui.Logo()
			delete()
		}
	}
}

func show() {
	var i int

	fmt.Scanln()
	fmt.Println(color.Title + "<APP - SAMPLE - SHOW....................>\n" + color.Reset)
	if global.NB != 0 {
		fmt.Println(color.DarkRed + "List Sample Bad:" + color.Reset)
		for i = 0; i < global.NB; i++ {
			fmt.Printf("%d. %-10s\n", i+1, global.Bad[i])
		}
	} else {
		fmt.Println(color.Gray + "Sample [bad] belum ada, silahkan isi sample terlebih dahulu." + color.Reset)
	}
	if global.NG != 0 {
		fmt.Println(color.Green + "\n\nList Sample Good:" + color.Reset)
		for i = 0; i < global.NG; i++ {
			fmt.Printf("%d. %-10s\n", i+1, global.Good[i])
		}
	} else {
		fmt.Println(color.Gray + "Sample [good] belum ada, silahkan isi sample terlebih dahulu." + color.Reset)
	}
	if global.NM != 0 {
		fmt.Println(color.Yellow + "\n\nList Sample Multiplier:" + color.Reset)
		for i = 0; i < global.NM; i++ {
			fmt.Printf("%d. %-10s \n", i+1, global.Mult[i])
		}
	} else {
		fmt.Println(color.Gray + "Sample [mult] belum ada, silahkan isi sample terlebih dahulu." + color.Reset)
	}
	ui.Pause()
}

func insert() {
	var tipe int
	var input string

	fmt.Println(color.Title + "<APP - SAMPLE - ADD.....................>\n" + color.Reset)
	fmt.Scanln()
	fmt.Println("Pilih tipe sample:")
	fmt.Println("1. Bad")
	fmt.Println("2. Good")
	fmt.Println("3. Multiplier")
	fmt.Print("tipe pilihan (1/2/3): " + color.Input)
	fmt.Scan(&tipe)
	fmt.Print(color.Reset)
	ui.ClearScrn()
	ui.Logo()
	fmt.Println(color.Title + "<APP - SAMPLE - ADD.....................>\n" + color.Reset)
	fmt.Println("Masukkan sample anda diakhiri dengan # sebagai sentinel:" + color.Input)
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
	fmt.Print(color.Reset)
	ui.Pause()
}

func edit() {
	var tipe int

	fmt.Println(color.Title + "<APP - SAMPLE - EDIT....................>\n" + color.Reset)
	fmt.Scanln()
	fmt.Println("Pilih tipe sample yang ingin di hapus:")
	fmt.Println("1. Bad")
	fmt.Println("2. Good")
	fmt.Println("3. Multiplier")
	fmt.Print("tipe pilihan (1/2/3): " + color.Input)
	fmt.Scan(&tipe)
	fmt.Print(color.Reset)
	ui.ClearScrn()
	if tipe == 1 {
		editPerType(&global.Bad, global.NB)
	} else if tipe == 2 {
		editPerType(&global.Good, global.NG)
	} else if tipe == 3 {
		editPerType(&global.Mult, global.NM)
	} else {
		fmt.Println(color.Warning + "Tipe tidak valid!" + color.Reset)
		ui.Pause()
	}
}

func editPerType(x *global.TabSamples, nX int) {
	var edNum int
	var i int

	ui.Logo()
	fmt.Println(color.Title + "<APP - SAMPLE - DELETE..................>\n" + color.Reset)
	for i = 0; i < nX; i++ {
		fmt.Printf("%d. %-10s\n", i+1, x[i])
	}
	fmt.Print("Pilih nomor yang ingin di edit: " + color.Input)
	fmt.Scan(&edNum)
	fmt.Print(color.Reset)
	if edNum-1 < 0 || edNum-1 >= nX {
		fmt.Println(color.Warning + "Nomor tidak valid!" + color.Reset)
	} else {
		fmt.Println("Masukkan sample  versi terbaru: " + color.Input)
		fmt.Scan(&x[edNum-1])
		fmt.Print(color.Log + "Sample berhasil di ubah!" + color.Reset)
	}
	ui.Pause()
}

func delete() {
	var tipe int

	fmt.Println(color.Title + "<APP - SAMPLE - DELETE..................>\n" + color.Reset)
	fmt.Scanln()
	fmt.Println("Pilih tipe sample yang ingin di hapus:")
	fmt.Println("1. Bad")
	fmt.Println("2. Good")
	fmt.Println("3. Multiplier")
	fmt.Print("tipe pilihan (1/2/3): " + color.Input)
	fmt.Scan(&tipe)
	fmt.Print(color.Reset)
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

	ui.Logo()
	fmt.Println(color.Title + "<APP - SAMPLE - DELETE..................>\n" + color.Reset)
	for i = 0; i < *nX; i++ {
		fmt.Printf("%d. %-10s\n", i+1, x[i])
	}
	fmt.Print("Pilih nomor yang ingin di hapus: " + color.Input)
	fmt.Scan(&delNum)
	fmt.Print(color.Reset)
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
