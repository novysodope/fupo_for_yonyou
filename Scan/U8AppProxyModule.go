package Scan

import (
	"fmt"
	"fupo_for_yonyou/Utils"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func U8AppProxyScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	U8AppProxy := "用友 GRP-U8 U8AppProxy 任意文件上传漏洞"
	urls := address + "/U8AppProxy?gnid=myinfo&id=saveheader&zydm=../../yongyouU8_test"
	response, err := client.Get(urls)
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, U8AppProxy)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, U8AppProxy)
		}
		return
	}
	defer response.Body.Close()

	bodywaf2, err := io.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%sERROR%s] %s 读取响应体失败: %v", Cyan, currentTime, Reset, Red, Reset, U8AppProxy, err)
		fmt.Println(output)
		return
	}

	if (response.StatusCode == http.StatusInternalServerError || response.StatusCode == http.StatusMethodNotAllowed || response.StatusCode == http.StatusUnsupportedMediaType) && !strings.Contains(string(bodywaf2), "GIF89a") && !strings.Contains(string(bodywaf2), "aliyun") && !strings.Contains(string(bodywaf2), "waf") {
		result14 := U8AppProxy + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result14, filename)
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Green, currentTime, Reset, Green, Reset, U8AppProxy, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, U8AppProxy, response.StatusCode)
		fmt.Println(output)
	}
}
