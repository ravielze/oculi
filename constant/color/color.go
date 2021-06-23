package color

const (

	// Text Color Group

	RESET        string = "\033[0m"
	BLACK        string = "\033[0;30m"
	RED          string = "\033[0;31m"
	GREEN        string = "\033[0;32m"
	YELLOW       string = "\033[0;33m"
	BLUE         string = "\033[0;34m"
	PURPLE       string = "\033[0;35m"
	CYAN         string = "\033[0;36m"
	WHITE        string = "\033[0;37m"
	GRAY         string = "\033[1;30m"
	LIGHT_RED    string = "\033[1;31m"
	LIME         string = "\033[1;32m"
	YELLOW_BOLD  string = "\033[1;33m"
	LIGHT_BLUE   string = "\033[1;34m"
	LIGHT_PURPLE string = "\033[1;35m"
	AQUA         string = "\033[1;36m"
	WHITE_BOLD   string = "\033[1;37m"

	// Highlight Group
	// Use this after text color group

	HL_BLACK  string = "\033[40m"
	HL_RED    string = "\033[41m"
	HL_GREEN  string = "\033[42m"
	HL_YELLOW string = "\033[43m"
	HL_BLUE   string = "\033[44m"
	HL_PURPLE string = "\033[45m"
	HL_CYAN   string = "\033[46m"
	HL_WHITE  string = "\033[47m"
)
