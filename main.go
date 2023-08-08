package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"golang.org/x/net/proxy"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"
)

// 保存扫描结果
func saveResultToFile(result, filename string) error {
	// 打开文件以追加模式
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将结果追加写入文件
	_, err = file.WriteString(result + "\n\n")
	if err != nil {
		return err
	}

	return nil
}

// 定义颜色
const (
	Black   = "\u001b[30m"
	Red     = "\u001b[31m"
	Green   = "\u001b[32m"
	Yellow  = "\u001b[33m"
	Blue    = "\u001b[34m"
	Magenta = "\u001b[35m"
	Cyan    = "\u001b[36m"
	White   = "\u001b[37m"
	Reset   = "\u001b[0m"
	flicker = "\u001b[5m"
)
//ui
func showStartupScreen() {
	fmt.Println(flicker + "\n███████╗██╗   ██╗██████╗  ██████╗ ")
	fmt.Println("██╔════╝██║   ██║██╔══██╗██╔═══██╗")
	fmt.Println("█████╗  ██║   ██║██████╔╝██║   ██║")
	fmt.Println("██╔══╝  ██║   ██║██╔═══╝ ██║   ██║")
	fmt.Println("██║     ╚██████╔╝██║     ╚██████╔╝")
	fmt.Println("╚═╝      ╚═════╝ ╚═╝      ╚═════╝ ")
	//fmt.Println(ascii)
	fmt.Println("                  " + Yellow + "用友全系列漏洞检测v2.0RC1  Fupo's series" + Reset)
	fmt.Println("——————————————————————————————————————————————————————————")
	//fmt.Println("*********************************************\n")
	fmt.Println(Reset)

}
func helpArg() {
	fmt.Println("单条检测：fupo_for_yonyou -u http[s]://1.1.1.1/")
	fmt.Println("批量检测：fupo_for_yonyou -f url.txt")
	fmt.Println("SOCKS5：-socks5 socks5://0.0.0.0:1080\n\n")
	//fmt.Println("./fupo_for_yonyou -socks5 socks5://xxx.xxx.xxx:xxx OR  xxx.xxxx.xxx:xxx\n\n")
	// log.Println("漏洞利用：./fupo_for_yonyou -c http[s]://1.1.1.1/ （暂未实现）\n\n")
	fmt.Println(Yellow + "目前支持的漏洞检测：\n")
	fmt.Println("用友 NC bsh.servlet.BshServlet 远程命令执行漏洞")
	fmt.Println("用友 U8 OA getSessionList.jsp 敏感信息泄漏漏洞")
	fmt.Println("用友 FE协作办公平台 templateOfTaohong_manager目录遍历漏洞")
	fmt.Println("用友 NCFindWeb 任意文件读取漏洞")
	fmt.Println("用友 GRP-U8 UploadFileData 任意文件上传漏洞")
	fmt.Println("用友 GRP-U8 Proxy SQL注入")
	fmt.Println("用友 U8 OA test.jsp SQL注入漏洞")
	fmt.Println("用友 Uapjs JNDI注入漏洞")
	fmt.Println("用友 畅捷通T-CRM get_usedspace.php SQL注入漏洞")
	fmt.Println("用友 畅捷通T+ Upload.aspx 任意文件上传漏洞")
	fmt.Println("用友 畅捷通T+ RecoverPassword.aspx 管理员密码修改漏洞")
	fmt.Println("用友 ServiceDispatcherServlet 反序列化漏洞")
	fmt.Println("用友 时空KSOA com.sksoft.bill.ImageUpload 任意文件上传漏洞")
	fmt.Println("用友 GRP-U8 U8AppProxy 任意文件上传漏洞")
	fmt.Println("用友 某jsp JNDI注入漏洞 一")
	fmt.Println("用友 某jsp JNDI注入漏洞 二")
	fmt.Println("用友 sync 反序列化漏洞")
	fmt.Println("用友 uapws 认证绕过漏洞")
	fmt.Println("用友 ajax JNDI注入漏洞")
	fmt.Println("用友 文件服务器 认证绕过漏洞")
	fmt.Println("用友 文件服务器 未授权访问漏洞")
	fmt.Println("用友 files 反序列化漏洞")
	fmt.Println("用友 文件服务器 反序列化漏洞")
	fmt.Println("用友 畅捷通T+ DownloadProxy任意文件读取漏洞")
	//log.Println("用友 GRP U8 XXNode SQL注入漏洞 -未实现")
	//log.Println("用友 GRP U8 forgetPassword_old.jsp SQL注入漏洞 -未实现")
	//log.Println("用友 畅捷通 ajaxpro 反序列化漏洞 -未实现")
	//log.Println("用友 畅捷通 Controller SQL注入漏洞 -未实现")
	fmt.Println("用友 NC FileReceiveServlet反序列化漏洞")
	fmt.Println("用友 NC cloud accept任意文件上传漏洞")
	fmt.Println("用友 NC MessageServlet反序列化漏洞")
	fmt.Println("用友 NC UploadServlet反序列化漏洞")
	fmt.Println("用友 NC MonitorServlet反序列化漏洞")
	fmt.Println("用友 NC service 接口信息泄露漏洞").Println("用友 NC IUpdateService XXE漏洞漏洞")
	fmt.Println(Reset)

}

// 漏洞检测
func scanPoc(address string, passArgb string, passArgu string, passArgy string, proxyAddr string) {

  }


func main() {
  showStartupScreen()
  fmt.println()
}
