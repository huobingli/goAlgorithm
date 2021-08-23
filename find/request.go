import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"unsafe"
)

// 获取页面请求数据
func GetHtml(request_url, request_id, request_cookie string) string {
	url := fmt.Sprintf(request_url, request_id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	// testjrds
	req.Header.Set("Cookie", request_cookie)
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