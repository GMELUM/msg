package msg

var errorProducsNotFount = serialize(errorData{
	Code:     20,
	Message:  "There is no such thing as a product or subscription.",
	Critical: true,
})

//	Use in shop VK callback.
//
//	The item or subscription does not exist.
func ProducsNotFount(ctx context) error {
	ctx.Write(errorProducsNotFount)
	return nil
}
