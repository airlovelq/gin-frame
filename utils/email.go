package utils

import (
	"crypto/tls"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(mailTo string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码

	//mailConn := map[string]string{
	//  "user": "xxx@163.com",
	//  "pass": "your password",
	//  "host": "smtp.163.com",
	//  "port": "465",
	//}

	mailConn := map[string]string{
		"user": GetEnvDefault("Sugar_EMAIL", "email"),
		"pass": GetEnvDefault("Sugar_EMAIL_PASSWORD", "password"),
		"host": GetEnvDefault("Sugar_EMAIL_HOST", "email.com"),
		"port": GetEnvDefault("Sugar_EMAIL_PORT", "25"),
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "Test")) //这种方式可以添加别名，即“XX官方”

	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo)       //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	return err
}
