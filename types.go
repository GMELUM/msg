package msg

import "encoding/json"

type errorMessage struct {
	Error ErrorData `json:"error"`
}

type responseMessage struct {
	Response interface{} `json:"response"`
}

type ErrorData struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Critical bool   `json:"critical,omitempty"`
}

type errorMessageBadRequest struct {
	Error errorDataBadRequest `json:"error"`
}

type errorDataBadRequest struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description,omitempty"`
	Critical    bool   `json:"critical,omitempty"`
}

func serialize(data ErrorData) []byte {
	message, err := json.Marshal(&errorMessage{Error: data})
	if err != nil {
		return []byte{}
	}
	return message
}

type context interface {
	Set(key string, val string)
	Write(p []byte) (int, error)
}
