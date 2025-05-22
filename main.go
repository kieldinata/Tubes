package main

import (
	"fmt"
	"os"
	"os/exec"
)

const MAXdat int = 1000
const MAXrune int = 1000
const MAXword int = 255
const MAXsample int = 120

type data struct {
	user     string
	komentar string
	score    float64
}

type karakter [MAXrune]rune
type tabData [MAXdat]data
type tabSamples [MAXsample]string

var d tabData
var bad tabSamples
var good tabSamples
var mult tabSamples

var nData int
var nB, nG, nM int

// FUNGSI UI===============================================================================>
func ClearScrn() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Pause() {
	fmt.Print("\n\nTekan enter untuk melanjutkan...")
	fmt.Scanln()
}

// FUNGSI UI===============================================================================>
// x
// x
// MENU SECTION============================================================================>
func main() {
	var pilih int
	for pilih != 4 {
		pilih = 0
		ClearScrn()
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
			ComMenu()
		case 2:
			SampleMenu()
		case 3:
			Stats()
		}
	}
}

func ComMenu() {
	var pilih int
	for pilih != 5 {
		ClearScrn()
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
			ClearScrn()
			ComShow()
			Pause()
		case 2:
			ClearScrn()
			ComIn()
		case 3:
			ClearScrn()
			ComEd()
		case 4:
			ClearScrn()
			ComDel()
		}
	}
}

func SampleMenu() {
	var pilih int
	for pilih != 5 {
		ClearScrn()
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
			ClearScrn()
			SamShow()
		case 2:
			ClearScrn()
			SamIn()
		case 3:
			ClearScrn()
			SamEd()
		case 4:
			ClearScrn()
			SamDel()
		}
	}
}

// MENU SECTION============================================================================>
// x
// x
// COMMENT SECTION=========================================================================>
func ComShow() {
	var i int

	fmt.Scanln()
	fmt.Println("List Semua Komentar:")
	for i = 0; i < nData; i++ {
		fmt.Printf("%d. %s\n", i+1, d[i].komentar)
		d[i].score = hitungScore(d[i].komentar)
	}
}
func ComIn() {
	var i int
	var char rune
	var teks string

	fmt.Scanln()
	fmt.Println("Masukkan komentar diakhiri dengan enter:")
	for char != '\n' && i < MAXrune {
		fmt.Scanf("%c", &char)
		if char != '\n' && char != '\r' {
			teks += string(char)
			i++
		}
	}
	d[nData].komentar = teks
	nData++
	fmt.Printf("Komentar ke-%d berhasil ditambahkan", nData)
	Pause()
}

func ComEd() {
	var editNum, i int
	var char rune
	var teks string = ""

	ComShow()
	fmt.Print("\nPilih nomor yang ingin di edit: ")
	fmt.Scan(&editNum)
	fmt.Scanln()
	fmt.Println("Tuliskan komentar setelah di edit: ")
	for char != '\n' && i < MAXrune {
		fmt.Scanf("%c", &char)
		if char != '\n' && char != '\r' {
			teks += string(char)
			i++
		}
	}
	d[editNum-1].komentar = teks
}

func ComDel() {
	var delNUm int
	var i int

	ComShow()
	fmt.Print("Pilih nomor yang ingin di hapus: ")
	fmt.Scan(&delNUm)

	for i = delNUm - 1; i < nData-1; i++ {
		d[i].komentar = d[i+1].komentar
		d[i].score = d[i+1].score
	}
	nData--
	fmt.Printf("Nomor %d telah berhasil di hapus.", delNUm)
	Pause()
}

func toLower(s string) string {
	var i int
	var converted [MAXrune]rune

	for i = 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			converted[i] = rune(s[i] + 32)
		} else {
			converted[i] = rune(s[i])
		}
	}
	return string(converted[:len(s)])
}

// Ilmu mistis, jangan disentuh tanpa di copy!!!!
func hitungScore(s string) float64 {
	var badSc, goodSc, netralSc int
	var bd, gd, ml bool
	var multSc int = 1
	var nKata int
	var i, j int
	var score float64
	var kata string

	for i = 0; i < len(s); i++ {
		if s[i] != ' ' {
			kata += string(s[i])
			kata = toLower(kata)
		} else {
			for j = 0; j < nB; j++ {
				if kata == bad[j] {
					badSc -= 2
					bd = true
				}
			}
			for j = 0; j < nG; j++ {
				if kata == good[j] {
					goodSc += 2
					gd = true
				}
			}
			for j = 0; j < nM; j++ {
				if kata == mult[j] {
					multSc++
					ml = true
				}
			}
			if !bd && !gd && !ml {
				netralSc++
			}
			kata = ""
			nKata++
			bd = false
			gd = false
			ml = false
		}
	}
	score = float64(multSc*(badSc+goodSc+netralSc)) / float64(nKata+1) //rumus penting
	return score
}

// COMMENT SECTION=========================================================================>
// x
// x
// SAMPLE SECTION==========================================================================>
func SamShow() {
	var i int

	fmt.Scanln()
	fmt.Println("List Sample Bad:")
	for i = 0; i < nB; i++ {
		fmt.Printf("%d. %-10s\n", i+1, bad[i])

	}
	fmt.Println("\n\nList Sample Good:")
	for i = 0; i < nG; i++ {
		fmt.Printf("%d. %-10s\n", i+1, good[i])
	}
	fmt.Println("\n\nList Sample Multiplier:")
	for i = 0; i < nM; i++ {
		fmt.Printf("%d. %-10s \n", i+1, mult[i])
	}
	Pause()
}

func SamIn() {
	var tipe int
	var input string

	fmt.Scanln()
	fmt.Println("Pilih tipe sample:")
	fmt.Println("1. Bad")
	fmt.Println("2. Good")
	fmt.Println("3. Multiplier")
	fmt.Print("tipe pilihan (1/2/3): ")
	fmt.Scan(&tipe)
	ClearScrn()
	fmt.Println("Masukkan sample anda diakhiri dengan # sebagai sentinel:")
	for input != "#" {
		fmt.Scan(&input)
		if input != "#" {
			if tipe == 1 {
				bad[nB] = input
				nB++
			} else if tipe == 2 {
				good[nG] = input
				nG++
			} else if tipe == 3 {
				mult[nM] = input
				nM++
			}
		}
	}
	Pause()

}

func SamEd() {
	var tipe int

	fmt.Scanln()
	fmt.Println("Pilih tipe sample yang ingin di hapus:")
	fmt.Println("1. Bad")
	fmt.Println("2. Good")
	fmt.Println("3. Multiplier")
	fmt.Print("tipe pilihan (1/2/3): ")
	fmt.Scan(&tipe)
	ClearScrn()
	if tipe == 1 {
		SampEdPerType(&bad, nB)
	} else if tipe == 2 {
		SampEdPerType(&good, nG)
	} else if tipe == 3 {
		SampEdPerType(&mult, nM)
	} else {
		fmt.Print("Tipe tidak valid!")
		Pause()
	}

}

func SampEdPerType(x *tabSamples, nX int) {
	var edNum int

	fmt.Print("Pilih nomor yang ingin di edit: ")
	fmt.Scan(&edNum)
	if edNum-1 < 0 && edNum-1 >= nX {
		fmt.Print("Nomor tidak valid!")
	} else {
		fmt.Println("Masukkan sample  versi terbaru: ")
		fmt.Scan(&x[edNum-1])
		fmt.Print("Sample berhasil di ubah!")
	}
	Pause()

}

func SamDel() {
	var tipe int

	fmt.Scanln()
	fmt.Println("Pilih tipe sample yang ingin di hapus:")
	fmt.Println("1. Bad")
	fmt.Println("2. Good")
	fmt.Println("3. Multiplier")
	fmt.Print("tipe pilihan (1/2/3): ")
	fmt.Scan(&tipe)
	ClearScrn()
	if tipe == 1 {
		SampDelPerType(&bad, &nB)
	} else if tipe == 2 {
		SampDelPerType(&good, &nG)
	} else if tipe == 3 {
		SampDelPerType(&mult, &nM)
	} else {
		fmt.Print("Tipe tidak valid.")
		Pause()
	}
}

func SampDelPerType(x *tabSamples, nX *int) {
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
	if delNum-1 < 0 && delNum-1 >= *nX {
		fmt.Print("Nomor tidak valid!")
	} else {
		for i = delNum - 1; i < *nX-1; i++ {
			x[i] = x[i+1]
		}
		*nX--
		fmt.Print("Nomor %d berhasil di hapus", delNum)
	}
	Pause()

}

// SAMPLE SECTION==========================================================================>
// x
// x
// STAT SECTION============================================================================>
func Stats() {
	var i int
	var komentar string
	var scr float64

	ClearScrn()
	fmt.Scanln()
	fmt.Println("List Semua Komentar:")
	for i = 0; i < nData; i++ {
		komentar = d[i].komentar
		scr = hitungScore(komentar)
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
	Pause()
}

// STAT SECTION============================================================================>
