package msg

import "encoding/json"

// Function returns an error with the corresponding 
// status code and message in JSON format.
// { 
//   code: 2, 
//   message: "Bad Request",
//   description: data
// }
func BadRequest(ctx context, data string) error {

	message, err := json.Marshal(&errorMessageBadRequest{Error: errorDataBadRequest{
		Code:        2,
		Message:     "Bad Request",
		Description: data,
	}})
	if err != nil {
		return err
	}

	ctx.Set("Content-type", "application/json; charset=utf-8")
	_, err = ctx.Write(message)
	return err

}
