package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvaildParam
	CodeUserExist
	CodeUserNoExist
	CodeInvaildLogin
	CodeInvaildBusy
)

var codeToMsg = map[ResCode]string{
	CodeSuccess:      "success",
	CodeInvaildParam: "Request parameter invalid",
	CodeUserExist:    "User already exists",
	CodeUserNoExist:  "User does not exist",
	CodeInvaildLogin: "Invalid user name or password",
	CodeInvaildBusy:  "Service is busy",
}

func (c ResCode) Msg() string {
	msg, ok := codeToMsg[c]
	if !ok {
		msg = codeToMsg[CodeInvaildBusy]
	}
	return msg
}
