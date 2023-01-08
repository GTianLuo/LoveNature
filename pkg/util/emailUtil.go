package util

import (
	"fmt"
	"github.com/jordan-wright/email"
	"lovenature/log"
	"net/smtp"
)

func formatCode(code string) string {
	return fmt.Sprintf("\t[Love Nature] 欢迎使用Love Nature app, 您的验证码为:%s (一分钟后过期)。\n\n\t请妥善保管，不要告诉他人！", code)
}

func SendCode(code string, e string) {

	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "g2985496686@163.com"

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{e}

	// 设置主题
	em.Subject = "Love Nature 验证码"

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte(formatCode(code))

	//设置服务器相关的配置
	err := em.Send("smtp.163.com:25", smtp.PlainAuth("", "g2985496686@163.com", "KQKSJMOKJGHHDJUT", "smtp.163.com"))
	if err != nil {
		log.Error(err)
	}
}
