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

func UapwsloginScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	uapwslogin := "用友 uapws 认证绕过漏洞"
	urls := address + "/uapws/login.ajax"
	response, err := client.Get(urls)
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, uapwslogin)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, uapwslogin)
		}
		return
	}
	defer response.Body.Close()
	body4, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Red+"x"+Reset+"] %s 读取响应失败: %s", currentTime, uapwslogin, err)
		fmt.Println(output)
		return
	}
	if len(body4) == 1 && strings.Contains(string(body4), "0") {
		result17 := uapwslogin + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result17, filename) // 保存结果到文本文件
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, uapwslogin, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, uapwslogin, response.StatusCode)
		fmt.Println(output)
	}
}
