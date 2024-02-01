package msg

var errorAlreadyExists = serialize(errorData{
	Code:    5,
	Message: "This user is already exists",
})

// Function returns an error with the corresponding 
// status code and message in JSON format.
// { 
//   code: 5, 
//   message: "This user is already exists"
// }
func AlreadyExists(ctx context) error {
	ctx.Set("Content-type", "application/json; charset=utf-8")
	_, err := ctx.Write(errorAlreadyExists)
	return err
}
