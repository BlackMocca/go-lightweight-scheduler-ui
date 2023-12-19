package constants

type Event int

const (
	/*
		Used for Input text element validate value form
	*/
	EVENT_ON_VALIDATE_INPUT_TEXT Event = iota

	/*
		Used Event When Component was change input
		example
			- Dropdown chooseItem
	*/
	EVENT_ON_SELECT

	/* command parent update */
	EVENT_UPDATE
)
