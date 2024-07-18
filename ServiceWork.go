package msg

var errorServiceWork = serialize(errorData{
	Code:     9,
	Message:  "Technical work is underway",
	Critical: true,
})

// Function returns an error with the
// corresponding status code and message in JSON format.
// {
//   code: 9,
//   message: "Technical work is underway"
// }
func ServiceWork(ctx context) error {
	ctx.Set("Content-type", "application/json; charset=utf-8")
	_, err := ctx.Write(errorServiceWork)
	return err
}
