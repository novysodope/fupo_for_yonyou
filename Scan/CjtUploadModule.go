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

func YCjtUploadScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	cjtUpload := "用友 畅捷通T+ Upload.aspx 任意文件上传漏洞"
	urls := address + "/tplus/SM/SetupAccount/Upload.aspx?preload=1"
	response, err := client.Get(urls)
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, cjtUpload)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, cjtUpload)
		}
		return
	}
	defer response.Body.Close()

	//body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Yellow, Reset, cjtUpload, err)
			fmt.Println(output)
			return
		}
		if strings.Contains(string(body), "submitPic") && !strings.Contains(string(body), "GIF89a") {
			result := cjtUpload + "：" + urls
			//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
			timestamp := time.Now().Format("20230712")
			filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
			err := Utils.SaveResultToFile(result, filename)
			if err != nil {
				log.Println("保存结果到文件出错:", err)
			} else {

			}
			output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, cjtUpload, urls)
			fmt.Println(output)
		}
	} else if response.StatusCode == http.StatusInternalServerError || response.StatusCode == http.StatusMethodNotAllowed {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Yellow, Reset, cjtUpload, err)
			fmt.Println(output)
			return
		}
		if strings.Contains(string(body), "GIF89a") || strings.Contains(string(body), "aliyun") || strings.Contains(string(body), "waf") {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, cjtUpload, response.StatusCode)
			fmt.Println(output)
		} else {
			result := cjtUpload + "：" + urls
			//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
			timestamp := time.Now().Format("20230712")
			filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
			err := Utils.SaveResultToFile(result, filename) // 保存结果到文本文件
			if err != nil {
				log.Println("保存结果到文件出错:", err)
			} else {

			}
			output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, cjtUpload, urls)
			fmt.Println(output)
		}
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, cjtUpload, response.StatusCode)
		fmt.Println(output)
	}
}
