package Scan

import (
	"fupo_for_yonyou/Utils"
	"fmt"
	"log"
	"net"
	"net/url"
	"time"
)

func SyncScan(synhostt string, Red string, Green string, Yellow string, Reset string, Cyan string, currentTime string) {
	stnde := "用友 sync反序列化漏洞"
	urls := "https://novysodope.github.io/2022/12/09/97/"
	parsedURL, _ := url.Parse(synhostt)
	//如果传入的地址带端口，那就只取主机名
	synhos := parsedURL.Hostname()
	target := net.JoinHostPort(synhos, "8821")
	dialer := net.Dialer{
		Timeout: 5 * time.Second,
	}
	_, err := dialer.Dial("tcp", target)
	if err != nil {
		//我感觉这里太杂乱了，应该还能再简化
		output := fmt.Sprintf("["+Cyan+"%s"+Reset+"] ["+Yellow+"-"+Reset+"] 不存在%s ", currentTime, stnde)
		fmt.Println(output)
	} else {
		result4 := stnde + "：" + urls
		//保存文件不能用传来的currentTime，得重新定义一个格式的时间戳
		timestamp := time.Now().Format("20230712")
		filename := fmt.Sprintf("scan_result_%s.txt", timestamp)
		err := Utils.SaveResultToFile(result4, filename)
		if err != nil {
			log.Println("保存结果到文件出错:", err)
		} else {

		}
		output := fmt.Sprintf("["+Green+"%s"+Reset+"] ["+Green+"+"+Reset+"] 存在%s：%s", currentTime, stnde, urls)
		fmt.Println(output)
	}
}
