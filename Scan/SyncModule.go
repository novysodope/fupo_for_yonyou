package Scan

import (
	"fmt"
	"fupo_for_yonyou/Utils"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

func SyncScan(synhostt string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	const stnde = "用友 sync反序列化漏洞"
	const pocURL = "https://novysodope.github.io/2022/12/09/97/"

	parsedURL, err := url.Parse(synhostt)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%sERROR%s] %s 解析URL失败: %v", Cyan, currentTime, Reset, Red, Reset, stnde, err)
		fmt.Println(output)
		return
	}

	target := net.JoinHostPort(parsedURL.Hostname(), "8821")
	dialer := &net.Dialer{
		Timeout: 5 * time.Second,
	}

	_, err = dialer.Dial("tcp", target)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s", Cyan, currentTime, Reset, Yellow, Reset, stnde)
		fmt.Println(output)
		return
	}

	result := fmt.Sprintf("%s：%s", stnde, pocURL)
	output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Green, currentTime, Reset, Green, Reset, stnde, pocURL)
	fmt.Println(output)

	timestamp := time.Now().Format("20230712")
	filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
	if err := Utils.SaveResultToFile(result, filename); err != nil {
		log.Printf("保存结果到文件出错: %v", err)
	}
}
