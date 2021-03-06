//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-10-14
//  Update on 2013-10-17
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//
//	time 辅助工具
//
package SFTimeUtil

import (
	"strings"
	"time"
)

var (
	//yyyyMMddhhmmssSSS
	_YMDHMSSFormat = []string{
		"MST", "MST",
		"Z07:00", "Z07:00",

		"_yyyy", "-2006", //由于转换为 _2006 "_2"会被golang进行格式化日 02,2,_2
		"yyyy", "2006",
		"yy", "06",
		"MM", "01",
		"M", "1",
		"dd", "02",
		"d", "2",

		"hh", "15",
		"h", "3",
		"mm", "04",
		"m", "4",
		"ss", "05",
		"s", "5",

		"SSSSSSSSS", ".9999999999",
		"SSSSSSSS", ".99999999",
		"SSSSSSS", ".9999999",
		"SSSSSS", ".999999",
		"SSSSS", ".99999",
		"SSSS", ".9999",
		"SSS", ".999",
		"SS", ".99",
		"S", ".9",
	}

	//	符号标记的格式化
	_YMDHMSSSignFormat = []string{
		"${MST}", "MST",
		"${Z07:00}", "Z07:00",

		"_${yyyy}", "-2006",
		"${yyyy}", "2006",
		"${yy}", "06",
		"${MM}", "01",
		"${M}", "1",
		"${dd}", "02",
		"${d}", "2",

		"${hh}", "15",
		"${h}", "3",
		"${mm}", "04",
		"${m}", "4",
		"${ss}", "05",
		"${s}", "5",

		"${SSSSSSSSS}", ".9999999999",
		"${SSSSSSSS}", ".99999999",
		"${SSSSSSS}", ".9999999",
		"${SSSSSS}", ".999999",
		"${SSSSS}", ".99999",
		"${SSSS}", ".9999",
		"${SSS}", ".999",
		"${SS}", ".99",
		"${S}", ".9",
	}

	// yyyyMMddhhmmssSSS replacer
	_YMDHMSSReplacer     = strings.NewReplacer(_YMDHMSSFormat...)
	_YMDHMSSSignReplacer = strings.NewReplacer(_YMDHMSSSignFormat...)
)

//	yyyyMMddhhmmssSSS格式化时间返回字符串
//	这里不进行时区的转换了，根据glang提供的格式可以自行添加
//	"yyyy-MM-dd hh:mm:ssSSSSSSSSS -0700 MST" = "2006-01-02 15:04:05.999999999 -0700 MST"
//
//	@t
//	@format
//	@return format time string
func YMDHMSSFormat(t time.Time, format string) string {
	layout := _YMDHMSSReplacer.Replace(format)
	return t.Format(layout)
}

//	使用${}个标记格式
func YMDHMSSSignFormat(t time.Time, format string) string {
	layout := _YMDHMSSSignReplacer.Replace(format)
	return t.Format(layout)
}

//	yyyyMMddhhmmssSSS格式化解析时间对象
//	@timeStr
//	@format
func YMDHMSSParse(timeStr, format string) (time.Time, error) {
	layout := _YMDHMSSReplacer.Replace(format)
	return time.Parse(layout, timeStr)
}

//	使用${}个标记格式
func YMDHMSSSignParse(timeStr, format string) (time.Time, error) {
	layout := _YMDHMSSSignReplacer.Replace(format)
	return time.Parse(layout, timeStr)
}

//	yyyyMMddhhmmssSSS格式化解析时间对象
//	@loc *time.Location
//	@return
func YMDHMSSParseInLocation(timeStr, format string, loc *time.Location) (time.Time, error) {
	layout := _YMDHMSSReplacer.Replace(format)
	return time.ParseInLocation(layout, timeStr, loc)
}

//	使用${}个标记格式
func YMDHMSSSignParseInLocation(timeStr, format string, loc *time.Location) (time.Time, error) {
	layout := _YMDHMSSSignReplacer.Replace(format)
	return time.ParseInLocation(layout, timeStr, loc)
}
