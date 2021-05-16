package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
)

const BaseUploadPath = "C:\\ci_auto_publish\\"

func TimeParseYYYYMMDD(in string, sub string) (out time.Time, err error) {
	layout := "2006" + sub + "01" + sub + "02"
	out, err = time.ParseInLocation(layout, in, time.Local)
	if err != nil {
		return
	}
	return
}

func getCurDay() (date int) {
	curTime := time.Now()
	year := curTime.Year()
	month := int(curTime.Month())
	day := curTime.Day()
	return year*10000 + month*100 + day
}

// 清理目录中一个月前的的临时文件  文件格式 日期_创建时间
func cleanfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "start clean file ... please wait !!")

	curDay := getCurDay()

	//files = []string{}

	_dir, err := ioutil.ReadDir(BaseUploadPath)
	if err != nil {
		return
	}

	for _, _file := range _dir {
		if _file.IsDir() {
			dirname := _file.Name()

			// 根据“_”分割文件名
			comm := strings.Index(dirname, "_")

			strDirDate := dirname[:comm]

			dirDate, err := strconv.Atoi(strDirDate)
			if err != nil {
				fmt.Print("strconv.Atoi, err:%v\n", err)
			}

			if dirDate < curDay-100 {
				fmt.Println(dirDate)
				removeDir := BaseUploadPath + dirname
				os.RemoveAll(removeDir)
			}
		}
	}

}

func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	print(percent[0])
	return percent[0]
}

// 后面想获取磁盘信息，预留一下口子与参考blog
// https://blog.csdn.net/whatday/article/details/109620192
func getdiskinfo(w http.ResponseWriter, r *http.Request) {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Print("get Partitions failed, err:%v\n", err)
		return
	}

	for _, part := range parts {
		fmt.Print("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Print("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}

	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Print("%v:%v\n", k, v)
	}

	// fmt.Fprintln(disk_info)
}

func main() {
	// 其他接口
	http.HandleFunc("/cleanfile", cleanfile)
	http.HandleFunc("/getdiskinfo", getdiskinfo)

	// 文件服务器
	http.Handle("/", http.FileServer(http.Dir(BaseUploadPath)))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Server run fail")
	}
}
