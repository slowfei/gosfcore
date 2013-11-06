//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-11-06
//  Update on 2013-11-06
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

// email handle
package SFLog

import (
	"fmt"
	"net/smtp"
	"runtime/debug"
	"strings"
)

const (
	DEFAULT_CONTENT_TYPE = "Content-Type: text/plain; charset=UTF-8"
)

//	appender email config
type AppenderEmailConfig struct {
	Host        string `json:"EmailHost"`        //	smtp host
	User        string `json:"EmailUser"`        // email send user address
	Password    string `json:"EmailPassword"`    //
	FromName    string `json:"EmailFromName"`    // send user show name 如果nil 直接使用EmailUser的值
	ToEmails    string `json:"EmailTo"`          //	send to email address "x1@gmail.com;x2@gmail.com"
	Pattern     string `json:"EmailPattern"`     //	信息内容输出格式
	Subject     string `json:"EmailSubject"`     // email subject
	ContentType string `json:"EmailContentType"` // 默认 "Content-Type: text/plain; charset=UTF-8"
}

// Appender impl email send
type AppenderEmail struct {
}

//	new email impl
func NewAppenderEmail() *AppenderEmail {
	return &AppenderEmail{}
}

//	email send
func (ae *AppenderEmail) emailSend(msg *LogMsg, config *AppenderEmailConfig) {
	host := config.Host
	user := config.User
	from := config.FromName
	pwd := config.Password
	to := config.ToEmails
	subject := config.Subject
	contentType := config.ContentType

	if 0 == len(host) || 0 == len(user) || 0 == len(pwd) || 0 == len(to) {
		fmt.Println("[%v]log email send ftal: Host=%v , User=%v , Password=%v , ToEmails=%v", msg.logGroup, host, user, pwd, to)
		return
	}
	if 0 == len(subject) {
		subject = "Golang Log Info"
	}
	if 0 == len(contentType) {
		contentType = DEFAULT_CONTENT_TYPE
	}
	if 0 == len(from) {
		from = user
	}

	format := config.Pattern
	if 0 == len(format) {
		format = DEFAULT_PATTERN
	}
	msgFormat := logMagFormat(format, msg)

	//	golang smtp
	auth := smtp.PlainAuth("", user, pwd, strings.Split(host, ":")[0])
	bodyMsg := []byte(
		"To:" + to + "\r\n" +
			"From:" + from + "<" + user + ">\r\n" +
			"Subject:" + subject + "\r\n" +
			contentType + "\r\n\r\n" +
			msgFormat)

	err := smtp.SendMail(host, auth, user, strings.Split(to, ";"), bodyMsg)

	if nil != err {
		fmt.Printf("%v\n%s\n", err, debug.Stack())
	}
}

//	#interface impl
//
func (ae *AppenderEmail) Write(msg *LogMsg, configInfo interface{}) {
	if emailConfig, ok := configInfo.(*AppenderEmailConfig); ok {
		ae.emailSend(msg, emailConfig)
	}
}

//	#interface impl
//	name = file
func (af *AppenderEmail) Name() string {
	return VAL_APPENDER_EMAIL
}
