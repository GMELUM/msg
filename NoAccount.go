package msg

var errorNoAccount = serialize(errorData{
	Code:    3,
	Message: "This account is not linked to any user page",
})

// Function returns an error with the 
// corresponding status code and message in JSON format.
// { 
//   code: 3, 
//   message: "This account is not linked to any user page"
// }
func NoAccount(ctx context) error {
	ctx.Set("Content-type", "application/json; charset=utf-8")
	_, err := ctx.Write(errorNoAccount)
	return err
}
