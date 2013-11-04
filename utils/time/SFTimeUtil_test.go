package SFTimeUtil

import (
	"fmt"
	"testing"
	"time"
)

//
func TestYMDHMSSFormat(t *testing.T) {
	fmt.Println(YMDHMSSFormat(time.Now(), "yyyy-MM-dd hh:mm:ssSSSSSSSSS -0700 MST"))
}

func TestYMDHMSSSignParse(t *testing.T) {
	fmt.Println(YMDHMSSSignFormat(time.Now(), "_${yyyy}"))
}
