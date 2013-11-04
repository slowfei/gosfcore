package SFLog

import (
	"testing"
	"time"
)

func TestGetFile(t *testing.T) {
	af := NewAppenderFile()
	af.getFile("file_${yyyy}-${MM}-${dd}.log", time.Now())
}
