package main

import (
	"net/http"
	"time"
)

func moduleConf(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string) {
	currentTime := time.Now().Format("15:04:05")
	syncScan(address, currentTime, Green, Yellow, Reset, Cyan, Red)
}
