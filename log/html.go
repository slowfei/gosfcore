//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-11-05
//  Update on 2013-11-06
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

// html handle
package SFLog

import (
	"fmt"
	"github.com/slowfei/gosfcore/utils/filemanager"
	"github.com/slowfei/gosfcore/utils/time"
)

const (
	//
	HTMLHandLayout = `
<!doctype html>
<html lang="en">
<head>
<meta charset="UTF-8"><title>%v</title>
<script type="text/javascript" src="http://ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script> 
<script> $(function() {$(".file_info").hide(); $(".stack").hide(); $(".show").click(function(event) {var $_this = $(this); var $logFrame = $_this.parent(); var $fileInfo = $logFrame.next(); var $stack = $logFrame.next().next(); $fileInfo.toggle(); $stack.toggle(); var char = $_this.html(); if (char == "▼"){$_this.html("▲"); }else{$_this.html("▼"); } }); $(".sort").click(function(event) {$(".sort").html(function(){return $(this).attr("t"); }); var s = $(this).attr("s"); var t = $(this).attr("t"); var index = $(this).index(); if (s == "1"){$(this).attr("s","2"); $(this).html(t+" ▼"); }else{$(this).attr("s","1"); $(this).html(t+" ▲"); } var tdLogInfos = $(".log_info").toArray().sort(function(a,b){if (t == "Time"){var date1 = new Date(Date.parse($("td",a).eq(1).text().replace(/-/g, "/"))); var date2 = new Date(Date.parse($("td",b).eq(1).text().replace(/-/g, "/"))); if (s == "1"){return date2 - date1; }else{return date1 - date2; } }else{var sa = $("td",a).eq(index).text(); var sb = $("td",b).eq(index).text(); if (s == "1"){return sa.localeCompare(sb); }else{return sb.localeCompare(sa); } } }); $.each(tdLogInfos,function(i,v){var fileInfo = $(v).next(); var stack = $(v).next().next(); $(v).appendTo('tbody'); $(fileInfo).appendTo('tbody'); $(stack).appendTo('tbody'); }); }); }); </script>
<style> table{background-color: #3a73ac;} thead th{background-color: #2d5a86; color: #fff; border-bottom: 1px solid #999;} table td{padding: 5px 5px;background-color: #fff; font-size: 14px;} pre{white-space: pre-wrap; /* css-3 */ white-space: -moz-pre-wrap; /* Mozilla, since 1999 */ white-space: -pre-wrap; /* Opera 4-6 */ white-space: -o-pre-wrap; /* Opera 7 */ word-wrap: break-word; /* Internet Explorer 5.5+ */ font-size: 14px; margin: 3px 3px;} .log_info{} .log_info td{border-top: 1px solid #999; } .file_info{} .stack{color: red} span{font-weight: bold; color: #fff; padding: 3px 3px; } span.info{background-color: #000; } span.debug{background-color:#080; } span.error{background-color:red; } span.warn{background-color:#dd0033; } span.fatal{background-color:#bb0033; } span.panic{background-color: #760c0c } </style> 
</head>
<body>
<div style="text-align:center;"><h1>%v</h1></div><table style="width: 100%;"><thead><tr><th width="10">&nbsp;</th><th width="160" class="sort" s="1" t="Time">Time</th> <th width="70" class="sort" s="1" t="Target">Target</th><th width="150" class="sort" s="1" t="Log Group">Log Group</th><th width="150" class="sort" s="1" t="Log Tag">Log Tag</th><th>Message</th></tr></thead><tbody>
	`
	//
	HTMLEndLayout = `</tbody></table></body></html> `

	//
	HTMLContentLayout = `<tr class="log_info"><td align=center class="show">▲</td>
<td align=center>%v</td>
<td align=center><span class="%v">%v</span></td>
<td>%v</td>
<td>%v</td>
<td><pre>%v</pre></td></tr>
<tr class="file_info"><td colspan=6>%v</td></tr>
<tr class="stack"><td colspan=6><b>Stack</b><pre>%v</pre></td></tr>`

	DEFAULT_HTML_MAX_SIZE = 3 << 20 // default 5M html max size
	DEFAULT_HTML_MAX_NUM  = 1000    // default same name log html max num
)

//	appender html
type AppenderHtmlConfig struct {
	SavePath   string `json:"HtmlSavePath"`       // 文件存储路径, 			默认执行文件目录
	Name       string `json:"HtmlName"`           // 文件名(可以输入时间格式)  默认"(ExceFileName)-${yyyy}-${MM}-${dd}.html"
	Title      string `json:"HtmlTitle"`          // html title
	TimeFormat string `json:"HtmlTimeFormat"`     // time format			"yyyy-MM-dd hh:mm:ssSSSSSSSSS"
	MaxSize    int64  `json:"HtmlMaxSize"`        // 文件大小 byte			默认3M
	SameMaxNum int    `json:"HtmlSameNameMaxNum"` // 日志相同名称的最大数量，例如log(1).html...log(1000).html。默认1000，超出建立的数量将不会创建日志文件
}

// Appender impl console write
type AppenderHtml struct {
	writeFile *AppenderFile
}

func NewAppenderHtml() *AppenderHtml {
	ah := &AppenderHtml{}
	ah.writeFile = NewAppenderFile()
	ah.writeFile.defaultFileName = SFFileManager.GetExecFileName() + "-${yyyy}${MM}${dd}.html"
	ah.writeFile.defaultMaxSize = DEFAULT_HTML_MAX_SIZE
	ah.writeFile.defaultSameMaxNum = DEFAULT_HTML_MAX_NUM
	return ah
}

//	html write
func (ah *AppenderHtml) htmlWrite(msg *LogMsg, config *AppenderHtmlConfig) {

	savePath := config.SavePath
	fileName := config.Name
	maxSize := config.MaxSize
	sameMaxNum := config.SameMaxNum

	timeFormat := config.TimeFormat
	if 0 == len(timeFormat) {
		timeFormat = "yyyy-MM-dd hh:mm:ss"
	}
	timeString := SFTimeUtil.YMDHMSSFormat(msg.dateTime, timeFormat)
	target := string(msg.target)
	htmlFI := fmt.Sprintf("L%d: %s(...) %s", msg.fileLine, msg.funcName, msg.filePath)

	ah.writeFile.rwmutex.Lock()
	defer ah.writeFile.rwmutex.Unlock()
	file := ah.writeFile.getFile(savePath, fileName, msg.dateTime, maxSize, sameMaxNum)

	if nil != file {
		fileInfo, _ := file.Stat()
		if nil == fileInfo {
			return
		}
		if 0 == fileInfo.Size() {
			htmlHand := fmt.Sprintf(HTMLHandLayout, config.Title, config.Title)
			fmt.Fprintln(file, htmlHand)
		}
		htmlContent := fmt.Sprintf(HTMLContentLayout, timeString, target, target, msg.logGroup, msg.logTag, msg.msg, htmlFI, msg.stack)
		fmt.Fprintln(file, htmlContent)

		//	TODO 正在考虑如何添加结尾标签
	}
}

//	关闭所有日志文件
func (af *AppenderHtml) CloseAllLogHtml() {
	af.writeFile.CloseAllLogFile()
}

//	#interface impl
func (ah *AppenderHtml) Write(msg *LogMsg, configInfo interface{}) {
	if htmlConfig, ok := configInfo.(*AppenderHtmlConfig); ok {
		ah.htmlWrite(msg, htmlConfig)
	}
}

//	name = html
func (ah *AppenderHtml) Name() string {
	return VAL_APPENDER_HTML
}
