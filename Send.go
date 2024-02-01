package msg

import "encoding/json"

// Function returns a response with data in JSON format.
//{
//	response: { 
//	...data
// 	}
//}
func Send(ctx context, data interface{}) error {

	message, err := json.Marshal(&responseMessage{Response: data})
	if err != nil {
		return err
	}

	ctx.Set("Content-type", "application/json; charset=utf-8")
	_, err = ctx.Write(message)
	return err

}
