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

func NCMessageServletScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	NCMessageServlet := "用友 NC MessageServlet反序列化漏洞"
	urls := address + "/servlet/~baseapp/nc.message.bs.NCMessageServlet"

	response, err := client.Post(urls, "multipart/form-data", strings.NewReader("aa"))
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, NCMessageServlet)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, NCMessageServlet)
		}
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应失败: %v", Cyan, currentTime, Reset, Yellow, Reset, NCMessageServlet, err)
		fmt.Println(output)
		return
	}

	if strings.Contains(string(body), "EOFExcepti") {
		result := NCMessageServlet + "：" + urls
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Green, currentTime, Reset, Green, Reset, NCMessageServlet, urls)
		fmt.Println(output)

		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		if err := Utils.SaveResultToFile(result, filename); err != nil {
			log.Printf("保存结果到文件出错: %v", err)
		}
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, NCMessageServlet, response.StatusCode)
		fmt.Println(output)
	}
}
