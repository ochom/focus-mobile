package domain

const (
	//usage help for the terminal application
	USAGE = `Invalid input. see usage below:
    usage: ./app action "argument"
    actions: an action can either be 'search' or 'describe'
    example: ./app search "text to search"`

	//the rxnav api route
	BaseRoute = "https://rxnav.nlm.nih.gov"
)
