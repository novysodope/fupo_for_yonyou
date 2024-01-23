# 遵纪守法
## 本工具仅用于企业资产隐患自查，请勿用于违法行为
**fupo_for_yonyou 3.0**
- 开源
- 重构代码
- 更新漏洞检测
- 其他

## 编译
通过下载源码编译时，注意`Scan`模块的头部引用，需要把`FoPoForYonyou2.0`改为`fupo_for_yonyou`

### 跨平台
可以通过`go env -w [variable]`使用交叉编译

### 打包
```bash
go install
```

编译好的文件在`$GOPATH/bin`下，可以通过执行`go env`查看文件夹位置，也可以到[releases](https://github.com/novysodope/fupo_for_yonyou/releases)下载已编译好的文件

## 使用
```bash
C:\Users\Administrator\go\bin>FoPoForYonyou.exe -h

███████╗██╗   ██╗██████╗  ██████╗
██╔════╝██║   ██║██╔══██╗██╔═══██╗
█████╗  ██║   ██║██████╔╝██║   ██║
██╔══╝  ██║   ██║██╔═══╝ ██║   ██║
██║     ╚██████╔╝██║     ╚██████╔╝
╚═╝      ╚═════╝ ╚═╝      ╚═════╝
                  用友全系列漏洞检测v3  Fupo's series
—————————————————————————————————————————————————————

单条检测：fupo_for_yonyou -u http[s]://1.1.1.1/
批量检测：fupo_for_yonyou -f url.txt
SOCKS5：-socks5 socks5://0.0.0.0:1080

目前支持的漏洞检测：

用友 NC bsh.servlet.BshServlet 远程命令执行漏洞
用友 U8 OA getSessionList.jsp 敏感信息泄漏漏洞
用友 FE协作办公平台 templateOfTaohong_manager目录遍历漏洞
用友 NCFindWeb 任意文件读取漏洞
用友 GRP-U8 UploadFileData 任意文件上传漏洞
用友 GRP-U8 Proxy SQL注入
用友 U8 OA test.jsp SQL注入漏洞
用友 Uapjs 远程代码执行漏洞
用友 畅捷通T-CRM get_usedspace.php SQL注入漏洞
用友 畅捷通T+ Upload.aspx 任意文件上传漏洞
用友 畅捷通T+ RecoverPassword.aspx 管理员密码修改漏洞
用友 ServiceDispatcherServlet 反序列化漏洞
用友 时空KSOA com.sksoft.bill.ImageUpload 任意文件上传漏洞
用友 GRP-U8 U8AppProxy 任意文件上传漏洞
用友 某jsp JNDI注入漏洞 一
用友 某jsp JNDI注入漏洞 二
用友 sync 反序列化漏洞
用友 uapws 认证绕过漏洞
用友 ajax JNDI注入漏洞
用友 文件服务器 认证绕过漏洞
用友 文件服务器 未授权访问漏洞
用友 files 反序列化漏洞
用友 文件服务器 反序列化漏洞
用友 畅捷通T+ DownloadProxy任意文件读取漏洞
用友 NC FileReceiveServlet反序列化漏洞
用友 NC MessageServlet反序列化漏洞
用友 NC UploadServlet反序列化漏洞
用友 NC MonitorServlet反序列化漏洞
用友 NC service 接口信息泄露漏洞
用友 NC IUpdateService XXE漏洞漏洞
用友 NC cloud accept任意文件上传漏洞
用友 U8 cloud KeyWordDetailReportQuery SQL注入漏洞
用友 GRP-U8 ufgovbank XXE漏洞
用友 移动系统管理 uploadIcon.do 任意文件上传漏洞
⽤友 NC-word.docx 文件读取漏洞
用友 portal 任意文件读取漏洞
用友 U8+ CRM help2.php 任意文件读取漏洞
用友 U8+ CRM getemaildata.php 任意文件读取漏洞
用友 GRP U8 license_check.jsp SQL注入漏洞
用友 GRP U8 SelectDMJE.jsp SQL注入漏洞
用友 GRP-U8 bx_historyDataCheck.jsp SQL注入漏洞
用友 U8 Cloud smartweb2.RPC.d XXE漏洞
用友 U8 obr_zdybxd_check SQL注入漏洞
用友 畅捷通Tplus GetStoreWarehouseByStore 远程代码执行漏洞
用友 畅捷通Tplus CheckMutex SQL注入漏洞
用友 畅捷通T+ GetDecAllUsers 信息泄露漏洞
用友 畅捷通远程通 GNRemote.dll SQL注入漏洞
```


