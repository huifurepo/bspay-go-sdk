/**
 * 微信支付宝预授权完成撤销 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePaymentPreauthpaycancelRefundRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePaymentPreauthpaycancelRefundRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePaymentPreauthpaycancelRefundRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 客户号
        HuifuId:"6666000108854952",
        // 原预授权完成交易请求日期
        OrgReqDate:"20221031",
        // 完成撤销金额
        OrdAmt:"0.02",
        // 风控信息
        RiskCheckInfo:getB6b94958D58f402dA1271d898ffe2200(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradePaymentPreauthpaycancelRefundRequest(dgReq)
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
    // 外部订单号
    extendInfoMap["out_ord_Id"] = ""
    // 原预授权完成交易请求流水号
    extendInfoMap["org_req_seq_id"] = "20211667205111"
    // 交易发起时间
    extendInfoMap["send_time"] = "312321321321"
    // 商品描述
    extendInfoMap["good_desc"] = "商户描述商户描述商户描述商户描述商户描述"
    // 是否人工介入
    extendInfoMap["is_manual_process"] = "Y"
    // 备注
    extendInfoMap["remark"] = "商户描述商户描述商户描述商户描述商户描述"
    // 汇付机具号
    extendInfoMap["devs_id"] = "SPINTP366061800405501"
    // 商户操作员号
    // extendInfoMap["mer_oper_id"] = ""
    // 批次号
    // extendInfoMap["batch_id"] = ""
    // 扩展域
    // extendInfoMap["mer_priv"] = ""
    // 设备信息
    extendInfoMap["terminal_device_info"] = getB4cf3131629a41389d2285e0202d1062()
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.baidu.com"
    return extendInfoMap
}

func getB6b94958D58f402dA1271d898ffe2200() string {
    dto := make(map[string]interface{})
    // 基站地址
    dto["base_station"] = "192.168.1.1"
    // ip地址
    dto["ip_addr"] = "192.168.1.1"
    // 纬度
    dto["latitude"] = "33.3"
    // 经度
    dto["longitude"] = "33.3"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getB4cf3131629a41389d2285e0202d1062() string {
    dto := make(map[string]interface{})
    // 交易设备GPS
    dto["device_gps"] = "192.168.0.0"
    // 交易设备ICCID
    dto["device_icc_id"] = ""
    // 交易设备IMEI
    dto["device_imei"] = ""
    // 交易设备IMSI
    dto["device_imsi"] = ""
    // 交易设备IP
    dto["device_ip"] = "10.10.0.1"
    // 交易设备MAC
    dto["device_mac"] = ""
    // 设备类型
    dto["device_type"] = "1"
    // 交易设备WIFIMAC
    dto["device_wifi_mac"] = ""
    // 汇付机具号
    dto["devs_id"] = "SPINTP366061800405501"
    // 操作员号
    dto["mer_oper_id"] = ""
    // 逻辑终端号
    dto["pnr_dev_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

