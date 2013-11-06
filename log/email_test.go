package SFLog

import (
	"testing"
	"time"
)

func TestEmailLogSend(t *testing.T) {
	ae := NewAppenderEmail()

	config := &AppenderEmailConfig{}
	config.Host = "smtp.126.com:25"
	config.User = "xxx@126.com"
	config.Password = "123"
	config.FromName = "sl-test"
	config.ToEmails = "xxx@gmail.com"

	msg := &LogMsg{}
	msg.dateTime = time.Now()
	msg.fileLine = 26
	msg.filePath = "email_test.go"
	msg.funcName = "TestEmailLogSend(...)"
	msg.logGroup = "logGroup"
	msg.logTag = "logTag"
	msg.msg = "test email send"
	msg.stack = ""

	ae.Write(msg, config)
}
