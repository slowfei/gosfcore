package SFLog

import (
	"fmt"
	"testing"
)

//	测试加载配置文件
func TestLoadConfig(t *testing.T) {
	err := LoadConfig("")
	if nil != err {
		fmt.Println(err)
		fmt.Println(_defaultConfig)
	} else {
		fmt.Println(_sharedLogConfig.ChannelSize)
		fmt.Println(_sharedLogConfig.LogTags["global_log_config"].Info)
	}

}
