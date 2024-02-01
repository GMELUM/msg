package msg

var errorOutOfStock = serialize(errorData{
	Code:     21,
	Message:  "Out of stock",
	Critical: true,
})

//	Use in shop VK callback.
//
//	The item is out of stock, user and/or global limits exceeded.
func OutOfStock(ctx context) error {
	ctx.Write(errorOutOfStock)
	return nil
}
