package msg

var ErrorInvalidFields = serialize(ErrorData{
	Code:     7,
	Message:  "Required fields are missing or their type does not match the declared one",
	Critical: true,
})

// Function returns an error with the corresponding
// status code and message in JSON format.
//
//	{
//	  code: 7,
//	  message: "Required fields are missing or their type does not match the declared one"
//	}
func InvalidFields(ctx context) error {
	ctx.Set("Content-type", "application/json; charset=utf-8")
	_, err := ctx.Write(ErrorInvalidFields)
	return err
}
