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

func ServiceDispatcherServletScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	ServiceDispatcherServlet := "用友 ServiceDispatcherServlet 反序列化漏洞"
	urls := address + "/ServiceDispatcherServlet"
	response, _ := client.Get(urls)
	defer response.Body.Close()
	bodywaf, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode == http.StatusInternalServerError {
		if strings.Contains(string(bodywaf), "aliyun") || strings.Contains(string(bodywaf), "waf") || !strings.Contains(string(bodywaf), "ServletException") {
			output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, ServiceDispatcherServlet)
			fmt.Println(output)
		} else {
			result13 := ServiceDispatcherServlet + "：" + urls
			//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
			timestamp := time.Now().Format("20230712")
			filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
			err := Utils.SaveResultToFile(result13, filename) // 保存结果到文本文件
			if err != nil {
				log.Println("保存结果到文件出错:", err)
			} else {

			}
			output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, ServiceDispatcherServlet, urls)
			fmt.Println(output)
		}
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, ServiceDispatcherServlet)
		fmt.Println(output)
	}
}
