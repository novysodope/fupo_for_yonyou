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

func FileReceiveServletcan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	urls := address + "/servlet/FileReceiveServlet"
	FileReceiveServlet := "用友 FileReceiveServlet反序列化漏洞"

	response, _ := client.Post(urls, "multipart/form-data;", strings.NewReader("ok"))
	defer response.Body.Close()
	// 读取响应内容
	body, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode == http.StatusOK && len(body) == 0 {
		result := FileReceiveServlet + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result, filename) // 保存结果到文本文件
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, FileReceiveServlet, urls)
		fmt.Println(output)
	} else if response.StatusCode == http.StatusInternalServerError {
		result := FileReceiveServlet + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result, filename) // 保存结果到文本文件
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, FileReceiveServlet, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, FileReceiveServlet)
		fmt.Println(output)
	}
}
