package msg

import "encoding/json"

func CustomError(ctx context, code int, msg string, critical ...bool) error {
	var isCritical bool = false
	if len(critical) > 0 {
		isCritical = critical[0]
	}

	message, err := json.Marshal(&errorMessageBadRequest{Error: errorDataBadRequest{
		Code:     code,
		Message:  msg,
		Critical: isCritical,
	}})
	if err != nil {
		return err
	}

	ctx.Set("Content-type", "application/json; charset=utf-8")
	_, err = ctx.Write(message)
	return err
}
