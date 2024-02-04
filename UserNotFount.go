package msg

var errorUserNotFound = serialize(errorData{
	Code:     22,
	Message:  "The user is missing from the store database.",
	Critical: true,
})

//	Use in shop VK callback.
//
//	The user is missing from the store database.
func UserNotFound(ctx context) error {
	ctx.Write(errorUserNotFound)
	return nil
}
