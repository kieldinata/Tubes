package color

const (
	Reset = "\033[0m"

	//warna
	DarkRed     = "\033[31m"
	Yellow      = "\033[33m"
	Green       = "\033[32m"
	LimeGreen   = "\033[38;5;154m"
	Orange      = "\033[38;5;208m"
	Red         = "\033[91m"
	LightBlue   = "\033[94m"
	Turquoise   = "\033[38;5;44m"
	Gray        = "\033[90m"
	White       = "\033[97m"
	BgLightBlue = "\033[104m"

	//efek
	Bold      = "\033[1m"
	Underline = "\033[4m"
	Blink     = "\033[5m"

	//untuk ui
	Input   = Turquoise
	Warning = Orange
	Show    = LightBlue
	Title   = BgLightBlue + Bold + White
	Result  = Underline + Bold
)
