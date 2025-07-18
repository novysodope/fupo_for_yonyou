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

func RecoverPasswordScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	RecoverPassword := "用友 畅捷通T+ RecoverPassword.aspx 管理员密码修改漏洞"
	urls := address + "/tplus/ajaxpro/RecoverPassword,App_Web_recoverpassword.aspx.cdcab7d2.ashx"
	response, err := client.Get(urls)
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, RecoverPassword)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, RecoverPassword)
		}
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应失败: %v", Cyan, currentTime, Reset, Yellow, Reset, RecoverPassword, err)
			fmt.Println(output)
			return
		}
		if strings.Contains(string(body), "pwdNew") {
			result := RecoverPassword + "：" + urls
			timestamp := time.Now().Format("20230712")
			filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
			err := Utils.SaveResultToFile(result, filename)
			if err != nil {
				log.Println("保存结果到文件出错:", err)
			}
			output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Cyan, currentTime, Reset, Green, Reset, RecoverPassword, urls)
			fmt.Println(output)
		} else {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, RecoverPassword, response.StatusCode)
			fmt.Println(output)
		}
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, RecoverPassword, response.StatusCode)
		fmt.Println(output)
	}
}
