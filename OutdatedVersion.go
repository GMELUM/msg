package msg

var ErrorOutdatedVersion = serialize(ErrorData{
	Code:     10,
	Message:  "Outdated version of the application",
	Critical: true,
})

// Function returns an error with the
// corresponding status code and message in JSON format.
//
//	{
//	  code: 10,
//	  message: "Outdated version of the application"
//	}
func OutdatedVersion(ctx context) error {
	ctx.Set("Content-type", "application/json; charset=utf-8")
	_, err := ctx.Write(ErrorOutdatedVersion)
	return err
}
