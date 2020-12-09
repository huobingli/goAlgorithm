package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"strings"
	"strconv"
	"unsafe"
)

// decode unicode to utf8  /u
func unicode2utf8(source string) string {
	var res = []string{""}
	sUnicode := strings.Split(source, "\\u")
	var context = ""
	for _, v := range sUnicode {
		var additional = ""
		if len(v) < 1 {
			continue
		}
		if len(v) > 4 {
			rs := []rune(v)
			v = string(rs[:4])
			additional = string(rs[4:])
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			context += v
		}
		context += fmt.Sprintf("%c", temp)
		context += additional
	}
	res = append(res, context)
	return strings.Join(res, "")
}

func GetHtml() string {
	url := "http://ft.10jqka.com.cn/thsft/iFindService/Chain/industry-map/get-all-company?chart_id=39323263116075000051554&chart_type=qc&ismine=0"
	//url := "http://gpc.10jqka.com.cn/dptrend/window/textinfo/windindicator"
	//url:="http://gdzx.10jqka.com.cn/sfzx/cpzg/index/version/1?usecef=1"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	// testjrds
	req.Header.Set("Cookie", "searchGuide=sg; Hm_lvt_78c58f01938e4d85eaf619eae71b4ed1=1605162252,1606104757,1606300022; user=MDp0ZXN0anJkczQ6Ok5vbmU6NTAwOjM5MzIzMjYzMTo0LDEwMDAwMDAwMTAwMTAwMDAwMDExMTExMSwyODA7NSwxLDQwOzYsMSw0MDs3LDExMTExMTExMTExMCw0MDs4LDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAxMDAwMDEwMDExLDQwOzIwLDEsMjgwOzIyLDEsMjgwOzI2LDEsMjgwOzI3LDEsMjgwOzI5LDEsMjgwOzMzLDAwMTEwMDAwMDAxMSwyODA7MzYsMTAwMTExMTExMDAwMTExMTEwMTExMTExLDI4MDs0MSwxMTExMSwyODA7NDIsMTAwMTAxLDI4MDs0MywxMDExMTExMTExMTExMTExMTExMTExMTEsMjgwOzQ0LDExMTExLDQwOzQ1LDEwMTExMTExMSwyODA7NDYsMDExMDExMTExMDAwMDAxMTExMTExMTExLDI4MDs0NywxMTExMTExMTExMTExMTEsMjgwOzQ4LDExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExLDI4MDs1MCwxMSwyODA7NTEsMTEwMDAwMDAwMDAwMDAwMCwyODA7NTIsMDExMDAsMjgwOzU0LDEwMDAwMDAwMDAwMDAwMDAwMDAwMDAxMDAwMDAwMDAwLDI4MDs1NiwxMSwyODA7NTcsMTExMTExMTExMTExMTExMTExMTExMTExMTExMSwyODA7NTgsMTExMTExMTExMTExMTExMTEsMjgwOzYwLDExMTExMTExMTExMTExMTExLDI4MDs2MSwxMSwyODA7NjIsMTEwMDAwMDAwMDAwMDAwMDEwMDAwMDAwLDI4MDs2MywxMDAxMDAwMTAwMDAwMDAwMDAwMDAwMDAsMjgwOzY0LDAwMDExMTEwMDEwMDAwMDAwMDAxMDAwMCwyODA7NjYsMTExMTExMTExMTExMTAsMjgwOzY4LDExMTExMTExMTAxMTExMTExMDAwLDI4MDs3MywxMTExMDAsMjgwOzc3LDEwMDAwMTAsMjgwOzc4LDEsMjgwOzkxLDExMTExMTExMTEsMjgwOzkyLDAwMDAwMTAwLDI4MDsxMDMsMTExMTExMTExMTEsMjgwOzEsMTAxLDQwOzIsMSw0MDszLDEsNDA7MTAyLDEsNDA6OTo6OjM4MzIzMjYzMToxNjA3MzkxOTgzOjo6MTQ4OTAzODE4MDoyMjc2MTc6MDoxODI2YTAzYjBiNjNlYWI3MzBiM2NlMDY3OTU0MjBhNWE6OjA%3D; userid=383232631; u_name=testjrds4; escapename=testjrds4; ticket=98054b818cbc25e2ea023cbd9c9a542d; user_status=0; @#!userid!#@=383232631; @#!sessionid!#@=1774d432ff741abbd931d3d9a22b50696; @#!rsa_version!#@=default_4; v=Avqxdsu9caVWx_3JQjg9-VtbSyseq36F8C_yKQTzpg1Y95CV7DvOlcC_QjvX; cyl_sess_key=f7af1d1a340e01867380cd11e70bcf95%7Cjrds%7Ctestjrds4%7C393232631")
	//req.Header.Set("Cookie", "searchGuide=sg; __utma=156575163.232964267.1582016786.1582016786.1582016786.1; Hm_lvt_78c58f01938e4d85eaf619eae71b4ed1=1598589480,1599119768; user=MDp0ZXN0anJkczQ6Ok5vbmU6NTAwOjM5MzIzMjYzMTo0LDEwMDAwMDAwMTAwMTAwMDAwMDExMTExMSwzNTc7NSwxLDQwOzYsMSw0MDs3LDExMTExMTExMTExMCw0MDs4LDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAxMDAwMDEwMDExLDQwOzIwLDEsMzU3OzIyLDEsMzU3OzI2LDEsMzU3OzI3LDEsMzU3OzI5LDEsMzU3OzMzLDAwMTEwMDAwMDAxMSwzNTc7MzYsMTAwMTExMTExMDAwMTExMTEwMTExMTExLDM1Nzs0MSwxMTExMSwzNTc7NDIsMTAwMTAxLDM1Nzs0MywxMDExMTExMTExMTExMTExMTExMTExMTEsMzU3OzQ0LDExMTExLDQwOzQ1LDEwMTExMTExMSwzNTc7NDYsMDExMDExMTExMDAwMDAxMTExMTExMTExLDM1Nzs0NywxMTExMTExMTExMTExMTEsMzU3OzQ4LDExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExLDM1Nzs1MCwxMSwzNTc7NTEsMTEwMDAwMDAwMDAwMDAwMCwzNTc7NTIsMDExMDAsMzU3OzU0LDEwMDAwMDAwMDAwMDAwMDAwMDAwMDAxMDAwMDAwMDAwLDM1Nzs1NiwxMSwzNTc7NTcsMTExMTExMTExMTExMTExMTExMTExMTExMTExMSwzNTc7NTgsMTExMTExMTExMTExMTExMTEsMzU3OzYwLDExMTExMTExMTExMTExMTExLDM1Nzs2MSwxMSwzNTc7NjIsMTEwMDAwMDAwMDAwMDAwMDEwMDAwMDAwLDM1Nzs2MywxMDAxMDAwMTAwMDAwMDAwMDAwMDAwMDAsMzU3OzY0LDAwMDExMTEwMDEwMDAwMDAwMDAxMDAwMCwzNTc7NjYsMTExMTExMTExMTExMTAsMzU3OzY4LDExMTExMTExMTAxMTExMTExMDAwLDM1Nzs3MywxMTExMDAsMzU3Ozc3LDEwMDAwMTAsMzU3Ozc4LDEsMzU3OzkxLDExMTExMTEsMzU3OzkyLDAwMDAwMTAwLDM1NzsxLDEwMSw0MDsyLDEsNDA7MywxLDQwOzEwMiwxLDQwOjk6OjozODMyMzI2MzE6MTYwMDczNzk1Nzo6OjE0ODkwMzgxODA6MjI4ODQzOjA6MWFhMmFiODk1YTE0NmZkN2VmOTI1YTAwZTE5MmFjZTE4Ojow; userid=383232631; u_name=testjrds4; escapename=testjrds4; ticket=df1384d64259c29360f28af895f70e89; user_status=0; laravel_session=hx10jqka480a17e262b5d24cdf07407e122d8ee6; @#!userid!#@=383232631; @#!sessionid!#@=18fd7689bf3dc402fc559af96140b20c7; @#!rsa_version!#@=default_4; v=AjZBpTdsde1rvwGJSJa5jTdvh2c6V3qRzJuu9aAeIpm049zZCOfKoZwr_gFz")
	//req.Header.Set("Cookie", "userid=383232631; u_name=testjrds4;")
	req.Header.Set("Host", "gpc.10jqka.com.cn")
	//req.Header.Set("Referer", "http://gpc.10jqka.com.cn/idea/dataspirit/chart1.html?left=1454&top=168&skin=brown")
	//req.Header.Set("Referer", "http://gpc.10jqka.com.cn/dptrend/fund/MarginTradingInfo/get?startDate=20181125&endDate=20201125")
	req.Header.Set("Referer", "http://ft.10jqka.com.cn/thsft/iFindService/Chain/base/overview?title=%E6%B1%BD%E8%BD%A6&type=qc&sel=survey&selectUrl=home&is_index=1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	str := (*string)(unsafe.Pointer(&body))
	result := unicode2utf8(*str)

	return result
}

func main()  {
	html:=GetHtml()
	fmt.Print(html)
}