package process

import (
	"app/global"
)

// mistis ini wak
func HitungScore(A global.TabData) global.TabData {
	var badSc, goodSc, netralSc int
	var bd, gd, ml bool
	var multSc int
	var nKata int
	var h, i, j int
	var kata, s string
	for h = 0; h < global.NData; h++ {

		s = A[h].Komentar + " "
		nKata = 0
		badSc = 0
		goodSc = 0
		netralSc = 0
		multSc = 1
		for i = 0; i < len(s); i++ {
			if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
				kata += string(s[i])
			} else if s[i] == ' ' {
				kata = toLower(kata)
				for j = 0; j < global.NB; j++ {
					if kata == global.Bad[j] {
						badSc -= 2
						bd = true
					}
				}
				for j = 0; j < global.NG; j++ {
					if kata == global.Good[j] {
						goodSc += 2
						gd = true
					}
				}
				for j = 0; j < global.NM; j++ {
					if kata == global.Mult[j] {
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
		A[h].Score = float64(multSc*(badSc+goodSc+netralSc)) / float64(nKata)
	}
	return A
}

func toLower(s string) string {
	var i int
	var converted [global.MAXrune]rune

	for i = 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			converted[i] = rune(s[i] + 32)
		} else {
			converted[i] = rune(s[i])
		}
	}
	return string(converted[:len(s)])
}

func SelectionSortData(A global.TabData, category string, comparator string) global.TabData {
	var parameter bool
	var i, j, idx int
	var temp global.Data

	for i = 0; i < global.NData-1; i++ {
		idx = i
		for j = i + 1; j < global.NData; j++ {
			switch category + comparator {
			case "ascKomentar":
				parameter = A[j].Komentar < A[idx].Komentar
			case "dscKomentar":
				parameter = A[j].Komentar > A[idx].Komentar
			}
			if parameter {
				idx = j
			}
		}
		temp = A[idx]
		A[idx] = A[i]
		A[i] = temp
	}
	return A
}

func InsertionSortData(A global.TabData, category string, comparator string) global.TabData {
	var i, j int
	var temp global.Data
	var parameter, lanjut bool

	for i = 1; i < global.NData; i++ {
		temp = A[i]
		j = i - 1
		lanjut = true

		for j >= 0 && lanjut {
			switch category + comparator {
			case "ascScore":
				parameter = temp.Score < A[j].Score
			case "dscScore":
				parameter = temp.Score > A[j].Score
			}

			if parameter {
				A[j+1] = A[j]
				j--
			} else {
				lanjut = false
			}
		}
		A[j+1] = temp
	}
	return A
}
