package msg

var ErrorForbidden = serialize(ErrorData{
	Code:    0,
	Message: "Forbidden",
})

// Function returns an error with the corresponding
// status code and message in JSON format.
//
//	{
//	  code: 0,
//	  message: "Forbidden"
//	}
func Forbidden(ctx context) error {
	ctx.Set("Content-type", "application/json; charset=utf-8")
	_, err := ctx.Write(ErrorForbidden)
	return err
}
