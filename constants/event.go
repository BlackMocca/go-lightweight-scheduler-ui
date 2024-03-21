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

	/* clear data fill data from connection */
	EVENT_CLEAR_DATA_FROM_CONNECTION

	/* delete data fill data from connection */
	EVENT_DELETE_DATA_FROM_CONNECTION

	/* sending data into parent */
	EVENT_GET_DATA

	/* update data something */
	EVENT_UPDATE
)
