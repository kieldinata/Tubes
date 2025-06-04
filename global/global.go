package global

const MAXdat int = 1000
const MAXrune int = 1000
const MAXword int = 255
const MAXsample int = 120

type Data struct {
	Komentar string
	Score    float64
	Category string
}

type Karakter [MAXrune]rune
type TabData [MAXdat]Data
type TabSamples [MAXsample]string

var D TabData
var Bad TabSamples
var Good TabSamples
var Mult TabSamples

var NData int
var NB, NG, NM int
var NResult int
