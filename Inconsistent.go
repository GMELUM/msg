package msg

var ErrorInconsistent = serialize(ErrorData{
	Code:     11,
	Message:  "The request is missing one or more mandatory fields",
	Critical: true,
})

// Use in shop VK callback.
//
// One of the following reasons:
//   - The request parameters do not match the specification.
//   - One or more mandatory fields are missing in the request.
//   - Request integrity errors.
func Inconsistent(ctx context) error {
	ctx.Write(ErrorInconsistent)
	return nil
}
