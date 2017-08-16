//  Copyright 2016 slowfei And The Contributors All rights reserved.
//
//  Software Source Code License Agreement (BSD License)
//
//  Create on 2016-07-18
//  Update on 2016-08-19
//  Email  slowfei@nnyxing.com
//  Home   http://www.slowfei.com

// initiation and load languages
package SFLocalize

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
)

var (
	_languages []Language = make([]Language, 0, 0)
)

/**
 *	localize directory path
 */
type LocDir string

/**
 *	load language directory info
 *
 *	@param `tagName` localize tag
 *	@param `locdir` localize directory path
 *	@return `ILocalize`
 *	@return `error`
 */
func LoadLanguages(tagName string, locdir LocDir) (ILocalize, error) {
	var dirPath string = string(locdir)
	// check directory
	fidir, err := os.Stat(dirPath)
	if nil != err {
		return nil, err
	}

	if !fidir.IsDir() {
		return nil, errors.New("It is not a valid directory path.")
	}

	// init localize struct
	local := new(localize)
	local.RootPath = dirPath
	local.TagName = tagName

	// init keystrings pathsfiles
	keystringspathfiles := parseKeyStringsFilePaths(dirPath)

	// languages source data
	langData := SourceData()
	dataLine := strings.Split(langData, "\n")

	for i := 0; i < len(dataLine); i++ {
		langInfo := strings.Split(dataLine[i], ";")

		langCode := ""
		langName := ""
		langArea := ""
		isLang := false

		// 3 is language, 4 is language and area
		if 3 == len(langInfo) {
			langName = langInfo[0]
			langCode = langInfo[2]
			isLang = true
		} else if 4 == len(langInfo) {
			langName = langInfo[0]
			langArea = langInfo[1]
			langCode = langInfo[3]
		}

		if 0 != len(langCode) {

			langStruct := Language{}
			langStruct.KeyFiles = make(map[string]KeyStrings)
			isDir := false
			isKSFiles := parseLangRootKeyStringsFiles(langCode, keystringspathfiles, langStruct.KeyFiles)

			joinPath := path.Join(dirPath, langCode)
			fi, err := os.Stat(joinPath)

			if nil == err && fi.IsDir() {
				isDir = true
				parseLangDirKeyStrings(joinPath, langStruct.KeyFiles)

			}

			if isDir || isKSFiles {
				langStruct.Code = langCode
				langStruct.IsLang = isLang
				langStruct.LangName = langName
				langStruct.Area = langArea
				local.Languages = append(local.Languages, langStruct)
				sort.Sort(LanguageToShort(local.Languages))
			}

		}
	}

	return local, nil
}

/**
 *	parse "*.keystrings" files to paths
 */
func parseKeyStringsFilePaths(dirPath string) []string {
	var filepaths []string

	suffix := "." + DEFAULT_KEYSTRINGS_FILE_SUFFIX
	files, _ := ioutil.ReadDir(dirPath)

	for _, file := range files {
		fileName := file.Name()
		if !file.IsDir() && strings.HasSuffix(fileName, suffix) {
			filepaths = append(filepaths, path.Join(dirPath, fileName))
		}
	}

	return filepaths
}

/**
 *	parse language code directory keystrings file
 */
func parseLangDirKeyStrings(dirpath string, mapkey map[string]KeyStrings) {

	suffix := "." + DEFAULT_KEYSTRINGS_FILE_SUFFIX
	suffixLen := len(suffix)
	files, _ := ioutil.ReadDir(dirpath)

	for _, file := range files {
		fileName := file.Name()
		fileNameLen := len(fileName)
		if !file.IsDir() && strings.HasSuffix(fileName, suffix) {
			keyFilename := fileName[:fileNameLen-suffixLen]

			ksv, ok := mapkey[keyFilename]
			if !ok {
				ksv = make(KeyStrings)
			}

			parseKeyStringsFile(path.Join(dirpath, fileName), ksv)

			mapkey[keyFilename] = ksv

		}

	}
}

/**
 *	parse language root directory keystrings file
 *
 *	filepaths is root directory the "*.keystrings" file
 *
 *	return bool is append KeyStrings the count
 */
func parseLangRootKeyStringsFiles(langCode string, filepaths []string, mapkey map[string]KeyStrings) bool {

	// suffix virtual value is "xxxx.zh-CN.keystrings"
	suffix := "." + langCode + "." + DEFAULT_KEYSTRINGS_FILE_SUFFIX
	suffixLen := len(suffix)
	isAdd := false

	for _, filepath := range filepaths {
		fileName := path.Base(filepath)
		fileNameLen := len(fileName)
		if strings.HasSuffix(fileName, suffix) {
			keyFilename := fileName[:fileNameLen-suffixLen]

			ksv := make(KeyStrings)
			parseKeyStringsFile(filepath, ksv)

			mapkey[keyFilename] = ksv
			isAdd = true
		}
	}
	return isAdd
}

/**
 *	parse .keystrings file key and value
 */
func parseKeyStringsFile(filepath string, ksv KeyStrings) {

	file, err := os.Open(filepath)
	if nil == err {
		defer func() {
			if err := file.Close(); err != nil {

			}
		}()

		br := bufio.NewReader(file)
		for {

			lineBytes, isPrefix, err := br.ReadLine()
			if nil != err || io.EOF == err || isPrefix {
				break
			}

			signIndex := bytes.IndexByte(lineBytes, '=')
			if 0 < signIndex {
				key := string(lineBytes[:signIndex])
				value := string(lineBytes[signIndex+1:])
				ksv[key] = value
			}

		}
	}
}

/***语言代码列表参考

Afrikaans;布尔语[南非荷兰语];af
Afrikaans;Namibia;纳米比亚;af-NA
Afrikaans;South Africa;南非;af-ZA

Aghem;;agq
Aghem;Cameroon;喀麦隆;agq-CM

Akan;库阿语;ak
Akan;Ghana;加纳;ak-GH

Albanian;阿尔巴尼亚语;sq
Albanian;Albania;阿尔巴尼亚;sq-AL
Albanian;Kosovo;科索沃;sq-XK
Albanian;Macedonia;马其顿;sq-MK

Amharic;阿姆哈拉语;am
Amharic;Ethiopia;埃塞俄比亚;am-ET

Arabic;阿拉伯语;ar
Arabic;Algeria;阿尔及利亚;ar-DZ
Arabic;Bahrain;巴林;ar-BH
Arabic;Chad;乍得;ar-TD
Arabic;Comoros;科摩罗;ar-KM
Arabic;Djibouti;吉布提;ar-DJ
Arabic;Egypt;埃及;ar-EG
Arabic;Eritrea;厄立特里亚;ar-ER
Arabic;Iraq;伊拉克;ar-IQ
Arabic;Israel;以色列;ar-IL
Arabic;Jordan;约旦;ar-JO
Arabic;Kuwait;科威特;ar-KW
Arabic;Lebanon;黎巴嫩;ar-LB
Arabic;Libya;利比亚;ar-LY
Arabic;Mauritania;毛里塔尼亚;ar-MR
Arabic;Morocco;摩洛哥;ar-MA
Arabic;Oman;阿曼;ar-OM
Arabic;Palestine;巴勒斯坦;ar-PS
Arabic;Qatar;卡塔尔;ar-QA
Arabic;Saudi Arabia;沙特阿拉伯;ar-SA
Arabic;Somalia;索马里;ar-SO
Arabic;South Sudan;南苏丹;ar-SS
Arabic;Sudan;苏丹;ar-SD
Arabic;Syria;叙利亚;ar-SY
Arabic;Tunisia;突尼斯;ar-TN
Arabic;United Arab Emirates;阿拉伯联合酋长国;ar-AE
Arabic;Western Sahara;西撒哈拉;ar-EH
Arabic;World;世界;ar-001
Arabic;Yemen;也门;ar-YE

Armenian;亚美尼亚语;hy
Armenian;Armenia;亚美尼亚;hy-AM

Assamese;阿萨姆语;as
Assamese;India;印度;as-IN

Asu; ;asa
Asu;Tanzania;坦桑尼亚;asa-TZ

Azerbaijani;阿塞拜疆语;az
Azerbaijani;Azerbaijan;阿塞拜疆;az-AZ
Azerbaijani;Cyrillic;西里尔;az-Cyrl
Azerbaijani;Cyrillic, Azerbaijan;西里尔，阿塞拜疆;az-CyrI-AZ

Bafia;巴菲亚语;ksf
Bafia;Cameroon;喀麦隆;ksf-CM

Bambara;班巴拉语;bm
Bambara;Mali;马里;bm-ML

Basaa;巴萨语;bas
Basaa;Cameroon;喀麦隆;bas-CM

Basque;巴斯克语;eu
Basque;Spain;西班牙;eu-ES

Belarusian;白俄罗斯语;be
Belarusian;Belarus;白俄罗斯;be-BY

Bemba;本巴语;bem
Bemba;Zambia;赞比亚;bem-ZM

Bena;比纳语;bez
Bena;Tanzania;坦桑尼亚;bez-TZ

Bengali;孟加拉语;bn
Bengali;Bangladesh;孟加拉;bn-BD
Bengali;India;印度;bn-IN

Bodo;博多语;brx
Bodo;India;印度;brx-IN

Bosnian;波斯尼亚语;bs
Bosnian;Bosnia & Herzegovina;波斯尼亚与黑塞哥维那;bs-BA
Bosnian;Cyrillic;西里尔;bs-Cyrl
Bosnian;Cyrillic, Bosnia & Herzegovina;西里尔，波斯尼亚与黑塞哥维那;bs-Cyrl-BA

Breton;布列塔尼语;br
Breton;France;法国;br-FR

Bulgarian;保加利亚语;bg
Bulgarian;Bulgaria;保加利亚;bg-BG

Burmese;缅甸语;my
Burmese;Myanmar;缅甸;my-MM

Catalan;加泰罗尼亚语;ca
Catalan;Andorra;安道尔;ca-AD
Catalan;France;法国;ca-FR
Catalan;Italy;意大利;ca-IT
Catalan;Spain;西班牙;ca-ES

Cebuano;宿务语;ceb

Central Atlas Tamazight;中央阿特拉斯塔马塞特语;tzm
Central Atlas Tamazight;Morocco;摩洛哥;tzm-MA

Central Kurdish;库尔德语;ckb
Central Kurdish;Iran;伊朗;ckb-IR
Central Kurdish;Iraq;伊拉克;ckb-IQ

Cherokee;切罗基语;chr
Cherokee;United States;美国;chr-US

Chichewa;齐切瓦语;ny

Chiga; ;cgg
Chiga;Uganda;乌干达;cgg-UG

Chinese Simplified;简体中文;zh-Hans
Chinese Traditional;繁体中文;zh-Hant
Chinese;Simplified, China;简体，中国;zh-CN
Chinese;Simplified, Hong Kong;简体，香港;zh-Hans-HK
Chinese;Simplified, Macau;简体，澳门;zh-Hans-MO
Chinese;Simplified, Singapore;简体，新加坡;zh-Hans-SG
Chinese;Simplified, Taiwan;简体，台湾;zh-Hans-TW
Chinese;Traditional, Hong Kong;繁体，香港;zh-Hant-HK
Chinese;Traditional, Macau;繁体，澳门;zh-Hant-MO
Chinese;Traditional, Singapore;繁体，新加坡;zh-Hant-SG
Chinese;Traditional, Taiwan;繁体，台湾;zh-Hant-TW

Colognian; ;ksh
Colognian;Germany;德国;ksh-DE

Cornish;康沃尔语;kw
Cornish;United Kingdom;英国;kw-GB

Corsican;科西嘉语;co

Croatian;克罗地亚语;hr
Croatian;Bosnia & Herzegovina;波斯尼亚和黑塞哥维那;hr-BA
Croatian;Croatia;克罗地亚;hr-HR

Czech;捷克语;cs
Czech;Czech Republic;捷克共和国;cs-CZ

Danish;丹麦语;da
Danish;Denmark;丹麦;da-DK
Danish;Greenland;格陵兰;da-GL

Duala;杜阿拉语;dua
Duala;Cameroon;喀麦隆;dua-CM

Dutch;荷兰语;nl
Dutch;Aruba;阿鲁巴;nl-AW
Dutch;Belgium;比利时;nl-BE
Dutch;Caribbean Netherlands;荷兰加勒比区;nl-BQ
Dutch;Curacao;库拉索;nl-CW
Dutch;Netherlands;荷兰;nl-NL
Dutch;Sint Maarten;圣马丁岛;nl-SX
Dutch;Suriname;苏里南;nl-SR

Dzongkha;宗喀语;dz
Dzongkha;Bhutan;不丹;dz-BT

Embu;恩布语;ebu
Embu;Kenya;肯尼亚;ebu-KE

English;英语;en
English;Albania;阿尔巴尼亚;en-AL
English;American Samoa;美属萨摩亚;en-AS
English;Andorra;安道尔共和国;en-AD
English;Anguilla;安圭拉岛;en-Al
English;Antigua & Barbuda;安提瓜和巴布达岛;en-AG
English;Australia;澳大利亚;en-AU
English;Austria;奥地利;en-AT
English;Bahamas;巴哈马群岛;en-BS
English;Barbados;巴巴多斯岛;en-BB
English;Belgium;比利时;en-BE
English;Belize;伯利兹;en-BZ
English;Bermuda;百慕大群岛;en-BM
English;Bosnia & Herzegovina;波斯尼亚和黑塞哥维纳;en-BA
English;Botswana;博茨瓦纳;en-BW
English;British Indian Ocean Territory;英属印度洋领地;en-IO
English;British Virgin Islands;英属维尔京群岛;en-VG
English;Cameroon;喀麦隆;en-CM
English;Canada;加拿大;en-CA
English;Cayman Islands;开曼群岛;en-KY
English;Christmas Island;圣诞岛;en-CX
English;Cocos [Keeling] Islands;科科斯基林群岛;en-CC
English;Cook Islands;库克群岛;en-CK
English;Croatia;克罗地亚;en-HR
English;Cyprus;塞浦路斯;en-CY
English;Czech Republic;捷克共和国;en-CZ
English;Denmark;丹麦;en-DK
English;Diego Garcia;迪戈加西亚;en-DG
English;Dominica;多米尼加;en-DM
English;Eritrea;厄立特里亚;en-ER
English;Estonia;爱沙尼亚;en-EE
English;Europe;欧洲;en-150
English;Falkland Islands;福克兰群岛;en-FK
English;Fiji;斐济;en-FJ
English;Finland;芬兰;en-Fl
English;France;法国;en-FR
English;Gambia;冈比亚;en-GM
English;Germany;德国;en-DE
English;Ghana;加纳;en-GH
English;Gibraltar;直布罗陀;en-GI
English;Greece;希腊;en-GR
English;Grenada;格林纳达;en-GD
English;Guam;关岛;en-GU
English;Guernsey;根西岛;en-GG
English;Guyana;圭亚那;en-GY
English;Hong Kong;香港;en-HK
English;Hungary;匈牙利;en-HU
English;Iceland;冰岛;en-IS
English;India;印度;en-IN
English;Ireland;爱尔兰;en-IE
English;Isle of Man;马恩岛;en-IM
English;Israel;以色列;en-IL
English;Italy;意大利;en-IT
English;Jamaica;牙买加;en-JM
English;Jersey;新泽西;en-JE
English;Kenya;肯尼亚;en-KE
English;Kiribati;基里巴斯;en-KI
English;Latvia;拉脱维亚;en-LV
English;Lesotho;莱索托;en-LS
English;Liberia;利比里亚;en-LR
English;Lithuania;立陶宛;en-LT
English;Luxembourg;卢森堡;en-LU
English;Macau;澳门;en-MO
English;Madagascar;马达加斯加岛;en-MG
English;Malawi;马拉维;en-MW
English;Malaysia;马来西亚;en-MY
English;Malta;马耳他;en-MT
English;Marshall Islands;马绍尔群岛;en-MH
English;Mauritius;毛里求斯;en-MU
English;Micronesia;密克罗尼西亚;en-FM
English;Montenegro;黑山共和国;en-ME
English;Montserrat;蒙特塞拉特岛;en-MS
English;Namibia;纳米比亚;en-NA
English;Nauru;瑙鲁岛;en-NR
English;Netherlands;荷兰;en-NL
English;New Zealand;新西兰;en-NZ
English;Nigeria;尼日利亚;en-NG
English;Niue;纽埃岛;en-NU
English;Norfolk Island;诺福克岛;en-NF
English;Northern Mariana Islands;北马里亚纳群岛;en-MP
English;Norway;挪威;en-No
English;Pakistan;巴基斯坦;en-PK
English;Palau;帕劳群岛;en-PW
English;Papua New Guinea;巴布亚新几内亚;en-PG
English;Philippines;菲律宾;en-PH
English;Pitcairn Islands;皮特凯恩群岛;en-PN
English;Poland;波兰;en-PL
English;Portugal;葡萄牙;en-PT
English;Puerto Rico;波多黎各;en-PR
English;Romania;罗马尼亚;en-Ro
English;Russia;俄国;en-RU
English;Rwanda;卢旺达;en-RW
English;Samoa;萨摩亚群岛;en-WS
English;Seychelles;塞舌尔;en-SC
English;Sierra Leone;塞拉利昂;en-SL
English;Singapore;新加坡;en-SG
English;Sint Maarten;圣马丁岛;en-SX
English;Slovakia;斯洛伐克;en-SK
English;Slovenia;斯洛文尼亚;en-SI
English;Solomon Islands;所罗门群岛;en-SB
English;South Africa;南非;en-ZA
English;South Sudan;南苏丹;en-SS
English;Spain;西班牙;en-ES
English;St. Helena;圣赫勒拿岛;en-SH
English;St. Kitts & Nevis;圣基茨和尼维斯;en-KN
English;St. Lucia;圣卢西亚;en-LC
English;St. Vincent & Grenadines;圣文森特和格林纳丁斯;en-VC
English;Sudan;苏丹;en-SD
English;Swaziland;斯威士兰;en-SZ
English;Sweden;瑞典;en-SE
English;Switzerland;瑞士;en-CH
English;Tanzania;坦桑尼亚;en-TZ
English;Tokelau;托克劳;en-TK
English;Tonga;汤加;en-TO
English;Trinidad & Tobago;特立尼达和多巴哥;en-TT
English;Turkey;土耳其;en-TR
English;Turks & Caicos Islands;特克斯和凯科斯群岛;en-TC
English;Tuvalu;图瓦卢;en-TV
English;US Outlying Islands;美国离岛;en-UM
English;US Virgin Islands;美国维尔京群岛;en-VI
English;Uganda;乌干达;en-UG
English;United Kingdom;英国;en-GB
English;United States;美国;en-US
English;United States, Computer;美国，计算机;en-US-POSIX
English;Vanuatu;瓦努阿图;en-VU
English;World;世界;en-001
English;Zambia;赞比亚;en-ZM
English;Zimbabwe;津巴布韦;en-ZW

Esperanto;世界语;eo

Estonian;爱沙尼亚语;et
Estonian;Estonia;爱沙尼亚;et-EE

Ewe;埃维语;ee
Ewe;Ghana;加纳;ee-GH
Ewe;Togo;多哥;ee-TG

Ewondo;埃翁多语;ewo
Ewondo;Cameroon;喀麦隆;ewo-CM

Faroese;法罗人[语];fo
Faroese;Faroe Islands;法罗群岛;fo-FO

Filipino;菲律宾语;fil
Filipino;Philippines;菲律宾;fil-PH

Finnish;芬兰语;fi
Finnish;Finland;芬兰;fi-Fl

French;法语;fr
French;Algeria;阿尔及利亚;fr-DZ
French;Belgium;比利时;fr-BE
French;Benin;贝宁湾;fr-BJ
French;Burkina Faso;布基纳法索;fr-BF
French;Burundi;布隆迪;fr-Bl
French;Cameroon;喀麦隆;fr-CM
French;Canada;加拿大;fr-CA
French;Central African Republic;中非共和国;fr-CF
French;Chad;乍得;fr-TD
French;Comoros;科摩罗;fr-KM
French;Congo - Brazzaville;刚果-布拉柴维尔;fr-CG
French;Congo - Kinshasa;刚果-金沙萨;fr-CD
French;Cote d’lvoire;科特迪瓦;fr-Cl
French;Djibouti;吉布提;fr-DJ
French;Equatorial Guinea;赤道几内亚;fr-GQ
French;France;法国;fr-FR
French;French Guiana;法属圭亚那;fr-GF
French;French Polynesia;法属波利尼西亚;fr-PF
French;Gabon;加蓬;fr-GA
French;Guadeloupe;瓜德罗普岛;fr-GP
French;Guinea;几内亚;fr-GN
French;Haiti;海地;fr-HT
French;Luxembourg;卢森堡;fr-LU
French;Madagascar;马达加斯加;fr-MG
French;Mali;马里;fr-ML
French;Martinique;马提尼克岛;fr-MQ
French;Mauritania;毛里塔尼亚;fr-MR
French;Mauritius;毛里求斯;fr-MU
French;Mayotte;马约特岛;fr-YT
French;Monaco;摩纳哥;fr-MC
French;Morocco;摩洛哥;fr-MA
French;New Caledonia;新喀里多尼亚;fr-NC
French;Niger;尼日尔;fr-NE
French;Reunion;留尼旺岛;fr-RE
French;Rwanda;卢旺达;fr-RW
French;Senegal;塞内加尔;fr-SN
French;Seychelles;塞舌尔;fr-SC
French;St. Barthelemy;圣巴托洛缪岛;fr-BL
French;St. Martin;圣马丁;fr-MF
French;St. Pierre & Miquelon;圣石和密克隆岛;fr-PM
French;Switzerland;瑞士;fr-CH
French;Syria;叙利亚共和国;fr-SY
French;Togo;多哥;fr-TG
French;Tunisia;突尼斯;fr-TN
French;Vanuatu;瓦努阿图;fr-VU
French;Wallis & Futuna;瓦利斯和富图纳;fr-WF

Friulian;弗留利语;fur
Friulian;Italy;意大利;fur-IT

Fulah;富拉赫语;ff
Fulah;Cameroon;喀麦隆;ff-CM
Fulah;Guinea;几内亚;ff-GN
Fulah;Mauritania;毛利塔尼亚;ff-MR
Fulah;Senegal;塞内加尔;ff-SN

Galician;加利西亚语;gl
Galician;Spain;西班牙;gl-ES

Ganda;干达人[语];lg
Ganda;Uganda;乌干达;lg-UG

Georgian;格鲁吉亚语;ka
Georgian;Georgia;格鲁吉亚;ka-GE

German;德语;de
German;Austria;奥地利;de-AT
German;Belgium;比利时;de-BE
German;Germany;德国;de-DE
German;Liechtenstein;列支敦士登;de-LI
German;Luxembourg;卢森堡公国;de-LU
German;Switzerland;瑞士;de-CH

Greek;希腊语;el
Greek;Cyprus;塞浦路斯;el-CY
Greek;Greece;希腊;el-GR

Gujarati;古吉拉特语;gu
Gujarati;India;印度;gu-IN

Gusii;古西语;guz
Gusii;Kenya;肯尼亚;guz-KE

Haitian Creole;海地克里奥尔语;ht

Hausa;豪萨语;ha
Hausa;Ghana;加纳;ha-GH
Hausa;Niger;尼日尔;ha-NE
Hausa;Nigeria;尼日利亚;ha-NG

Hawaiian;夏威夷语;haw
Hawaiian;United States;美国;haw-US

Hebrew;希伯来语;he
Hebrew;Israel;以色列;he-IL

Hindi;印地语;hi
Hindi;India;印度;hi-IN

Hmong;苗语;hmong

Hungarian;匈牙利语;hu
Hungarian;Hungary;匈牙利;hu-HU

Icelandic;冰岛语;is
Icelandic;Iceland;冰岛;is-IS

Igbo;伊博语;ig
Igbo;Nigeria;尼日利亚;ig-NG

Inari Sami;伊纳里萨米语;smn
Inari Sami;Finland;芬兰;smn-FI

Indonesian;印尼语;id
Indonesian;Indonesia;印尼;id-ID

Inuktitut;因纽特语;iu
Inuktitut;Unified Canadian Aboriginal Syllabics;统一加拿大土著语音节;iu-Cans
Inuktitut;Unified Canadian Aboriginal Syllabics, Canada;统一加拿大土著语音节，加拿大;iu-Cans-CA

Irish;爱尔兰语;ga
Irish;Ireland;爱尔兰;ga-IE

Italian;意大利语;it
Italian;Italy;意大利;it-IT
Italian;San Marino;圣马力诺;it-SM
Italian;Switzerland;瑞士;it-CH

Japanese;日语;ja
Japanese;Japan;日本;ja-JP

Javanese;印尼爪哇语;jv

Jola-Fonyi; ;dyo
Jola-Fonyi;Senegal;塞内加尔;dyo-SN

Kabuverdianu; ;kea
Kabuverdianu;Cape Verde;佛得角;kea-CV

Kabyle;卡比尔语;kab
Kabyle;Algeria;阿尔及利亚;kab-DZ

Kako; ;kkj
Kako;Cameroon;喀麦隆;kkj-CM

Kalaallisut;格陵兰语;kl
Kalaallisut;Greenland;格陵兰;kl-GL

Kalenjin;卡伦津语;kln
Kalenjin;Kenya;肯尼亚;kln-KE

Kamba;坎巴语;kam
Kamba;Kenya;肯尼亚;kam-KE

Kannada;卡纳达语;kn
Kannada;India;印度;kn-IN

Kashmiri;克什米尔人[语];ks
Kashmiri;Arabic;阿拉伯;ks-Arab
Kashmiri;Arabic, India;阿拉伯，印度;ks-Arab-IN

Kazakh;哈萨克语;kk
Kazakh;Kazakhstan;哈萨克斯坦;kk-KZ

Khmer;高棉语;km
Khmer;Cambodia;柬埔寨;km-KH

Kikuyu;基库尤语;ki
Kikuyu;Kenya;肯尼亚;ki-KE

Kinyarwanda;卢旺达语;rw
Kinyarwanda;Rwanda;卢旺达;rw-RW

Konkani;孔卡尼语;kok
Konkani;India;印度;kok-IN

Korean; ;ko
Korean;North Korea;朝鲜;ko-KP
Korean;South Korea;韩国;ko-KR

Koyra Chiini; ;khq
Koyra Chiini;Mali;马里;khq-ML

Koyraboro Senni; ;ses
Koyraboro Senni;Mali;马里;ses-ML

Kwasio; ;nmg
Kwasio;Cameroon;喀麦隆;nmg-CM

Kyrgyz;吉尔吉斯语;ky
Kyrgyz;Kyrgyzstan;吉尔吉斯斯坦;ky-KG

Lakota;拉科塔语;lkt
Lakota;United States;美国;lkt-US

Langi; ;lag
Langi;Tanzania;坦桑尼亚;lag-TZ

Lao;老挝语;lo
Lao;Laos;老挝;lo-LA

Latin;拉丁语;la

Latvian;拉脱维亚语;lv
Latvian;Latvia;拉脱维亚;lv-LV

Lingala;林加拉语;ln
Lingala;Angola;安哥拉;ln-Ao
Lingala;Central African Republic;中非共和国;ln-CF
Lingala;Congo - Brazzaville;刚果-布拉柴维尔;ln-CG
Lingala;Congo - Kinshasa;刚果-金沙萨;ln-CD

Lithuanian;立陶宛语;lt
Lithuanian;Lithuania;立陶宛;lt-LT

Lower Sorbian;下索布语;dsb
Lower Sorbian;Germany;德国;dsb-DE

Luba-Katanga;卢巴加丹加语;lu
Luba-Katanga;Congo - Kinshasa;刚果-金沙萨;lu-CD

Luo; ;luo
Luo;Kenya;肯尼亚;luo-KE

Luxembourgish;卢森堡语;lb
Luxembourgish;Luxembourg;卢森堡;lb-LU

Luyia; ;luy
Luyia;Kenya;肯尼亚;luy-KE

Macedonian;马其顿语;mk
Macedonian;Macedonia;马其顿;mk-MK

Machame; ;jmc
Machame;Tanzania;坦桑尼亚;jmc-TZ

Makhuwa-Meetto;mgh
Makhuwa-Meetto;Mozambique;莫桑比克;mgh-MZ

Makonde; ;kde
Makonde;Tanzania;坦桑尼亚;kde-TZ

Maori;毛利语;mi

Malagasy;马尔加什语;mg
Malagasy;Madagascar;马达加斯加岛;mg-MG

Malay;马来语;ms
Malay;Arabic;阿拉伯;ms-Arab
Malay;Arabic, Brunei;阿拉伯，文莱;ms-Arab-BN
Malay;Arabic, Malaysia;阿拉伯，马来西亚;ms-Arab-MY
Malay;Brunei;文莱;ms-BN
Malay;Malaysia;马来西亚;ms-MY
Malay;Singapore;新加坡;ms-SG

Malayalam;马拉雅拉姆语;ml
Malayalam;India;印度;ml-IN

Maltese;马耳他语;mt
Maltese;Malta;马耳他;mt-MT

Manx;马恩岛语;gv
Manx;Isle of Man;马恩岛;gv-IM

Marathi;马拉地语;mr
Marathi;India;印度;mr-IN

Masai; ;mas
Masai;Kenya;肯尼亚;mas-KE
Masai;Tanzania;坦桑尼亚;mas-TZ

Meru; ;mer
Meru;Kenya;肯尼亚;mer-KE

Meta'; ;mgo
Meta';Cameroon;喀麦隆;mgo-CM

Mongolian;蒙古语;mn
Mongolian;Mongolia;蒙古;mn-MN

Morisyen; ;mfe
Morisyen;Mauritius;毛里求斯;mfe-MU

Mundang; ;mua
Mundang;Cameroon;喀麦隆;mua-CM

Nama; ;naq
Nama;Namibia;纳米比亚;naq-NA

Nepali;尼泊尔语;ne
Nepali;India;印度;ne-IN
Nepali;Nepal;尼泊尔;ne-NP

Ngiemboon; ;nnh
Ngiemboon;Cameroon;喀麦隆;nnh-CM

Ngomba; ;jgo
Ngomba;Cameroon;喀麦隆;jgo-CM

North Ndebele; ;nd
North Ndebele;Zimbabwe;津巴布韦;nd-ZW

Northern Sami; ;se
Northern Sami;Finland;芬兰;se-FI
Northern Sami;Norway;挪威;se-NO
Northern Sami;Sweden;瑞典;se-SE

Norwegian Bokmal;挪威语;nb
Norwegian Bokmal;Norway;挪威;nb-NO
Norwegian Bokmal;Svalbard & Jan Mayen;斯瓦尔巴特;nb-SJ

Norwegian Nynorsk;新挪威语;nn
Norwegian Nynorsk;Norway;挪威;nn-NO

Nuer;努尔语;nus
Nuer;Sudan;苏丹;nus-SD

Nyankole;言克尔语;nyn
Nyankole;Uganda;乌干达;nyn-UG

Oriya;奥里雅语;or
Oriya;India;印度;or-IN

Oromo;奥罗莫语;om
Oromo;Ethiopia;埃塞俄比亚;om-ET
Oromo;Kenya;肯尼亚;om-KE

Ossetic;奥塞特语;os
Ossetic;Georgia;格鲁吉亚;os-GE
Ossetic;Russia;俄罗斯;os-RU

Pashto;普什图语;ps
Pashto;Afghanistan;阿富汗;ps-AF

Persian;波斯语;fa
Persian;Afghanistan;阿富汗;fa-AF
Persian;Iran;伊朗;fa-IR

Polish;波兰语;pl
Polish;Poland;波兰;pl-PL

Portuguese;葡萄牙语;pt
Portuguese;Angola;安哥拉;pt-Ao
Portuguese;Cape Verde;佛得角;pt-CV
Portuguese;Guinea-Bissau;几内亚比绍共和国;pt-GW
Portuguese;Macau;澳门;pt-MO
Portuguese;Mozambique;莫桑比克;pt-MZ
Portuguese;Seio Tome & Principe;圣多美和普林西比;pt-ST
Portuguese;Timor-Leste;东帝汶;pt-TL

Punjabi;旁遮普语;pa
Punjabi;Arabic;阿拉伯;pa-Arab
Punjabi;Arabic, Pakistan;阿拉伯语，巴基斯坦;pa-Arab-PK
Punjabi;India;印度;pa-IN

Quechua;克丘亚语;qu
Quechua;Bolivia;玻利维亚;qu-BO
Quechua;Ecuador;厄瓜多尔;qu-EC
Quechua;Peru;秘鲁;qu-PE

Romanian;罗马尼亚语;ro
Romanian;Moldova;摩尔多瓦;ro-MD
Romanian;Romania;罗马尼亚;ro-RO

Romansh;罗曼什语;rm
Romansh;Switzerland;瑞士;rm-CH

Rombo; ;rof
Rombo;Tanzania;坦桑尼亚;rof-TZ

Rundi;隆迪语;rn
Rundi;Burundi;布隆迪;rn-Bl

Russian;俄语;ru
Russian;Belarus;白俄罗斯;ru-BY
Russian;Kazakhstan;哈萨克斯坦;ru-KZ
Russian;Kyrgyzstan;吉尔吉斯斯坦;ru-KG
Russian;Moldova;摩尔多瓦;ru-MD
Russian;Russia;俄罗斯;ru-RU
Russian;Ukraine;乌克兰;ru-UA

Rwa; ;rwk
Rwa;Tanzania;坦桑尼亚;rwk-TZ

Sakha; ;sah
Sakha;Russia;俄罗斯;sah-RU

Samburu;saq
Samburu;Kenya;肯尼亚;saq-KE

Samoan;萨摩亚语;sm

Sango;桑戈语;sg
Sango;Central African Republic;中非共和国;sg-CF

Sangu; ;sbp
Sangu;Tanzania;坦桑尼亚;sbp-TZ

Scottish Gaelic;苏格兰盖尔语;gd
Scottish Gaelic;United Kingdom;英国;gd-GB

Sesotho;塞索托语;st

Sena; ;seh
Sena;Mozambique;莫桑比克;seh-MZ

Serbian;塞尔维亚语;sr
Serbian;Bosnia & Herzegovina;波斯尼亚和黑塞哥维纳;sr-BA
Serbian;Kosovo;科索沃;sr-XK
Serbian;Latin;拉丁;sr-Latn
Serbian;Latin, Bosnia & Herzegovina;拉丁，波斯尼亚和黑塞哥维那;sr-Latn-BA
Serbian;Latin, Kosovo;拉丁，科索沃;sr-Latn-XK
Serbian;Latin, Montenegro;拉丁，黑山;sr-Latn-ME
Serbian;Latin, Serbia;拉丁，塞尔维亚;sr-Latn-RS
Serbian;Montenegro;黑山共和国;sr-ME
Serbian;Serbia;塞尔维亚;sr-RS

Shambala; ;ksb
Shambala;Tanzania;坦桑尼亚;ksb-TZ

Shona;修纳语;sn
Shona;Zimbabwe;津巴布韦;sn-ZW

Sichuan Yi;彝族语;ii
Sichuan Yi;China;中国;ii-CN

Sindhi;信德语;sd

Sinhala;僧伽罗语;si
Sinhala;Sri Lanka;斯里兰卡;si-LK

Slovak;斯洛伐克语;sk
Slovak;Slovakia;斯洛伐克;sk-SK

Slovenian;斯洛文尼亚语;sl
Slovenian;Slovenia;斯洛文尼亚;sl-Sl

Soga; ;xog
Soga;Uganda;乌干达;xog-UG

Somali;索马里语;so
Somali;Djibouti;吉布提;so-DJ
Somali;Ethiopia;埃塞俄比亚;so-ET
Somali;Kenya;肯尼亚;so-KE
Somali;Somalia;索马里;so-So

Spanish;西班牙语;es
Spanish;Argentina;阿根廷;es-AR
Spanish;Bolivia;玻利维亚;es-BO
Spanish;Canary Islands;加那利群岛;es-IC
Spanish;Ceuta & Melilla;休达和梅利利亚;es-EA
Spanish;Chile;智利;es-CL
Spanish;Colombia;哥伦比亚;es-Co
Spanish;Costa Rica;哥斯达黎加;es-CR
Spanish;Cuba;古巴;es-CU
Spanish;Dominican Republic;多米尼加共和国;es-Do
Spanish;Ecuador;厄瓜多尔;es-EC
Spanish;El Salvador;萨尔瓦多;es-SV
Spanish;Equatorial Guinea;赤道几内亚;es-GQ
Spanish;Guatemala;危地马拉;es-GT
Spanish;Honduras;洪都拉斯;es-HN
Spanish;Latin America;拉丁美洲;es-419
Spanish;Nicaragua;尼加拉瓜;es-NI
Spanish;Panama;巴拿马;es-PA
Spanish;Paraguay;巴拉圭;es-PY
Spanish;Peru;秘鲁;es-PE
Spanish;Philippines;菲律宾;es-PH
Spanish;Puerto Rico;波多黎各;es-PR
Spanish;Spain;西班牙;es-ES
Spanish;United States;美国;es-US
Spanish;Uruguay;乌拉圭;es-UY
Spanish;Venezuela;委内瑞拉;es-VE

Standard Moroccan Tamazight;摩洛哥塔马塞特语;zgh
Standard Moroccan Tamazight;Morocco;摩洛哥;zgh-MA

Sundanese;印尼巽他语;su

Swahili;斯瓦希里语;sw
Swahili;Congo - Kinshasa;刚果-金沙萨;sw-CD
Swahili;Kenya;肯尼亚;sw-KE
Swahili;Tanzania;坦桑尼亚;sw-TZ
Swahili;Uganda;乌干达;sw-UG

Swedish;瑞典语;sv
Swedish;Aland Islands;奥兰群岛;sv-AX
Swedish;Finland;芬兰;sv-Fl
Swedish;Sweden;瑞典;sv-SE

Swiss German;瑞士德语;gsw
Swiss German;France;法国;gsw-FR
Swiss German;Liechtenstein;列支敦士登;gsw-LI
Swiss German;Switzerland;瑞士;gsw-CH

Tachelhit; ;shi
Tachelhit;Morocco;摩洛哥;shi-MA
Tachelhit;Tifinagh;提非纳文;shi-Tfng
Tachelhit;Tifinagh, Morocco;提非纳文，摩洛哥;shi-Tfng-MA

Taita; ;dav
Taita;Kenya;肯尼亚;dav-KE

Tajik;塔吉克语;tg
Tajik;Tajikistan;塔吉克斯坦;tg-TJ

Tamil;泰米尔语;ta
Tamil;India;印度;ta-IN
Tamil;Malaysia;马来西亚;ta-MY
Tamil;Singapore;新加坡;ta-SG
Tamil;Sri Lanka;斯里兰卡;ta-LK

Tasawaq; ;twq
Tasawaq;Niger;尼日尔;twq-NE

Telugu;泰卢固语;te
Telugu;India;印度;te-IN

Teso; ;teo
Teso;Kenya;肯尼亚;teo-KE
Teso;Uganda;乌干达;teo-UG

Thai;泰语;th
Thai;Thailand;泰国;th-TH

Tibetan;藏语;bo
Tibetan;China;中国;bo-CN
Tibetan;India;印度;bo-IN

Tigrinya;提格里尼亚语;ti
Tigrinya;Eritrea;厄立特里亚;ti-ER
Tigrinya;Ethiopia;埃塞俄比亚;ti-ET

Tongan;汤加语;to
Tongan;Tonga;汤加;to-TO

Turkish;土耳其语;tr
Turkish;Cyprus;塞浦路斯共和国;tr-CY
Turkish;Turkey;土耳其;tr-TR

Turkmen;土库曼语;tk
Turkmen;Turkmenistan;土库曼斯坦;tk-TM

Ukrainian;乌克兰语;uk
Ukrainian;Ukraine;乌克兰;uk-UA

Upper Sorbian;上索布语;hsb
Upper Sorbian;Germany;德国;hsb-DE

Urdu;乌尔都语;ur
Urdu;India;印度;ur-IN
Urdu;Pakistan;巴基斯坦;ur-PK

Uyghur;维吾尔语;ug
Uyghur;Arabic;阿拉伯;ug-Arab
Uyghur;Arabic, China;阿拉伯，中国;ug-Arab-CN

Uzbek;乌兹别克语;uz
Uzbek;Arabic;阿拉伯;uz-Arab
Uzbek;Arabic, Afghanistan;阿拉伯，阿富汗;uz-Arab-AF
Uzbek;Latin;拉丁;uz-Latn
Uzbek;Latin, Uzbekistan;拉丁，乌兹别克斯坦;uz-Latn-UZ
Uzbek;Uzbekistan;乌兹别克斯坦;uz-UZ

Vai;;;vai
Vai;Latin;拉丁;vai-Latn
Vai;Latin, Liberia;拉丁，利比里亚;vai-Latn-LR
Vai;Liberia;利比里亚;vai-LR

Vietnamese;越南语;vi
Vietnamese;Vietnam;越南;vi-VN

Vunjo; ;vun
Vunjo;Tanzania;坦桑尼亚;vun-TZ

Walser; ;wae
Walser;Switzerland;瑞士;wae-CH

Welsh;威尔士语;cy
Welsh;United Kingdom;英国;cy-GB

Western Frisian;西弗里西亚语;fy
Western Frisian;Netherlands;荷兰;fy-NL

Yangben; ;yav
Yangben;Cameroon;喀麦隆;yav-CM

Yiddish;意第绪语;yi
Yiddish;World;世界;vi-001

Yoruba;约鲁巴语;yo
Yoruba;Benin;贝宁湾;yo-BJ
Yoruba;Nigeria;尼日利亚;yo-NG

Xhosa;南非科萨语;xh

Zarma; ;dje
Zarma;Niger;尼日尔;dje-NE

Zulu;祖鲁语;zu
Zulu;South Africa;南非;zu-ZA

*/

/**
 *	get languages source data
 *
 */
func SourceData() string {
	return `	
Afrikaans; ;af
Afrikaans;Namibia; ;af-NA
Afrikaans;South Africa; ;af-ZA

Aghem; ;agq
Aghem;Cameroon; ;agq-CM

Akan; ;ak
Akan;Ghana; ;ak-GH

Albanian; ;sq
Albanian;Albania; ;sq-AL
Albanian;Kosovo; ;sq-XK
Albanian;Macedonia; ;sq-MK

Amharic; ;am
Amharic;Ethiopia; ;am-ET

Arabic; ;ar
Arabic;Algeria; ;ar-DZ
Arabic;Bahrain; ;ar-BH
Arabic;Chad; ;ar-TD
Arabic;Comoros; ;ar-KM
Arabic;Djibouti; ;ar-DJ
Arabic;Egypt; ;ar-EG
Arabic;Eritrea; ;ar-ER
Arabic;Iraq; ;ar-IQ
Arabic;Israel; ;ar-IL
Arabic;Jordan; ;ar-JO
Arabic;Kuwait; ;ar-KW
Arabic;Lebanon; ;ar-LB
Arabic;Libya; ;ar-LY
Arabic;Mauritania; ;ar-MR
Arabic;Morocco; ;ar-MA
Arabic;Oman; ;ar-OM
Arabic;Palestine; ;ar-PS
Arabic;Qatar; ;ar-QA
Arabic;Saudi Arabia; ;ar-SA
Arabic;Somalia; ;ar-SO
Arabic;South Sudan; ;ar-SS
Arabic;Sudan; ;ar-SD
Arabic;Syria; ;ar-SY
Arabic;Tunisia; ;ar-TN
Arabic;United Arab Emirates; ;ar-AE
Arabic;Western Sahara; ;ar-EH
Arabic;World; ;ar-001
Arabic;Yemen; ;ar-YE

Armenian; ;hy
Armenian;Armenia; ;hy-AM

Assamese; ;as
Assamese;India; ;as-IN

Asu; ;asa
Asu;Tanzania; ;asa-TZ

Azerbaijani; ;az
Azerbaijani;Azerbaijan; ;az-AZ
Azerbaijani;Cyrillic; ;az-Cyrl
Azerbaijani;Cyrillic, Azerbaijan; ;az-CyrI-AZ

Bafia; ;ksf
Bafia;Cameroon; ;ksf-CM

Bambara; ;bm
Bambara;Mali; ;bm-ML

Basaa; ;bas
Basaa;Cameroon; ;bas-CM

Basque; ;eu
Basque;Spain; ;eu-ES

Belarusian; ;be
Belarusian;Belarus; ;be-BY

Bemba; ;bem
Bemba;Zambia; ;bem-ZM

Bena; ;bez
Bena;Tanzania; ;bez-TZ

Bengali; ;bn
Bengali;Bangladesh; ;bn-BD
Bengali;India; ;bn-IN

Bodo; ;brx
Bodo;India; ;brx-IN

Bosnian; ;bs
Bosnian;Bosnia & Herzegovina; ;bs-BA
Bosnian;Cyrillic; ;bs-Cyrl
Bosnian;Cyrillic, Bosnia & Herzegovina; ;bs-Cyrl-BA

Breton; ;br
Breton;France; ;br-FR

Bulgarian; ;bg
Bulgarian;Bulgaria; ;bg-BG

Burmese; ;my
Burmese;Myanmar; ;my-MM

Catalan; ;ca
Catalan;Andorra; ;ca-AD
Catalan;France; ;ca-FR
Catalan;Italy; ;ca-IT
Catalan;Spain; ;ca-ES

Cebuano; ;ceb

Central Atlas Tamazight; ;tzm
Central Atlas Tamazight;Morocco; ;tzm-MA

Central Kurdish; ;ckb
Central Kurdish;Iran; ;ckb-IR
Central Kurdish;Iraq; ;ckb-IQ

Cherokee; ;chr
Cherokee;United States; ;chr-US

Chichewa; ;ny

Chiga; ;cgg
Chiga;Uganda; ;cgg-UG

Chinese Simplified; ;zh-Hans
Chinese Traditional; ;zh-Hant
Chinese;Simplified, China; ;zh-CN
Chinese;Simplified, Hong Kong; ;zh-Hans-HK
Chinese;Simplified, Macau; ;zh-Hans-MO
Chinese;Simplified, Singapore; ;zh-Hans-SG
Chinese;Simplified, Taiwan; ;zh-Hans-TW
Chinese;Traditional, Hong Kong; ;zh-Hant-HK
Chinese;Traditional, Macau; ;zh-Hant-MO
Chinese;Traditional, Singapore; ;zh-Hant-SG
Chinese;Traditional, Taiwan; ;zh-Hant-TW

Colognian; ;ksh
Colognian;Germany; ;ksh-DE

Cornish; ;kw
Cornish;United Kingdom; ;kw-GB

Corsican; ;co

Croatian; ;hr
Croatian;Bosnia & Herzegovina; ;hr-BA
Croatian;Croatia; ;hr-HR

Czech; ;cs
Czech;Czech Republic; ;cs-CZ

Danish; ;da
Danish;Denmark; ;da-DK
Danish;Greenland; ;da-GL

Duala; ;dua
Duala;Cameroon; ;dua-CM

Dutch; ;nl
Dutch;Aruba; ;nl-AW
Dutch;Belgium; ;nl-BE
Dutch;Caribbean Netherlands; ;nl-BQ
Dutch;Curacao; ;nl-CW
Dutch;Netherlands; ;nl-NL
Dutch;Sint Maarten; ;nl-SX
Dutch;Suriname; ;nl-SR

Dzongkha; ;dz
Dzongkha;Bhutan; ;dz-BT

Embu; ;ebu
Embu;Kenya; ;ebu-KE

English; ;en
English;Albania; ;en-AL
English;American Samoa; ;en-AS
English;Andorra; ;en-AD
English;Anguilla; ;en-Al
English;Antigua & Barbuda; ;en-AG
English;Australia; ;en-AU
English;Austria; ;en-AT
English;Bahamas; ;en-BS
English;Barbados; ;en-BB
English;Belgium; ;en-BE
English;Belize; ;en-BZ
English;Bermuda; ;en-BM
English;Bosnia & Herzegovina; ;en-BA
English;Botswana; ;en-BW
English;British Indian Ocean Territory; ;en-IO
English;British Virgin Islands; ;en-VG
English;Cameroon; ;en-CM
English;Canada; ;en-CA
English;Cayman Islands; ;en-KY
English;Christmas Island; ;en-CX
English;Cocos [Keeling] Islands; ;en-CC
English;Cook Islands; ;en-CK
English;Croatia; ;en-HR
English;Cyprus; ;en-CY
English;Czech Republic; ;en-CZ
English;Denmark; ;en-DK
English;Diego Garcia; ;en-DG
English;Dominica; ;en-DM
English;Eritrea; ;en-ER
English;Estonia; ;en-EE
English;Europe; ;en-150
English;Falkland Islands; ;en-FK
English;Fiji; ;en-FJ
English;Finland; ;en-Fl
English;France; ;en-FR
English;Gambia; ;en-GM
English;Germany; ;en-DE
English;Ghana; ;en-GH
English;Gibraltar; ;en-GI
English;Greece; ;en-GR
English;Grenada; ;en-GD
English;Guam; ;en-GU
English;Guernsey; ;en-GG
English;Guyana; ;en-GY
English;Hong Kong; ;en-HK
English;Hungary; ;en-HU
English;Iceland; ;en-IS
English;India; ;en-IN
English;Ireland; ;en-IE
English;Isle of Man; ;en-IM
English;Israel; ;en-IL
English;Italy; ;en-IT
English;Jamaica; ;en-JM
English;Jersey; ;en-JE
English;Kenya; ;en-KE
English;Kiribati; ;en-KI
English;Latvia; ;en-LV
English;Lesotho; ;en-LS
English;Liberia; ;en-LR
English;Lithuania; ;en-LT
English;Luxembourg; ;en-LU
English;Macau; ;en-MO
English;Madagascar; ;en-MG
English;Malawi; ;en-MW
English;Malaysia; ;en-MY
English;Malta; ;en-MT
English;Marshall Islands; ;en-MH
English;Mauritius; ;en-MU
English;Micronesia; ;en-FM
English;Montenegro; ;en-ME
English;Montserrat; ;en-MS
English;Namibia; ;en-NA
English;Nauru; ;en-NR
English;Netherlands; ;en-NL
English;New Zealand; ;en-NZ
English;Nigeria; ;en-NG
English;Niue; ;en-NU
English;Norfolk Island; ;en-NF
English;Northern Mariana Islands; ;en-MP
English;Norway; ;en-No
English;Pakistan; ;en-PK
English;Palau; ;en-PW
English;Papua New Guinea; ;en-PG
English;Philippines; ;en-PH
English;Pitcairn Islands; ;en-PN
English;Poland; ;en-PL
English;Portugal; ;en-PT
English;Puerto Rico; ;en-PR
English;Romania; ;en-Ro
English;Russia; ;en-RU
English;Rwanda; ;en-RW
English;Samoa; ;en-WS
English;Seychelles; ;en-SC
English;Sierra Leone; ;en-SL
English;Singapore; ;en-SG
English;Sint Maarten; ;en-SX
English;Slovakia; ;en-SK
English;Slovenia; ;en-SI
English;Solomon Islands; ;en-SB
English;South Africa; ;en-ZA
English;South Sudan; ;en-SS
English;Spain; ;en-ES
English;St. Helena; ;en-SH
English;St. Kitts & Nevis; ;en-KN
English;St. Lucia; ;en-LC
English;St. Vincent & Grenadines; ;en-VC
English;Sudan; ;en-SD
English;Swaziland; ;en-SZ
English;Sweden; ;en-SE
English;Switzerland; ;en-CH
English;Tanzania; ;en-TZ
English;Tokelau; ;en-TK
English;Tonga; ;en-TO
English;Trinidad & Tobago; ;en-TT
English;Turkey; ;en-TR
English;Turks & Caicos Islands; ;en-TC
English;Tuvalu; ;en-TV
English;US Outlying Islands; ;en-UM
English;US Virgin Islands; ;en-VI
English;Uganda; ;en-UG
English;United Kingdom; ;en-GB
English;United States; ;en-US
English;United States, Computer; ;en-US-POSIX
English;Vanuatu; ;en-VU
English;World; ;en-001
English;Zambia; ;en-ZM
English;Zimbabwe; ;en-ZW

Esperanto; ;eo

Estonian; ;et
Estonian;Estonia; ;et-EE

Ewe; ;ee
Ewe;Ghana; ;ee-GH
Ewe;Togo; ;ee-TG

Ewondo; ;ewo
Ewondo;Cameroon; ;ewo-CM

Faroese; ;fo
Faroese;Faroe Islands; ;fo-FO

Filipino; ;fil
Filipino;Philippines; ;fil-PH

Finnish; ;fi
Finnish;Finland; ;fi-Fl

French; ;fr
French;Algeria; ;fr-DZ
French;Belgium; ;fr-BE
French;Benin; ;fr-BJ
French;Burkina Faso; ;fr-BF
French;Burundi; ;fr-Bl
French;Cameroon; ;fr-CM
French;Canada; ;fr-CA
French;Central African Republic; ;fr-CF
French;Chad; ;fr-TD
French;Comoros; ;fr-KM
French;Congo - Brazzaville; ;fr-CG
French;Congo - Kinshasa; ;fr-CD
French;Cote d’lvoire; ;fr-Cl
French;Djibouti; ;fr-DJ
French;Equatorial Guinea; ;fr-GQ
French;France; ;fr-FR
French;French Guiana; ;fr-GF
French;French Polynesia; ;fr-PF
French;Gabon; ;fr-GA
French;Guadeloupe; ;fr-GP
French;Guinea; ;fr-GN
French;Haiti; ;fr-HT
French;Luxembourg; ;fr-LU
French;Madagascar; ;fr-MG
French;Mali; ;fr-ML
French;Martinique; ;fr-MQ
French;Mauritania; ;fr-MR
French;Mauritius; ;fr-MU
French;Mayotte; ;fr-YT
French;Monaco; ;fr-MC
French;Morocco; ;fr-MA
French;New Caledonia; ;fr-NC
French;Niger; ;fr-NE
French;Reunion; ;fr-RE
French;Rwanda; ;fr-RW
French;Senegal; ;fr-SN
French;Seychelles; ;fr-SC
French;St. Barthelemy; ;fr-BL
French;St. Martin; ;fr-MF
French;St. Pierre & Miquelon; ;fr-PM
French;Switzerland; ;fr-CH
French;Syria; ;fr-SY
French;Togo; ;fr-TG
French;Tunisia; ;fr-TN
French;Vanuatu; ;fr-VU
French;Wallis & Futuna; ;fr-WF

Friulian; ;fur
Friulian;Italy; ;fur-IT

Fulah; ;ff
Fulah;Cameroon; ;ff-CM
Fulah;Guinea; ;ff-GN
Fulah;Mauritania; ;ff-MR
Fulah;Senegal; ;ff-SN

Galician; ;gl
Galician;Spain; ;gl-ES

Ganda; ;lg
Ganda;Uganda; ;lg-UG

Georgian; ;ka
Georgian;Georgia; ;ka-GE

German; ;de
German;Austria; ;de-AT
German;Belgium; ;de-BE
German;Germany; ;de-DE
German;Liechtenstein; ;de-LI
German;Luxembourg; ;de-LU
German;Switzerland; ;de-CH

Greek; ;el
Greek;Cyprus; ;el-CY
Greek;Greece; ;el-GR

Gujarati; ;gu
Gujarati;India; ;gu-IN

Gusii; ;guz
Gusii;Kenya; ;guz-KE

Haitian Creole; ;ht

Hausa; ;ha
Hausa;Ghana; ;ha-GH
Hausa;Niger; ;ha-NE
Hausa;Nigeria; ;ha-NG

Hawaiian; ;haw
Hawaiian;United States; ;haw-US

Hebrew; ;he
Hebrew;Israel; ;he-IL

Hindi; ;hi
Hindi;India; ;hi-IN

Hmong; ;hmong

Hungarian; ;hu
Hungarian;Hungary; ;hu-HU

Icelandic; ;is
Icelandic;Iceland; ;is-IS

Igbo; ;ig
Igbo;Nigeria; ;ig-NG

Inari Sami; ;smn
Inari Sami;Finland; ;smn-FI

Indonesian; ;id
Indonesian;Indonesia; ;id-ID

Inuktitut; ;iu
Inuktitut;Unified Canadian Aboriginal Syllabics; ;iu-Cans
Inuktitut;Unified Canadian Aboriginal Syllabics, Canada; ;iu-Cans-CA

Irish; ;ga
Irish;Ireland; ;ga-IE

Italian; ;it
Italian;Italy; ;it-IT
Italian;San Marino; ;it-SM
Italian;Switzerland; ;it-CH

Japanese; ;ja
Japanese;Japan; ;ja-JP

Javanese; ;jv

Jola-Fonyi; ;dyo
Jola-Fonyi;Senegal; ;dyo-SN

Kabuverdianu; ;kea
Kabuverdianu;Cape Verde; ;kea-CV

Kabyle; ;kab
Kabyle;Algeria; ;kab-DZ

Kako; ;kkj
Kako;Cameroon; ;kkj-CM

Kalaallisut; ;kl
Kalaallisut;Greenland; ;kl-GL

Kalenjin; ;kln
Kalenjin;Kenya; ;kln-KE

Kamba; ;kam
Kamba;Kenya; ;kam-KE

Kannada; ;kn
Kannada;India; ;kn-IN

Kashmiri; ;ks
Kashmiri;Arabic; ;ks-Arab
Kashmiri;Arabic, India; ;ks-Arab-IN

Kazakh; ;kk
Kazakh;Kazakhstan; ;kk-KZ

Khmer; ;km
Khmer;Cambodia; ;km-KH

Kikuyu; ;ki
Kikuyu;Kenya; ;ki-KE

Kinyarwanda; ;rw
Kinyarwanda;Rwanda; ;rw-RW

Konkani; ;kok
Konkani;India; ;kok-IN

Korean; ;ko
Korean;North Korea; ;ko-KP
Korean;South Korea; ;ko-KR

Koyra Chiini; ;khq
Koyra Chiini;Mali; ;khq-ML

Koyraboro Senni; ;ses
Koyraboro Senni;Mali; ;ses-ML

Kwasio; ;nmg
Kwasio;Cameroon; ;nmg-CM

Kyrgyz; ;ky
Kyrgyz;Kyrgyzstan; ;ky-KG

Lakota; ;lkt
Lakota;United States; ;lkt-US

Langi; ;lag
Langi;Tanzania; ;lag-TZ

Lao; ;lo
Lao;Laos; ;lo-LA

Latin; ;la

Latvian; ;lv
Latvian;Latvia; ;lv-LV

Lingala; ;ln
Lingala;Angola; ;ln-Ao
Lingala;Central African Republic; ;ln-CF
Lingala;Congo - Brazzaville; ;ln-CG
Lingala;Congo - Kinshasa; ;ln-CD

Lithuanian; ;lt
Lithuanian;Lithuania; ;lt-LT

Lower Sorbian; ;dsb
Lower Sorbian;Germany; ;dsb-DE

Luba-Katanga; ;lu
Luba-Katanga;Congo - Kinshasa; ;lu-CD

Luo; ;luo
Luo;Kenya; ;luo-KE

Luxembourgish; ;lb
Luxembourgish;Luxembourg; ;lb-LU

Luyia; ;luy
Luyia;Kenya; ;luy-KE

Macedonian; ;mk
Macedonian;Macedonia; ;mk-MK

Machame; ;jmc
Machame;Tanzania; ;jmc-TZ

Makhuwa-Meetto;mgh
Makhuwa-Meetto;Mozambique; ;mgh-MZ

Makonde; ;kde
Makonde;Tanzania; ;kde-TZ

Maori; ;mi

Malagasy; ;mg
Malagasy;Madagascar; ;mg-MG

Malay; ;ms
Malay;Arabic; ;ms-Arab
Malay;Arabic, Brunei; ;ms-Arab-BN
Malay;Arabic, Malaysia; ;ms-Arab-MY
Malay;Brunei; ;ms-BN
Malay;Malaysia; ;ms-MY
Malay;Singapore; ;ms-SG

Malayalam; ;ml
Malayalam;India; ;ml-IN

Maltese; ;mt
Maltese;Malta; ;mt-MT

Manx; ;gv
Manx;Isle of Man; ;gv-IM

Marathi; ;mr
Marathi;India; ;mr-IN

Masai; ;mas
Masai;Kenya; ;mas-KE
Masai;Tanzania; ;mas-TZ

Meru; ;mer
Meru;Kenya; ;mer-KE

Meta'; ;mgo
Meta';Cameroon; ;mgo-CM

Mongolian; ;mn
Mongolian;Mongolia; ;mn-MN

Morisyen; ;mfe
Morisyen;Mauritius; ;mfe-MU

Mundang; ;mua
Mundang;Cameroon; ;mua-CM

Nama; ;naq
Nama;Namibia; ;naq-NA

Nepali; ;ne
Nepali;India; ;ne-IN
Nepali;Nepal; ;ne-NP

Ngiemboon; ;nnh
Ngiemboon;Cameroon; ;nnh-CM

Ngomba; ;jgo
Ngomba;Cameroon; ;jgo-CM

North Ndebele; ;nd
North Ndebele;Zimbabwe; ;nd-ZW

Northern Sami; ;se
Northern Sami;Finland; ;se-FI
Northern Sami;Norway; ;se-NO
Northern Sami;Sweden; ;se-SE

Norwegian Bokmal; ;nb
Norwegian Bokmal;Norway; ;nb-NO
Norwegian Bokmal;Svalbard & Jan Mayen; ;nb-SJ

Norwegian Nynorsk; ;nn
Norwegian Nynorsk;Norway; ;nn-NO

Nuer; ;nus
Nuer;Sudan; ;nus-SD

Nyankole; ;nyn
Nyankole;Uganda; ;nyn-UG

Oriya; ;or
Oriya;India; ;or-IN

Oromo; ;om
Oromo;Ethiopia; ;om-ET
Oromo;Kenya; ;om-KE

Ossetic; ;os
Ossetic;Georgia; ;os-GE
Ossetic;Russia; ;os-RU

Pashto; ;ps
Pashto;Afghanistan; ;ps-AF

Persian; ;fa
Persian;Afghanistan; ;fa-AF
Persian;Iran; ;fa-IR

Polish; ;pl
Polish;Poland; ;pl-PL

Portuguese; ;pt
Portuguese;Angola; ;pt-Ao
Portuguese;Cape Verde; ;pt-CV
Portuguese;Guinea-Bissau; ;pt-GW
Portuguese;Macau; ;pt-MO
Portuguese;Mozambique; ;pt-MZ
Portuguese;Seio Tome & Principe; ;pt-ST
Portuguese;Timor-Leste; ;pt-TL

Punjabi; ;pa
Punjabi;Arabic; ;pa-Arab
Punjabi;Arabic, Pakistan; ;pa-Arab-PK
Punjabi;India; ;pa-IN

Quechua; ;qu
Quechua;Bolivia; ;qu-BO
Quechua;Ecuador; ;qu-EC
Quechua;Peru; ;qu-PE

Romanian; ;ro
Romanian;Moldova; ;ro-MD
Romanian;Romania; ;ro-RO

Romansh; ;rm
Romansh;Switzerland; ;rm-CH

Rombo; ;rof
Rombo;Tanzania; ;rof-TZ

Rundi; ;rn
Rundi;Burundi; ;rn-Bl

Russian; ;ru
Russian;Belarus; ;ru-BY
Russian;Kazakhstan; ;ru-KZ
Russian;Kyrgyzstan; ;ru-KG
Russian;Moldova; ;ru-MD
Russian;Russia; ;ru-RU
Russian;Ukraine; ;ru-UA

Rwa; ;rwk
Rwa;Tanzania; ;rwk-TZ

Sakha; ;sah
Sakha;Russia; ;sah-RU

Samburu;saq
Samburu;Kenya; ;saq-KE

Samoan; ;sm

Sango; ;sg
Sango;Central African Republic; ;sg-CF

Sangu; ;sbp
Sangu;Tanzania; ;sbp-TZ

Scottish Gaelic; ;gd
Scottish Gaelic;United Kingdom; ;gd-GB

Sesotho; ;st

Sena; ;seh
Sena;Mozambique; ;seh-MZ

Serbian; ;sr
Serbian;Bosnia & Herzegovina; ;sr-BA
Serbian;Kosovo; ;sr-XK
Serbian;Latin; ;sr-Latn
Serbian;Latin, Bosnia & Herzegovina; ;sr-Latn-BA
Serbian;Latin, Kosovo; ;sr-Latn-XK
Serbian;Latin, Montenegro; ;sr-Latn-ME
Serbian;Latin, Serbia; ;sr-Latn-RS
Serbian;Montenegro; ;sr-ME
Serbian;Serbia; ;sr-RS

Shambala; ;ksb
Shambala;Tanzania; ;ksb-TZ

Shona; ;sn
Shona;Zimbabwe; ;sn-ZW

Sichuan Yi; ;ii
Sichuan Yi;China; ;ii-CN

Sindhi; ;sd

Sinhala; ;si
Sinhala;Sri Lanka; ;si-LK

Slovak; ;sk
Slovak;Slovakia; ;sk-SK

Slovenian; ;sl
Slovenian;Slovenia; ;sl-Sl

Soga; ;xog
Soga;Uganda; ;xog-UG

Somali; ;so
Somali;Djibouti; ;so-DJ
Somali;Ethiopia; ;so-ET
Somali;Kenya; ;so-KE
Somali;Somalia; ;so-So

Spanish; ;es
Spanish;Argentina; ;es-AR
Spanish;Bolivia; ;es-BO
Spanish;Canary Islands; ;es-IC
Spanish;Ceuta & Melilla; ;es-EA
Spanish;Chile; ;es-CL
Spanish;Colombia; ;es-Co
Spanish;Costa Rica; ;es-CR
Spanish;Cuba; ;es-CU
Spanish;Dominican Republic; ;es-Do
Spanish;Ecuador; ;es-EC
Spanish;El Salvador; ;es-SV
Spanish;Equatorial Guinea; ;es-GQ
Spanish;Guatemala; ;es-GT
Spanish;Honduras; ;es-HN
Spanish;Latin America; ;es-419
Spanish;Nicaragua; ;es-NI
Spanish;Panama; ;es-PA
Spanish;Paraguay; ;es-PY
Spanish;Peru; ;es-PE
Spanish;Philippines; ;es-PH
Spanish;Puerto Rico; ;es-PR
Spanish;Spain; ;es-ES
Spanish;United States; ;es-US
Spanish;Uruguay; ;es-UY
Spanish;Venezuela; ;es-VE

Standard Moroccan Tamazight; ;zgh
Standard Moroccan Tamazight;Morocco; ;zgh-MA

Sundanese; ;su

Swahili; ;sw
Swahili;Congo - Kinshasa; ;sw-CD
Swahili;Kenya; ;sw-KE
Swahili;Tanzania; ;sw-TZ
Swahili;Uganda; ;sw-UG

Swedish; ;sv
Swedish;Aland Islands; ;sv-AX
Swedish;Finland; ;sv-Fl
Swedish;Sweden; ;sv-SE

Swiss German; ;gsw
Swiss German;France; ;gsw-FR
Swiss German;Liechtenstein; ;gsw-LI
Swiss German;Switzerland; ;gsw-CH

Tachelhit; ;shi
Tachelhit;Morocco; ;shi-MA
Tachelhit;Tifinagh; ;shi-Tfng
Tachelhit;Tifinagh, Morocco; ;shi-Tfng-MA

Taita; ;dav
Taita;Kenya; ;dav-KE

Tajik; ;tg
Tajik;Tajikistan; ;tg-TJ

Tamil; ;ta
Tamil;India; ;ta-IN
Tamil;Malaysia; ;ta-MY
Tamil;Singapore; ;ta-SG
Tamil;Sri Lanka; ;ta-LK

Tasawaq; ;twq
Tasawaq;Niger; ;twq-NE

Telugu; ;te
Telugu;India; ;te-IN

Teso; ;teo
Teso;Kenya; ;teo-KE
Teso;Uganda; ;teo-UG

Thai; ;th
Thai;Thailand; ;th-TH

Tibetan; ;bo
Tibetan;China; ;bo-CN
Tibetan;India; ;bo-IN

Tigrinya; ;ti
Tigrinya;Eritrea; ;ti-ER
Tigrinya;Ethiopia; ;ti-ET

Tongan; ;to
Tongan;Tonga; ;to-TO

Turkish; ;tr
Turkish;Cyprus; ;tr-CY
Turkish;Turkey; ;tr-TR

Turkmen; ;tk
Turkmen;Turkmenistan; ;tk-TM

Ukrainian; ;uk
Ukrainian;Ukraine; ;uk-UA

Upper Sorbian; ;hsb
Upper Sorbian;Germany; ;hsb-DE

Urdu; ;ur
Urdu;India; ;ur-IN
Urdu;Pakistan; ;ur-PK

Uyghur; ;ug
Uyghur;Arabic; ;ug-Arab
Uyghur;Arabic, China; ;ug-Arab-CN

Uzbek; ;uz
Uzbek;Arabic; ;uz-Arab
Uzbek;Arabic, Afghanistan; ;uz-Arab-AF
Uzbek;Latin; ;uz-Latn
Uzbek;Latin, Uzbekistan; ;uz-Latn-UZ
Uzbek;Uzbekistan; ;uz-UZ

Vai; ;vai
Vai;Latin; ;vai-Latn
Vai;Latin, Liberia; ;vai-Latn-LR
Vai;Liberia; ;vai-LR

Vietnamese; ;vi
Vietnamese;Vietnam; ;vi-VN

Vunjo; ;vun
Vunjo;Tanzania; ;vun-TZ

Walser; ;wae
Walser;Switzerland; ;wae-CH

Welsh; ;cy
Welsh;United Kingdom; ;cy-GB

Western Frisian; ;fy
Western Frisian;Netherlands; ;fy-NL

Yangben; ;yav
Yangben;Cameroon; ;yav-CM

Yiddish; ;yi
Yiddish;World; ;vi-001

Yoruba; ;yo
Yoruba;Benin; ;yo-BJ
Yoruba;Nigeria; ;yo-NG

Xhosa; ;xh

Zarma; ;dje
Zarma;Niger; ;dje-NE

Zulu; ;zu
Zulu;South Africa; ;zu-ZA
`
}
