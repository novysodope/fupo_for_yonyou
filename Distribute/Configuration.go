package Distribute

import (
	"fmt"
	"fupo_for_yonyou/Scan"
	"log"
	"net/http"
	"time"
)

func ModuleConf(address string, client *http.Client, Red, Green, Yellow, Reset, Cyan string) {
	currentTime := time.Now().Format("15:04:05")

	// 1. 把所有要 runScan 的函数和名称都放到一个切片里
	scans := []struct {
		fn   func(string, *http.Client, string, string, string, string, string, string)
		name string
	}{
		{Scan.SyncScan, "SyncScan"},
		{Scan.GetSessionListScan, "GetSessionListScan"},
		{Scan.BshScan, "BshScan"},
		{Scan.YyOaSqlScan, "YyOaSqlScan"},
		{Scan.TemplateOfTaohong_managerScan, "TemplateOfTaohong_managerScan"},
		{Scan.NCFindWebScan, "NCFindWebScan"},
		{Scan.UploadFileDataScan, "UploadFileDataScan"},
		{Scan.GRPProxyScan, "GRPProxyScan"},
		{Scan.UapjsjndiScan, "UapjsjndiScan"},
		{Scan.GetusedspacesqlScan, "GetusedspacesqlScan"},
		{Scan.YCjtUploadScan, "YCjtUploadScan"},
		{Scan.RecoverPasswordScan, "RecoverPasswordScan"},
		{Scan.ServiceDispatcherServletScan, "ServiceDispatcherServletScan"},
		{Scan.U8AppProxyScan, "U8AppProxyScan"},
		{Scan.JspjndiyScan, "JspjndiyScan"},
		{Scan.JspjndieScan, "JspjndieScan"},
		{Scan.UapwsloginScan, "UapwsloginScan"},
		{Scan.AjaxjndiScan, "AjaxjndiScan"},
		{Scan.UapwsauhtScan, "UapwsauhtScan"},
		{Scan.FilesdeScan, "FilesdeScan"},
		{Scan.UapwsauthScan, "UapwsauthScan"},
		{Scan.FsdeScan, "FsdeScan"},
		{Scan.DownloadProxyScan, "DownloadProxyScan"},
		{Scan.ImageUploadScan, "ImageUploadScan"},
		{Scan.FileReceiveServletcan, "FileReceiveServletcan"},
		{Scan.AcceptuploadScan, "AcceptuploadScan"},
		{Scan.NCMessageServletScan, "NCMessageServletScan"},
		{Scan.UploadServletScan, "UploadServletScan"},
		{Scan.MonitorServletScan, "MonitorServletScan"},
		{Scan.IUpdateServicexxeScan, "IUpdateServicexxeScan"},
		{Scan.ServiceinforScan, "ServiceinforScan"},
		{Scan.KeyWordDetailReportQueryScan, "KeyWordDetailReportQueryScan"},
		{Scan.UfgovbankScan, "UfgovbankScan"},
		{Scan.MobileUploadIconScan, "MobileUploadIconScan"},
		{Scan.NcwordScan, "NcwordScan"},
		{Scan.PortalreadfileScan, "PortalreadfileScan"},
		{Scan.U8help2Scan, "U8help2Scan"},
		{Scan.U8getemaildataScan, "U8getemaildataScan"},
		{Scan.License_checkSQLiScan, "License_checkSQLiScan"},
		{Scan.SelectDMJEScan, "SelectDMJEScan"},
		{Scan.Bx_historyDataCheckScan, "Bx_historyDataCheckScan"},
		{Scan.Smartweb2XXEScan, "Smartweb2XXEScan"},
		{Scan.Obr_zdybxd_checkScan, "Obr_zdybxd_checkScan"},
		{Scan.GetStoreWarehouseByStoreScan, "GetStoreWarehouseByStoreScan"},
		{Scan.CheckMutexScan, "CheckMutexScan"},
		{Scan.GetDecAllUsersScan, "GetDecAllUsersScan"},
		{Scan.GNRemoteScan, "GNRemoteScan"},
		{Scan.DynamaticExportscan, "DynamaticExportscan"},
		{Scan.FilterCondActionscan, "FilterCondActionscan"},
		{Scan.ProcessSQLScan, "ProcessSQLScan"},
		{Scan.FormItemServletScan, "FormItemServletScan"},
		{Scan.BlobRefClassSearchScan, "BlobRefClassSearchScan"},
		{Scan.EsnserverScan, "EsnserverScan"},
		{Scan.AjaxgetborrowdataScan, "AjaxgetborrowdataScan"},
		{Scan.QueryPsnInfoScan, "QueryPsnInfoScan"},
		{Scan.MeasureQResultActionScan, "MeasureQResultActionScan"},
		{Scan.Yerfile_downscan, "Yerfile_downscan"},
		{Scan.CheckekeyScan, "CheckekeyScan"},
		{Scan.Link_contentScan, "Link_contentScan"},
		{Scan.ReservationcompletetScan, "ReservationcompletetScan"},
		{Scan.ReservationcompleteScan, "ReservationcompleteScan"},
		{Scan.ApproveservletScan, "ApproveservletScan"},
		{Scan.YonbiploginScan, "YonbiploginScan"},
		{Scan.MultiRepChooseActionScan, "MultiRepChooseActionScan"},
		{Scan.IMetaWebService4BqCloudScan, "IMetaWebService4BqCloudScan"},
	}

	for _, sc := range scans {
		runScan := func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("[%s%s%s] [%sERROR%s] %s 扫描发生错误: %v", Cyan, currentTime, Reset, Red, Reset, sc.name, r)
				}
			}()
			sc.fn(address, client, Red, Green, Yellow, Reset, Cyan, currentTime)
		}
		runScan()
	}

	// 3. 最后打印总数
	fmt.Printf("扫描完成，共扫描了 %d 个模块，请注意查看保存的结果\n", len(scans))
}
