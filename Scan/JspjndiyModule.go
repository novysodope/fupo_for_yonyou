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

func JspjndiyScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	jspjndiy := "用友 testperformance JNDI注入漏洞"
	urls := address + "/testperformance.jsp"
	response, _ := client.Get(urls)
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body2, _ := ioutil.ReadAll(response.Body)
		if strings.Contains(string(body2), "compute") {
			result15 := jspjndiy + ":" + urls
			//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
			timestamp := time.Now().Format("20230712")
			filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
			err := Utils.SaveResultToFile(result15, filename) // 保存结果到文本文件
			if err != nil {
				log.Println("保存结果到文件出错:", err)
			} else {

			}
			output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, jspjndiy, urls)
			fmt.Println(output)
		} else {
			output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, jspjndiy)
			fmt.Println(output)
		}
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, jspjndiy)
		fmt.Println(output)
	}
}
