package e

var msg map[int]string

func init() {
	msg = make(map[int]string)
	msg[Success] = "ok"
	msg[Error] = "服务器内部错误"

	msg[InvalidEmail] = "邮箱格式不正确"
	msg[RepeatSending] = "发送过于频繁"
	msg[InvalidParam] = "参数解析异常"
	msg[WrongCode] = "验证码错误"
	msg[RepeatRegister] = "该邮箱已经注册"
	msg[WrongAccountOrPassword] = "账号或密码不正确"
}

func GetMsg(code int) string {
	return msg[code]
}
