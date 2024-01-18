package Distribute

import (
	"FoPoForYonyou2.0/Scan"
	"net/http"
	"time"
)

func ModuleConf(address string, client *http.Client, Red string, Green string, Yellow string, Reset string, Cyan string) {
	currentTime := time.Now().Format("15:04:05")
	Scan.SyncScan(address, Red, Green, Yellow, Reset, Cyan, currentTime)
}
