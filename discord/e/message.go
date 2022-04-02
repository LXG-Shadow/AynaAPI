package e

import "errors"

var CODE_MESSAGE = map[int]string{
	SUCCESS:                            "Success",
	INTERNAL_ERROR:                     "Internal err",
	MUSIC_ERROR_PROVIDER_NOT_AVAILABLE: "MUSIC_ERROR_PROVIDER_NOT_AVAILABLE",
	MUSIC_ERROR_SEARCH_FAIL:            "MUSIC_ERROR_SEARCH_FAIL",
	MUSIC_ERROR_INITIALIZE_FAIL:        "MUSIC_ERROR_INITIALIZE_FAIL",
	MUSIC_ERROR_GET_DATA_FAIL:          "MUSIC_ERROR_GET_DATA_FAIL",

	CONTEXT_ERROR_NOT_IN_VOICE_CHANNEL: "CONTEXT_ERROR_NOT_IN_VOICE_CHANNEL",

	MUSICBOT_ERROR_QUEUE_EMPTY:             "MUSICBOT_ERROR_QUEUE_EMPTY",
	MUSICBOT_ERROR_ALREADY_PAUSED_UNPAUSED: "MUSICBOT_ERROR_ALREADY_PAUSED_UNPAUSED",
}

// GetMessage get error information based on Code
func GetMessage(code int) string {
	msg, ok := CODE_MESSAGE[code]
	if ok {
		return msg
	}
	return CODE_MESSAGE[INTERNAL_ERROR]
}

func NewError(code int) error {
	return errors.New(GetMessage(code))
}
