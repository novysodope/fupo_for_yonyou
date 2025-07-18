package Scan

import (
	"fmt"
	"fupo_for_yonyou/Utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func IUpdateServicexxeScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	const IUpdateServicexxe = "用友 NC IUpdateService XXE漏洞"
	urls := address + "/uapws/service/nc.uap.oba.update.IUpdateService?xsd="

	// 发送 GET 请求
	response, err := client.Get(urls)
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, IUpdateServicexxe)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, IUpdateServicexxe)
		}
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Yellow, Reset, IUpdateServicexxe, err)
		fmt.Println(output)
		return
	}

	// 检查漏洞是否存在
	if strings.Contains(string(body), "getDocument") {
		result := IUpdateServicexxe + "：" + urls
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Cyan, currentTime, Reset, Green, Reset, IUpdateServicexxe, urls)
		fmt.Println(output)

		// 保存结果到文件
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		if err := Utils.SaveResultToFile(result, filename); err != nil {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 保存结果到文件出错: %v", Cyan, currentTime, Reset, Yellow, Reset, IUpdateServicexxe, err)
			fmt.Println(output)
		}
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, IUpdateServicexxe, response.StatusCode)
		fmt.Println(output)
	}
}
