package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvaildParam
	CodeUserExist
	CodeUserNoExist
	CodeInvaildLogin
	CodeInvaildBusy

	CodeAuthIsEmpty
	CodeAuthFormatError
	CodeInvalidToken
)

var codeToMsg = map[ResCode]string{
	CodeSuccess:      "success",
	CodeInvaildParam: "Request parameter invalid",
	CodeUserExist:    "User already exists",
	CodeUserNoExist:  "User does not exist",
	CodeInvaildLogin: "Invalid user name or password",
	CodeInvaildBusy:  "Service is busy",

	CodeAuthIsEmpty:     "auth is empty",
	CodeAuthFormatError: "error in auth format",
	CodeInvalidToken:    "invaild token",
}

func (c ResCode) Msg() string {
	msg, ok := codeToMsg[c]
	if !ok {
		msg = codeToMsg[CodeInvaildBusy]
	}
	return msg
}
