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

func MonitorServletScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	MonitorServlet := "用友 NC MonitorServlet反序列化漏洞"
	urls := address + "/servlet/~ic/nc.bs.framework.mx.monitor.MonitorServlet"
	response, _ := client.Post(urls, "multipart/form-data;", strings.NewReader("aa"))
	defer response.Body.Close()
	response2, _ := client.Get(urls)
	defer response2.Body.Close()

	// 读取响应内容
	body21, _ := ioutil.ReadAll(response.Body)
	body22, _ := ioutil.ReadAll(response.Body)
	if strings.Contains(string(body21), "EOFException") || strings.Contains(string(body22), "EOFException") && response2.StatusCode == http.StatusInternalServerError {
		result18 := MonitorServlet + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result18, filename) // 保存结果到文本文件
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, MonitorServlet, urls)
		fmt.Println(output)
		output2 := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在/~xx/系列接口（参考https://mp.weixin.qq.com/s/xVKuJb3DbKH0em0HoMZ4ZQ），请自行尝试", currentTime)
		fmt.Println(output2)
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, MonitorServlet)
		fmt.Println(output)
	} // time.Sleep(1 * time.Second)
}
