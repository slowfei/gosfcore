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
 *	load language directory info
 */
func LoadLanguages(tagName, dirPath string) (ILocalize, error) {

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
