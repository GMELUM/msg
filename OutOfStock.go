package msg

var ErrorOutOfStock = serialize(ErrorData{
	Code:     21,
	Message:  "Out of stock",
	Critical: true,
})

// Use in shop VK callback.
//
// The item is out of stock, user and/or global limits exceeded.
func OutOfStock(ctx context) error {
	ctx.Write(ErrorOutOfStock)
	return nil
}
