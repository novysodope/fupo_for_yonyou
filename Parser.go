package main

import (
	"crypto/tls"
	"fmt"
	"fupo_for_yonyou/Distribute"
	"golang.org/x/net/proxy"
	"net/http"
	"net/url"
	"time"
)

func TargetParse(address string, httpProxy string, socks5Addr string, Red string, Green string, Yellow string, Reset string, Cyan string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if httpProxy != "" {
		if socks5Addr != "" {
			fmt.Println("[" + Magenta + "WARN" + Reset + "] HTTP 代理启用中，SOCKS5 代理已被忽略")
		}
		proxyURL, err := url.Parse(httpProxy)
		if err != nil {
			fmt.Println("解析 HTTP 代理地址失败:", err)
			return
		}
		tr.Proxy = http.ProxyURL(proxyURL)

	} else if socks5Addr != "" {
		dialer, err := proxy.SOCKS5("tcp", socks5Addr, nil, proxy.Direct)
		if err != nil {
			fmt.Println("创建 SOCKS5 代理失败:", err)
			return
		}
		tr.Dial = dialer.Dial
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	Distribute.ModuleConf(address, client, Red, Green, Yellow, Reset, Cyan)
}
