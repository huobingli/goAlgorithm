package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetHtml() string {
	url := "http://gpc.10jqka.com.cn/dptrend/window/textinfo/windindicator"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	req.Header.Set("Cookie", "JSESSIONID=4ED8149EBB89667A589F78088B8282E0; searchGuide=sg; __utma=156575163.232964267.1582016786.1582016786.1582016786.1; __utmz=156575163.1582016786.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); Hm_lvt_78c58f01938e4d85eaf619eae71b4ed1=1592967889,1593653099,1594275403; user=MDp0ZXN0anJkczQ6Ok5vbmU6NTAwOjM5MzIzMjYzMTo0LDEwMDAwMDAwMTAwMTAwMDAwMDExMTExMSwyNDI7NSwxLDQwOzYsMSw0MDs3LDExMTExMTExMTExMCw0MDs4LDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAxMDAwMDEwMDExLDQwOzIwLDEsMjQyOzIyLDEsMjQyOzI2LDEsMjQyOzI3LDEsMjQyOzI5LDEsMjQyOzMzLDAwMTEwMDAwMDAxMSwyNDI7MzYsMTAwMTExMTExMDAwMTExMTEwMTExMTExLDI0Mjs0MSwxMTExMSwyNDI7NDIsMTAwMTAxLDI0Mjs0MywxMDExMTExMTExMTExMTExMTExMTExMTEsMjQyOzQ0LDExMTExLDQwOzQ1LDEwMTExMTExMSwyNDI7NDYsMDExMDExMTExMDAwMDAxMTExMTExMTExLDI0Mjs0NywxMTExMTExMTExMTExMTEsMjQyOzQ4LDExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExLDI0Mjs1MCwxMSwyNDI7NTEsMTEwMDAwMDAwMDAwMDAwMCwyNDI7NTIsMDEwMCwyNDI7NTQsMTAwMDAwMDAwMDAwMDAwMDAwMDAwMDEwMDAwMDAwMDAsMjQyOzU2LDExLDI0Mjs1NywxMTExMTExMTExMTExMTExMTExMTExMTExMTExLDI0Mjs1OCwxMTExMTExMTExMTExMTExMSwyNDI7NjAsMTExMTExMTExMTExMTExMTEsMjQyOzYxLDExLDI0Mjs2MiwxMTAwMDAwMDAwMDAwMDAwMTAwMDAwMDAsMjQyOzYzLDEwMDEwMDAxMDAwMDAwMDAwMDAwMDAwMCwyNDI7NjQsMDAwMTExMTAwMTAwMDAwMDAwMDEwMDAwLDI0Mjs2NiwxMTExMTExMTExMTExMCwyNDI7NjgsMTExMTExMTExMDExMTExMTEwMDAsMjQyOzczLDExMTEwMCwyNDI7NzcsMTAwMDAxMCwyNDI7NzgsMSwyNDI7OTEsMTExMTExMSwyNDI7OTIsMDAwMDAxMDAsMjQyOzEsMTAxLDQwOzIsMSw0MDszLDEsNDA6MjU6OjozODMyMzI2MzE6MTU5NDg2MjEzMTo6OjE0ODkwMzgxODA6NDAyMjY5OjA6MTNhODc0NDkxYWU5ODMwOGFhZTkyNDU5M2I0NDljYjYwOjow; userid=383232631; u_name=testjrds4; escapename=testjrds4; ticket=edb3851a2cea6aa0f3d23b2d9c078652; @#!userid!#@=383232631; @#!sessionid!#@=13238245a8a8645c18bb63571f0560c45; @#!rsa_version!#@=default_4; v=Agx7I7Gev05JJasyPUEjI9nd3WE8RbDvsunEs2bNGLda8aZfjlWAfwL5lEK1")
	req.Header.Set("Host", "gpc.10jqka.com.cn")
	req.Header.Set("Referer", "http://gpc.10jqka.com.cn/idea/dataspirit/chart1.html?left=1454&top=168&skin=brown")
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
	return string(body)
}

func main()  {
	html:=GetHtml()
	fmt.Print(html)
}