package msg

var errorExpiration = serialize(errorData{
	Code:     8,
	Message:  "Token expiration",
	Critical: true,
})

// Function returns an error with the
// corresponding status code and message in JSON format.
// {
//   code: 8,
//   message: "Token expiration"
// }
func Expiration(ctx context) error {
	ctx.Set("Content-type", "application/json; charset=utf-8")
	_, err := ctx.Write(errorExpiration)
	return err
}
