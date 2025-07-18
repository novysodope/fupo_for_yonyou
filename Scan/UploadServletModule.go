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

func UploadServletScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	const UploadServlet = "用友 NC UploadServlet反序列化漏洞"
	urls := address + "/servlet/~ic/nc.document.pub.fileSystem.servlet.UploadServlet"

	// 发送 POST 请求
	response, err := client.Post(urls, "multipart/form-data", strings.NewReader("aa"))
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, UploadServlet)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, UploadServlet)
		}
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%sERROR%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Red, Reset, UploadServlet, err)
		fmt.Println(output)
		return
	}

	// 检查漏洞是否存在
	if strings.Contains(string(body), "EOFExcepti") {
		result := UploadServlet + "：" + urls
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Green, currentTime, Reset, Green, Reset, UploadServlet, urls)
		fmt.Println(output)

		// 保存结果到文件
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		if err := Utils.SaveResultToFile(result, filename); err != nil {
			log.Printf("保存结果到文件出错: %v", err)
		}
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, UploadServlet, response.StatusCode)
		fmt.Println(output)
	}
}
