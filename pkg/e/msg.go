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
	msg[WrongPasswordFormat] = "密码格式不正确"
	msg[UserNotLogin] = "用户未登录"
	msg[IconTooBig] = "只允许2MB以下的图片作为头像"
	msg[WrongPictureFormat] = "不支持该格式的图片"
}

func GetMsg(code int) string {
	return msg[code]
}
