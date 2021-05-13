package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

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

// 文件格式 日期_创建时间
func cleanfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "start clean file ... please wait !!")

	curTime := time.Now()

	year := curTime.Year()
	fmt.Print(year)
	month := curTime.Month()
	fmt.Print(month)
	day := curTime.Day()
	fmt.Print(day)

	fmt.Fprint("YYY")

	//fmt.Print(year + month + day)

	// fmt.Print(curTime.Year(), curTi)
	// files = []string{}

	// _dir, err := ioutil.ReadDir(BaseUploadPath)
	// if err != nil {
	// 	return nil, err
	// }

	// suffix = strings.ToLower(suffix) //匹配后缀

	// for _, _file := range _dir {
	// 	if _file.IsDir() {
	// 		continue //忽略目录
	// 	}
	// 	if len(suffix) == 0 || strings.HasSuffix(strings.ToLower(_file.Name()), suffix) {
	// 		//文件后缀匹配
	// 		files = append(files, path.Join(dir, _file.Name()))
	// 	}
	// }

}

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

// func showall(w http.ResponseWriter, r *http.Request) {
// 	// print("list ")
// 	fmt.Fprintln(w, "show file list!")
// 	http.Handle("/", http.FileServer(http.Dir(".")))
// }

func main() {
	// 文件服务器
	http.HandleFunc("/cleanfile", cleanfile)
	http.HandleFunc("/getdiskinfo", getdiskinfo)

	http.Handle("/", http.FileServer(http.Dir(BaseUploadPath)))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Server run fail")
	}
}
