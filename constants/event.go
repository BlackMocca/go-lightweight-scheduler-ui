package constants

type Event int

const (
	/* Used for Input text element validate value form
	 */
	EVENT_ON_VALIDATE_INPUT_TEXT Event = iota

	/* command parent update */
	EVENT_UPDATE
)
