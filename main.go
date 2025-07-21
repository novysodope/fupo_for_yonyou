package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

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

func showStartupScreen() {
	fmt.Println(flicker + "\n███████╗██╗   ██╗██████╗  ██████╗ ")
	fmt.Println("██╔════╝██║   ██║██╔══██╗██╔═══██╗")
	fmt.Println("█████╗  ██║   ██║██████╔╝██║   ██║")
	fmt.Println("██╔══╝  ██║   ██║██╔═══╝ ██║   ██║")
	fmt.Println("██║     ╚██████╔╝██║     ╚██████╔╝")
	fmt.Println("╚═╝      ╚═════╝ ╚═╝      ╚═════╝ ")
	fmt.Println("                  " + Yellow + "用友全系列漏洞检测v3.1  Fupo's series" + Reset)
	fmt.Println("—————————————————————————————————————————————————————")
	fmt.Println(Reset)

}

func main() {
	showStartupScreen()

	if len(os.Args) < 2 {
		fmt.Println("[" + Red + "ERROR" + Reset + "] -h 查看使用帮助")
		return
	}

	var (
		address    string
		socks5Addr string
		httpProxy  string
		filePath   string
	)

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-h":
			helpArg()
			return

		case "-u":
			if i+1 < len(os.Args) {
				address = strings.TrimSuffix(os.Args[i+1], "/")
				fmt.Printf("["+Yellow+"INFO"+Reset+"] %s\n", address)
				i++
			}

		case "-socks5":
			if i+1 < len(os.Args) {
				raw := os.Args[i+1]
				if u, err := url.Parse(raw); err == nil && u.Host != "" {
					socks5Addr = u.Host
				} else {
					socks5Addr = raw
				}
				fmt.Printf("["+Green+"SOCKS5"+Reset+"] %s\n", socks5Addr)
				i++
			}

		case "-proxy":
			if i+1 < len(os.Args) {
				raw := os.Args[i+1]
				u, err := url.Parse(raw)
				if err != nil || u.Host == "" {
					fmt.Println("HTTP 代理地址格式错误，应为 http://host:port")
					os.Exit(1)
				}
				httpProxy = u.String()
				fmt.Printf("["+Blue+"HTTP_PROXY"+Reset+"] %s\n", httpProxy)
				i++
			}

		case "-f":
			if i+1 < len(os.Args) {
				filePath = os.Args[i+1]
				i++
			}
		}
	}

	if socks5Addr != "" && httpProxy != "" {
		fmt.Println("[" + Magenta + "WARN" + Reset + "] 同时指定了 HTTP 代理 和 SOCKS5 代理，将优先使用 HTTP 代理")
	}

	if address != "" {
		TargetParse(address, httpProxy, socks5Addr, Red, Green, Yellow, Reset, Cyan)
	}

	if filePath != "" {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("打开文件失败:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		count := 0
		for scanner.Scan() {
			url := strings.TrimSpace(scanner.Text())
			if url == "" {
				continue
			}
			url = strings.TrimSuffix(url, "/")
			count++
			fmt.Printf("["+Yellow+"INFO"+Reset+"] %s\n", url)
			TargetParse(address, httpProxy, socks5Addr, Red, Green, Yellow, Reset, Cyan)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("读取文件出错:", err)
		} else {
			fmt.Printf("初始化完成，共 %d 条 URL\n", count)
		}
	}
}

func helpArg() {
	fmt.Println("单条检测：fupo_for_yonyou -u http[s]://1.1.1.1/")
	fmt.Println("批量检测：fupo_for_yonyou -f url.txt")
	fmt.Println("SOCKS5：-socks5 socks5://0.0.0.0:1080")
	fmt.Println("HTTP代理：-proxy http://127.0.0.1:8080\n")

	fmt.Println(Green + "目前支持的漏洞检测：\n" + Reset)
	fmt.Println(Yellow + "用友 NC bsh.servlet.BshServlet 远程命令执行漏洞")
	fmt.Println("用友 U8 OA getSessionList.jsp 敏感信息泄漏漏洞")
	fmt.Println("用友 FE协作办公平台 templateOfTaohong_manager目录遍历漏洞")
	fmt.Println("用友 NCFindWeb 任意文件读取漏洞")
	fmt.Println("用友 GRP-U8 UploadFileData 任意文件上传漏洞")
	fmt.Println("用友 GRP-U8 Proxy SQL注入漏洞")
	fmt.Println("用友 U8 OA test.jsp SQL注入漏洞")
	fmt.Println("用友 Uapjs 远程代码执行漏洞")
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
	fmt.Println("用友 NC FileReceiveServlet反序列化漏洞")
	fmt.Println("用友 NC MessageServlet反序列化漏洞")
	fmt.Println("用友 NC UploadServlet反序列化漏洞")
	fmt.Println("用友 NC MonitorServlet反序列化漏洞")
	fmt.Println("用友 NC service 接口信息泄露漏洞")

	//3.0优化
	fmt.Println("用友 NC IUpdateService XXE漏洞")
	fmt.Println("用友 NC cloud accept任意文件上传漏洞")

	//3.0新增
	fmt.Println("用友 U8 cloud KeyWordDetailReportQuery SQL注入漏洞")
	fmt.Println("用友 GRP-U8 ufgovbank XXE漏洞")
	fmt.Println("用友 移动系统管理 uploadIcon.do 任意文件上传漏洞")
	fmt.Println("⽤友 NC-word.docx 文件读取漏洞")
	fmt.Println("用友 portal 任意文件读取漏洞")
	fmt.Println("用友 U8+ CRM help2.php 任意文件读取漏洞")
	fmt.Println("用友 U8+ CRM getemaildata.php 任意文件读取漏洞")
	fmt.Println("用友 GRP U8 license_check.jsp SQL注入漏洞")
	fmt.Println("用友 GRP U8 SelectDMJE.jsp SQL注入漏洞")
	fmt.Println("用友 GRP-U8 bx_historyDataCheck.jsp SQL注入漏洞")
	fmt.Println("用友 U8 Cloud smartweb2.RPC.d XXE漏洞")
	fmt.Println("用友 U8 obr_zdybxd_check SQL注入漏洞")
	fmt.Println("用友 畅捷通Tplus GetStoreWarehouseByStore 远程代码执行漏洞")
	fmt.Println("用友 畅捷通Tplus CheckMutex SQL注入漏洞")
	fmt.Println("用友 畅捷通T+ GetDecAllUsers 信息泄露漏洞")
	fmt.Println("用友 畅捷通远程通 GNRemote.dll SQL注入漏洞")

	//3.1新增
	fmt.Println("用友 U9 DynamaticExport.aspx 任意文件读取漏洞")
	fmt.Println("用友 U8Cloud FilterCondAction SQL注入漏洞")
	fmt.Println("用友 NC-Cloud process SQL注入漏洞")
	fmt.Println("用友 NC FormItemServlet SQL注入漏洞")
	fmt.Println("用友 NC-Cloud blobRefClassSearch FastJson反序列化漏洞")
	fmt.Println("用友 U8cloud esnserver 远程代码执行漏洞")
	fmt.Println("用友 U8-CRM ajaxgetborrowdata.php SQL注入漏洞")
	fmt.Println("用友 NC Cloud queryPsnInfo SQL注入漏洞")
	fmt.Println("用友 U8 Cloud MeasureQResultAction SQL注入漏洞")
	fmt.Println("用友 NC yerfile/down SQL注入漏洞")
	fmt.Println("用友 NC checkekey SQL 注入漏洞")
	fmt.Println("用友 NC content SQL注入漏洞")
	fmt.Println("用友 U8-CRM-reservationcomplete SQL注入漏洞")
	fmt.Println("用友 U8-CRM-reservationcomplete 身份认证绕过漏洞")
	fmt.Println("用友 U8 Cloud approveservlet SQL注入漏洞")
	fmt.Println("用友 YonBIP yonbiplogin 任意文件读取漏洞")
	fmt.Println("用友 U8 Cloud MultiRepChooseAction SQL注入漏洞")
	fmt.Println("用友 NC IMetaWebService4BqCloud SQL注入漏洞")

	fmt.Println(Reset)
}
