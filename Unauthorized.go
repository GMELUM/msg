package msg

var ErrorUnauthorized = serialize(ErrorData{
	Code:    1,
	Message: "Unauthorized",
})

// Function returns an error with the
// corresponding status code and message in JSON format.
//
//	{
//		code: 1,
//		message: "Unauthorized"
//	}
func Unauthorized(ctx context) error {
	ctx.Set("Content-type", "application/json; charset=utf-8")
	_, err := ctx.Write(ErrorUnauthorized)
	return err
}
