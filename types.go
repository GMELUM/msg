package msg

import "encoding/json"

type errorMessage struct {
	Error errorData `json:"error"`
}

type responseMessage struct {
	Response interface{} `json:"response"`
}

type errorData struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Critical bool   `json:"critical"`
}

type errorMessageBadRequest struct {
	Error errorDataBadRequest `json:"error"`
}

type errorDataBadRequest struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
	Critical    bool   `json:"critical"`
}

func serialize(data errorData) []byte {
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
