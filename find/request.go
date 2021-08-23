import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"unsafe"
)

// 获取页面请求数据
func GetHtml(id string) string {
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