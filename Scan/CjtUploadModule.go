package Scan

import (
	"FoPoForYonyou2.0/Utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func YCjtUploadScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	cjtUpload := "用友 畅捷通T+ Upload.aspx 任意文件上传漏洞"
	urls := address + "/tplus/SM/SetupAccount/Upload.aspx?preload=1"
	response, _ := client.Get(urls)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode == http.StatusOK {
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
		if strings.Contains(string(body), "GIF89a") || strings.Contains(string(body), "aliyun") || strings.Contains(string(body), "waf") {
			output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, cjtUpload)
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
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, cjtUpload)
		fmt.Println(output)
	}
}
