package constants

type Event int

const (
	/*
		Used for Input text element validate value form
	*/
	EVENT_ON_VALIDATE_INPUT Event = iota

	/*
		Used Event When Component was change input
		example
			- Dropdown chooseItem
	*/
	EVENT_ON_SELECT

	/* fill data form connection */
	EVENT_FILL_DATA_FORM_CONNECTION

	/* update data something */
	EVENT_UPDATE
)
