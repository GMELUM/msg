package msg

var ErrorManyRequest = serialize(ErrorData{
	Code:    6,
	Message: "To many request",
})

// Function returns an error with the
// corresponding status code and message in JSON format.
//
//	{
//	  code: 6,
//	  message: "To many request"
//	}
func ManyRequest(ctx context) error {
	ctx.Set("Content-type", "application/json; charset=utf-8")
	_, err := ctx.Write(ErrorManyRequest)
	return err
}
