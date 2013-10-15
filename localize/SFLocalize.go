//	本地化语言
//	TODO 还未实现
//
//	copyright 2013 slowfei
//	email		slowfei@foxmail.com
//	createTime 	2013-8-24
//	updateTime	2013-8-24
package SFLocalize

import (
	"errors"
)

const (
	//	语言标识
	LOCALIZE_CHINESE    = "zh-CN"
	LOCALIZE_CHINESE_TW = "zh-TW"
	LOCALIZE_CHINESE_HK = "zh-HK"
	LOCALIZE_ENGLISH    = "en-US"
)

var (
	//	程序唯一的本地化信息
	thisLocalize *localize
)

type localize struct {
	currentLang string
	langs       map[string]map[string]string
}

func (l *localize) string(key, comment string) string {
	return key
}

//	初始化本地化语言信息
func InitLocalize() {
	thisLocalize = &localize{currentLang: LOCALIZE_CHINESE}
}

//	切换语言
func ChangeLang(langTag string) error {
	//	由于没有具体实现到，所以展示直接设置 TODO
	thisLocalize.currentLang = langTag

	isSet := false
	for k, _ := range thisLocalize.langs {
		if langTag == k {
			isSet = true
			break
		}
	}

	if isSet {
		thisLocalize.currentLang = langTag
		return nil
	} else {
		return errors.New("can not find:" + langTag)
	}

}

//	获取本地化语言信息
//	@key		key查询不到相应的本地化信息将直接返回key
//	@comment	只为了进行注解，没有其他什么作用
func String(key, comment string) string {
	return thisLocalize.string(key, comment)
}

//	获取当前语言标识
func CurrentLang() string {
	return thisLocalize.currentLang
}
