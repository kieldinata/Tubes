package process

import (
	"app/global"
)

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

// Ilmu mistis, jangan disentuh tanpa di copy!!!!
func HitungScore(s string) float64 {
	var badSc, goodSc, netralSc int
	var bd, gd, ml bool
	var multSc int = 1
	var nKata int
	var i, j int
	var score float64
	var kata string

	s += " "
	for i = 0; i < len(s); i++ {
		if s[i] != ' ' {
			kata += string(s[i])
			kata = toLower(kata)
		} else {
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
	score = float64(multSc*(badSc+goodSc+netralSc)) / float64(nKata+1)
	return score
}
