package SFLog

import (
	"fmt"
	"testing"
)

//	测试加载配置文件
func testLoadConfig(t *testing.T) {
	err := LoadConfig("")
	if nil != err {
		fmt.Println(err)
		fmt.Println(_defaultConfig)
	} else {
		fmt.Println(_sharedLogConfig.LogGroups[KEY_GLOBAL_GROUP_LOG_CONFIG].AppenderConsoleConfig)
		fmt.Println(_sharedLogConfig.LogGroups[KEY_GLOBAL_GROUP_LOG_CONFIG].Info.AppenderConsoleConfig)
		fmt.Println(_sharedLogConfig.LogGroups[KEY_GLOBAL_GROUP_LOG_CONFIG].Info.AppenderNoneConfig)
		fmt.Println(_sharedLogConfig.LogGroups[KEY_GLOBAL_GROUP_LOG_CONFIG].Debug)
	}
}
