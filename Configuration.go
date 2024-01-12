package main

import (
	"net/http"
	"time"
	"fupo_for_yonyou/Scan"
)

func moduleConf(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string) {
	currentTime := time.Now().Format("15:04:05")
	syncScan(address, Red, Green, Yellow, Reset, Cyan, currentTime)
	//增加自定义yaml支持
}
