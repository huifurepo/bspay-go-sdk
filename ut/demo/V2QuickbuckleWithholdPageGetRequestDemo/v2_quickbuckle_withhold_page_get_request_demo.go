/**
 * 代扣绑卡页面版 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2QuickbuckleWithholdPageGetRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2QuickbuckleWithholdPageGetRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2QuickbuckleWithholdPageGetRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 汇付Id
        HuifuId:"6666000103422897",
        // 订单号
        OrderId:"SAS20230807102136898274149",
        // 订单日期
        OrderDate:"20230807",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2QuickbuckleWithholdPageGetRequest(dgReq)
  	if err != nil {
		fmt.Println("请求异常:", err)
	}

	fmt.Println("返回数据:", resp)
}

/**
 * 非必填字段
 */
func getExtendInfos() map[string]interface{} {
    // 设置非必填字段
    extendInfoMap := make(map[string]interface{})
    // 页面有效期
    extendInfoMap["expire_time"] = "20"
    // 页面跳转地址
    extendInfoMap["return_url"] = "https://api.huifu.com"
    // 客户用户号
    extendInfoMap["out_cust_id"] = "Q837467382819"
    // 用户汇付号
    extendInfoMap["user_huifu_id"] = "6666000107386236"
    // 异步通知地址
    extendInfoMap["notify_url"] = "https://api.huifu.com"
    // 设备信息域
    extendInfoMap["trx_device_info"] = getTrxDeviceInfo()
    // 风控信息
    extendInfoMap["risk_info"] = getRiskInfo()
    return extendInfoMap
}

func getTrxDeviceInfo() string {
    dto := make(map[string]interface{})
    // 银行预留手机号
    dto["trx_mobile_num"] = "13428722321"
    // 设备类型
    dto["trx_device_type"] = "1"
    // 交易设备IP
    dto["trx_device_ip"] = "192.168.1.1"
    // 交易设备MAC
    dto["trx_device_mac"] = "10.10.0.1"
    // 交易设备IMEI
    dto["trx_device_imei"] = "imei"
    // 交易设备IMSI
    dto["trx_device_imsi"] = "imsi"
    // 交易设备ICCID
    dto["trx_device_icc_id"] = "icc"
    // 交易设备WIFIMAC
    dto["trx_device_wfifi_mac"] = "wfifi"
    // 交易设备GPS
    dto["trx_device_gps"] = "gps"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getRiskInfo() string {
    dto := make(map[string]interface{})
    // IP类型
    dto["ip_type"] = "04"
    // IP值
    dto["source_ip"] = "1.1.1.1"
    // 设备标识
    dto["device_id"] = ""
    // 设备类型
    dto["device_type"] = ""
    // 银行预留手机号
    dto["mobile"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

