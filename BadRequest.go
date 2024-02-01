package msg

import "encoding/json"

// Function returns an error with the corresponding
// status code and message in JSON format.
// {
//   code: 2,
//   message: "Bad Request",
//   description: data
// }
func BadRequest(ctx context, msg string, critical ...bool) error {
	var isCritical bool = false
	if len(critical) > 0 {
		isCritical = critical[0]
	}

	message, err := json.Marshal(&errorMessageBadRequest{Error: errorDataBadRequest{
		Code:     2,
		Message:  msg,
		Critical: isCritical,
	}})
	if err != nil {
		return err
	}

	ctx.Set("Content-type", "application/json; charset=utf-8")
	_, err = ctx.Write(message)
	return err
}
