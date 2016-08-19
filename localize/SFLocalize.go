//  Copyright 2016 slowfei And The Contributors All rights reserved.
//
//  Software Source Code License Agreement (BSD License)
//
//  Create on 2013-08-24
//  Update on 2016-08-19
//  Email  slowfei@nnyxing.com
//  Home   http://www.slowfei.com

// 本地化语言
package SFLocalize

import (
	"os"
	"path"
	"strings"
)

const (
	// default localize directory name
	DEFAULT_DIRNAME = "Localize"

	// default localize file name
	DEFAULT_KEYSTRINGS_NAME = "localize"

	DEFAULT_KEYSTRINGS_FILE_SUFFIX = "keystrings"

	// default language directory name
	DEFAULT_LANG = "default"
)

/**
 *	localize interface
 */
type ILocalize interface {
	/**
	 *	get localize keystrings key on string value
	 *	specified file localize-[LanguageCode].keystrings
	 *
	 *	@param `langCode` language code
	 *	@param `key` keystrings file key @param `comt` comment be empty @return `code` language code
	 *	@return `keyVal` key on value
	 *	@return `isExist` whether find localize info, false is not.
	 */
	KeyValue(langCode, key, comt string) (code, keyVal string, isExist bool)

	/**
	 *	by filepath get localize keystrings key on string value
	 *
	 *	@param `langCode` language code
	 *	@param `key` keystrings file key
	 *	@param `filename` keystrings file name,suffix is not required
	 *	@param `comt` comment be empty
	 *	@return `code` language code
	 *	@return `keyVal` key on value
	 *	@return `isExist` whether find localize info, false is not.
	 */
	KeyValueByFilename(langCode, key, filename, comt string) (code, keyVal string, isExist bool)

	/**
	 *	by language code get localize file info
	 *
	 *	@param `langCode` language code
	 *	@param `filepath` relative path
	 *	@return `fullPath` absolute path
	 *	@return `fi` file info
	 */
	FileInfo(langCode, filepath string) (fullPath string, fi os.FileInfo)
}

/***2-keystrings文件的使用说明
 *	keystrings文件是key on value键与值对应的本地化存储文件。
 *	文件内容：
 *	key=localize string
 *	key2=localize value
 *
 *	文件命名格式：
 *	localize.[zh-CN].keystrings (注：localize是默认使用的文件名)
 *	如需要自定义文件名也可以替换“localize”名称，然后使用LocalizeByFilename(...)函数方法。
 *
 *	关于语言代码可以参考语言代码列表
 */

// localize keys on values type
type KeyStrings map[string]string

/**
 *	language struct info, Code is only
 */
type Language struct {
	Code     string                // language code
	LangName string                // language name
	Area     string                // use area
	IsLang   bool                  // language tag
	KeyFiles map[string]KeyStrings // keystrings files
}

/**
 *	language struct by Code var short
 */
type LanguageToShort []Language

func (lts LanguageToShort) Len() int           { return len(lts) }
func (lts LanguageToShort) Less(i, j int) bool { return len(lts[i].Code) > len(lts[j].Code) }
func (lts LanguageToShort) Swap(i, j int)      { lts[i], lts[j] = lts[j], lts[i] }

/**
 *	ILocalize implement
 */
type localize struct {
	TagName   string
	RootPath  string
	Languages []Language
}

/**
 *	implement ILocalize
 */
func (l *localize) KeyValue(langCode, key, comt string) (code, keyVal string, isExist bool) {
	code, keyVal, isExist = l.KeyValueByFilename(langCode, key, DEFAULT_KEYSTRINGS_NAME, comt)
	return
}

/**
 *	implement ILocalize
 */
func (l *localize) KeyValueByFilename(langCode, key, filename, comt string) (code, keyVal string, isExist bool) {

	var lang Language

	lang, isExist = languageByCode(langCode, l.Languages)
	if isExist {
		var kfs KeyStrings
		if kfs, isExist = lang.KeyFiles[filename]; isExist {
			if keyVal, isExist = kfs[key]; isExist {
				code = lang.Code
			} else {
				keyVal = key
			}
		}
	}

	return
}

/**
 * implement ILocalize
 */
func (l *localize) FileInfo(langCode, filepath string) (fullPath string, fi os.FileInfo) {
	lang, isExist := languageByCode(langCode, l.Languages)

	if isExist {
		fullPath = path.Join(l.RootPath, lang.Code, filepath)
		fi, _ = os.Stat(fullPath)
	}

	return
}

/**
 *	find language by language code
 */
func languageByCode(langCode string, languages []Language) (Language, bool) {
	var lang Language
	isExist := false

	for i := 0; i < len(languages); i++ {
		tempLang := languages[i]

		// Language.Code 已经过排序处理，所以不用担心zh-Hant\zh-Hant-HK 先后顺序问题
		if 0 <= strings.Index(langCode, tempLang.Code) {
			isExist = true
			lang = tempLang
			break
		}
	}

	return lang, isExist
}
