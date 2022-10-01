package util

import "fmt"

type ErrorCodeEnum string

func AssertError(err error, errorCode ErrorCodeEnum, statusCode int, errorMsg ...string) {
	if err != nil {
		msg := ""
		if len(errorMsg) == 0 || errorMsg[0] == "" {
			// msg = ErrorMsgs[errorCode] + " - " + err.Error()
			msg = err.Error()
		} else {
			msg = errorMsg[0] + " - " + err.Error()
		}
		ToError(errorCode, statusCode, msg)
	}
}

func ToError(errorCode ErrorCodeEnum, statusCode int, errorMsg string) {
	panic(fmt.Sprint(statusCode) + ":" + string(errorCode) + ":" + errorMsg)
}
