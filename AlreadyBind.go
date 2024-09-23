package msg

var ErrorAlreadyBind = serialize(ErrorData{
	Code:    4,
	Message: "This user is already bind",
})

// Function returns an error with the corresponding
// status code and message in JSON format.
//
//	{
//	  code: 4,
//	  message: "This user is already bind"
//	}
func AlreadyBind(ctx context) error {
	ctx.Set("Content-type", "application/json; charset=utf-8")
	_, err := ctx.Write(ErrorAlreadyBind)
	return err
}
