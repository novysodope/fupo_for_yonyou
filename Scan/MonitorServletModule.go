package Scan

import (
	"fmt"
	"fupo_for_yonyou/Utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func MonitorServletScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	const MonitorServlet = "用友 NC MonitorServlet反序列化漏洞"
	urls := address + "/servlet/~ic/nc.bs.framework.mx.monitor.MonitorServlet"

	// 发送 POST 请求
	response, err := client.Post(urls, "multipart/form-data", strings.NewReader("aa"))
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, MonitorServlet)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, MonitorServlet)
		}
		return
	}
	defer response.Body.Close()

	// 发送 GET 请求
	response2, err := client.Get(urls)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 请求失败: %v", Cyan, currentTime, Reset, Yellow, Reset, MonitorServlet, err)
		fmt.Println(output)
		return
	}
	defer response2.Body.Close()

	// 读取 POST 响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Yellow, Reset, MonitorServlet, err)
		fmt.Println(output)
		return
	}

	// 读取 GET 响应内容
	body2, err := ioutil.ReadAll(response2.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Yellow, Reset, MonitorServlet, err)
		fmt.Println(output)
		return
	}

	// 检查漏洞是否存在
	if strings.Contains(string(body), "EOFException") || (strings.Contains(string(body2), "EOFException")) {
		result := MonitorServlet + "：" + urls
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Cyan, currentTime, Reset, Green, Reset, MonitorServlet, urls)
		fmt.Println(output)
		output2 := fmt.Sprintf("[%s%s%s] [%s+%s] 存在/~xx/系列接口（参考https://mp.weixin.qq.com/s/xVKuJb3DbKH0em0HoMZ4ZQ），请自行尝试", Cyan, currentTime, Reset, Green, Reset)
		fmt.Println(output2)

		// 保存结果到文件
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		if err := Utils.SaveResultToFile(result, filename); err != nil {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 保存结果到文件出错: %v", Cyan, currentTime, Reset, Yellow, Reset, MonitorServlet, err)
			fmt.Println(output)
		}
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, MonitorServlet, response.StatusCode)
		fmt.Println(output)
	}
}
