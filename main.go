package main

import (
	"bufio"
	"fmt"
	"io"
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
	fmt.Println("                  " + Yellow + "用友全系列漏洞检测v3  Fupo's series" + Reset)
	fmt.Println("—————————————————————————————————————————————————————")
	fmt.Println(Reset)

}

func main() {
	showStartupScreen()
	if len(os.Args) < 2 {
		output := fmt.Sprintf("[" + Red + "ERROR" + Reset + "] -h查看使用帮助\n")
		fmt.Println(output)
		return
	}
	//取出控制台参数，判断参数
	helpArgFunc := false
	for _, arg := range os.Args {
		if arg == "-h" {
			helpArgFunc = true
			break
		}
	}
	//根据参数进入对应的方法
	if helpArgFunc {
		helpArg()
		return
	}

	//扫描参数，从第二个参数开始获取
	args := os.Args[1:]
	argBatch := os.Args[1:]
	argsocks := os.Args[1:]
	var address string
	var filePath string
	var socks5 string

	//socks5代理
	for i := 0; i < len(argsocks); i++ {
		if argsocks[i] == "-socks5" && i+1 < len(argsocks) {
			socks5 = argsocks[i+1]
			if strings.HasPrefix(socks5, "socks5://") {
				socks5 = socks5[len("socks5://"):]
			}
			output := fmt.Sprintf("["+Green+"SOCKS5"+Reset+"] %s ", socks5)
			fmt.Println(output)
		}
	}

	//单个扫描
	for i := 0; i < len(args); i++ {
		if args[i] == "-u" && i+1 < len(args) {
			address = args[i+1]
			output := fmt.Sprintf("["+Yellow+"INFO"+Reset+"] %s ", address)
			fmt.Println(output)
		}
	}
	if address != "" {
		address = strings.TrimSuffix(address, "/")

		TargetParse(address, socks5, Red, Green, Yellow, Reset, Cyan)
	}

	//批量扫描
	//https://github.com/novysodope/fupo_for_yonyou/issues/2
	for i := 0; i < len(argBatch); i++ {
		if argBatch[i] == "-f" && i+1 < len(argBatch) {

			filePath = argBatch[i+1]
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Println("打开文件失败:", err)
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			urlCount := 0
			for scanner.Scan() {
				urls := strings.TrimSpace(scanner.Text())
				if urls == "" {
					continue // 跳过空行
				}
				urlCount++
				if err := scanner.Err(); err != nil {
					if err == io.EOF {
						// 文件已到达末尾，正常结束
						fmt.Printf("初始化完成，一共有 %d 条 URL\n", urlCount)
					} else {
						fmt.Println("读取文件出错:", err)
					}
				}
				urls = strings.TrimSuffix(urls, "/")
				output := fmt.Sprintf("["+Yellow+"INFO"+Reset+"] %s ", urls)
				fmt.Println(output)
				TargetParse(urls, socks5, Red, Green, Yellow, Reset, Cyan)
			}
		}
	}

}

func helpArg() {
	fmt.Println("单条检测：fupo_for_yonyou -u http[s]://1.1.1.1/")
	fmt.Println("批量检测：fupo_for_yonyou -f url.txt")
	fmt.Println("SOCKS5：-socks5 socks5://0.0.0.0:1080\n")

	fmt.Println(Green + "目前工具内部支持的漏洞检测：\n" + Reset)
	fmt.Println(Yellow + "用友 NC bsh.servlet.BshServlet 远程命令执行漏洞")
	fmt.Println("用友 U8 OA getSessionList.jsp 敏感信息泄漏漏洞")
	fmt.Println("用友 FE协作办公平台 templateOfTaohong_manager目录遍历漏洞")
	fmt.Println("用友 NCFindWeb 任意文件读取漏洞")
	fmt.Println("用友 GRP-U8 UploadFileData 任意文件上传漏洞")
	fmt.Println("用友 GRP-U8 Proxy SQL注入")
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
	fmt.Println("用友 NC cloud accept任意文件上传漏洞")
	fmt.Println("用友 NC MessageServlet反序列化漏洞")
	fmt.Println("用友 NC UploadServlet反序列化漏洞")
	fmt.Println("用友 NC MonitorServlet反序列化漏洞")
	fmt.Println("用友 NC service 接口信息泄露漏洞")
	fmt.Println("用友 NC IUpdateService XXE漏洞漏洞")
	fmt.Println(Reset)

}
