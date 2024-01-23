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

func ServiceinforScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	serviceinfor := "用友 NC service 接口信息泄露"
	urls := address + "/uapws/service"
	response, _ := client.Get(urls)
	defer response.Body.Close()

	// 读取响应内容
	body23, _ := ioutil.ReadAll(response.Body)

	if strings.Contains(string(body23), "{http") {
		result18 := serviceinfor + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result18, filename) // 保存结果到文本文件
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, serviceinfor, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, serviceinfor)
		fmt.Println(output)
	}
	fmt.Println(Reset)
}
