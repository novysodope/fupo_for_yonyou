package main

import (
	"fupo_for_yonyou/Distribute"
	"crypto/tls"
	"fmt"
	"golang.org/x/net/proxy"
	"net/http"
	"time"
)

func TargetParse(address string, proxyAddr string, Red string, Green string, Yellow string, Reset string, Cyan string) {
	if proxyAddr != "" {
		// 使用 socks5 代理创建 Dialer
		dialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
		if err != nil {
			fmt.Println("创建 SOCKS5 代理失败:", err)
			return
		}
		tr := &http.Transport{
			//socks5代理
			Dial: dialer.Dial,
		}
		client := &http.Client{
			Transport: tr,
			Timeout:   10 * time.Second,
		}
		client.Transport = tr
	}
	tr := &http.Transport{
		//忽略https证书错误
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}

	Distribute.ModuleConf(address, client, Red, Green, Yellow, Reset, Cyan)
}
