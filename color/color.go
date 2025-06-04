package color

const (
	Reset = "\033[0m"

	//warna
	Green      = "\033[32m"
	LightGreen = "\033[92m"
	LimeGreen  = "\033[38;5;154m"
	Yellow     = "\033[33m"
	Orange     = "\033[38;5;208m"
	Red        = "\033[91m"
	DarkRed    = "\033[31m"
	LightBlue  = "\033[94m"
	Teal       = "\033[38;5;30m"
	Turquoise  = "\033[38;5;44m"
	Purple     = "\033[35m"
	Pink       = "\033[38;5;213m"
	Gray       = "\033[90m"
	White      = "\033[97m"
	Black      = "\033[30m"

	//bg
	BgGreen      = "\033[42m"
	BgLightGreen = "\033[102m"
	BgLimeGreen  = "\033[48;5;154m"
	BgYellow     = "\033[43m"
	BgOrange     = "\033[48;5;208m"
	BgRed        = "\033[101m"
	BgDarkRed    = "\033[41m"
	BgLightBlue  = "\033[104m"
	BgTeal       = "\033[48;5;30m"
	BgTurquoise  = "\033[48;5;44m"
	BgPurple     = "\033[45m"
	BgPink       = "\033[48;5;213m"
	BgGray       = "\033[100m"
	BgWhite      = "\033[107m"

	//efek
	Bold      = "\033[1m"
	Underline = "\033[4m"
	Italic    = "\033[3m"
	Blink     = "\033[5m"
	Focus     = "\033[H"
)

var Input = Turquoise
var Warning = Orange
var Log = Gray
var Show = LightBlue
var Title = BgLightBlue + Bold + White
