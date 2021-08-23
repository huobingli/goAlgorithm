package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
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
	//关联关系
	url := "http://ft.10jqka.com.cn/thsft/iFindService/CompanyLibrary/graph/ajax-get-firm-graph?orgid=T000025698&v=v2&position="
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	// testjrds
	req.Header.Set("Cookie", "jgbsessid=fd32b80df811d9e160ad11d5547a69ec; THSFT_USERID=huobingli; u_name=huobingli; userid=377434299; user=Omh1b2JpbmdsaTo6Ojo6Ojo6OjM3NzQzNDI5OToxNjI5NjgxMDg1Ojo6Ojg2NDAwOjoxNDZjMTJlODU4ZmNmYThiOTAyMDExMjhkMTU1NDRmOGU%3D; ticket=6c7a9a6dd2da2e365dfd55cc8181fc72; escapename=huobingli; version=1.10.12.369; securities=0; platform=w; jgblang=cn; searchGuide=sg; log=; newguidever=2; v=AziBRajYrbi8B8E8Mbnf5wVBCe3JoZwp_gVwr3KphHMmjdLRGrFsu04VQDvB")
	req.Header.Set("Host", "ft.10jqka.com.cn")
	req.Header.Set("Referer", "http://ft.10jqka.com.cn/thsft/iFindService/CompanyLibrary/graph/graph-f9?type=FirmGraph&orgid=T000025698")
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

// 获取页面请求数据
func GetHtmlV2(id string) string {
	//关联关系
	// url := "http://ft.10jqka.com.cn/thsft/iFindService/CompanyLibrary/graph/ajax-get-firm-graph?orgid=%s&v=v2&position=" % "T000393218"
	url := fmt.Sprintf("http://ft.10jqka.com.cn/thsft/iFindService/CompanyLibrary/graph/ajax-get-firm-graph?orgid=%s&v=v2&position=", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	// testjrds
	req.Header.Set("Cookie", "jgbsessid=fd32b80df811d9e160ad11d5547a69ec; THSFT_USERID=huobingli; u_name=huobingli; userid=377434299; user=Omh1b2JpbmdsaTo6Ojo6Ojo6OjM3NzQzNDI5OToxNjI5NjgxMDg1Ojo6Ojg2NDAwOjoxNDZjMTJlODU4ZmNmYThiOTAyMDExMjhkMTU1NDRmOGU%3D; ticket=6c7a9a6dd2da2e365dfd55cc8181fc72; escapename=huobingli; version=1.10.12.369; securities=0; platform=w; jgblang=cn; searchGuide=sg; log=; newguidever=2; v=AziBRajYrbi8B8E8Mbnf5wVBCe3JoZwp_gVwr3KphHMmjdLRGrFsu04VQDvB")
	req.Header.Set("Host", "ft.10jqka.com.cn")
	req.Header.Set("Referer", "http://ft.10jqka.com.cn/thsft/iFindService/CompanyLibrary/graph/graph-f9?type=FirmGraph&orgid=T000025698")
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

func main() {

	// html := GetHtml()
	// fmt.Print(html)
	id := "T000025698"
	ul := "http://ft.10jqka.com.cn/thsft/iFindService/CompanyLibrary/graph/ajax-get-firm-graph?orgid=%s&v=v2&position="
	url := fmt.Sprintf(ul, id)
	print(url)

	// 保存文件
}
