/**
 * 线上交易退款 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeOnlinepaymentRefundRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeOnlinepaymentRefundRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeOnlinepaymentRefundRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000108854952",
        // 退款金额
        OrdAmt:"0.01",
        // 设备信息
        TerminalDeviceData:getTerminalDeviceData(),
        // 安全信息
        RiskCheckData:getRiskCheckData(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeOnlinepaymentRefundRequest(dgReq)
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
    // 原交易请求日期
    extendInfoMap["org_req_date"] = "20221110"
    // 原交易全局流水号
    extendInfoMap["org_hf_seq_id"] = ""
    // 原交易请求流水号
    extendInfoMap["org_req_seq_id"] = "RQ1212333113"
    // 分账对象
    extendInfoMap["acct_split_bunch"] = getAcctSplitBunchRucan()
    // 补贴支付信息
    extendInfoMap["combinedpay_data"] = getCombinedpayData()
    // 大额转账支付账户信息数据
    // extendInfoMap["bank_info_data"] = getBankInfoData()
    // 备注
    extendInfoMap["remark"] = "remark123"
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.baidu.com"
    return extendInfoMap
}

func getAcctInfosRucan() interface{} {
    dto := make(map[string]interface{})
    // 商户号
    // dto["huifu_id"] = ""
    // 支付金额
    // dto["div_amt"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getAcctSplitBunchRucan() string {
    dto := make(map[string]interface{})
    // 分账信息列表
    // dto["acct_infos"] = getAcctInfosRucan()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getCombinedpayData() string {
    dto := make(map[string]interface{})
    // 补贴方汇付编号
    // dto["huifu_id"] = "test"
    // 补贴方类型
    // dto["user_type"] = "test"
    // 补贴方账户号
    // dto["acct_id"] = "test"
    // 补贴金额
    // dto["amount"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getBankInfoData() string {
    dto := make(map[string]interface{})
    // 付款方账户类型
    // dto["card_acct_type"] = "test"
    // 省份
    // dto["province"] = ""
    // 地区
    // dto["area"] = ""
    // 银行编号
    // dto["bank_code"] = ""
    // 联行号
    // dto["correspondent_code"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getTerminalDeviceData() string {
    dto := make(map[string]interface{})
    // 交易设备ip
    dto["device_ip"] = "172.31.31.145"
    // 设备类型
    dto["device_type"] = "1"
    // 交易设备gps
    dto["device_gps"] = "07"
    // 交易设备iccid
    dto["device_icc_id"] = "05"
    // 交易设备imei
    dto["device_imei"] = "02"
    // 交易设备imsi
    dto["device_imsi"] = "03"
    // 交易设备mac
    dto["device_mac"] = "01"
    // 交易设备wifimac
    dto["device_wifi_mac"] = "06"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getRiskCheckData() string {
    dto := make(map[string]interface{})
    // ip地址
    // dto["ip_addr"] = ""
    // 基站地址
    // dto["base_station"] = ""
    // 纬度
    // dto["latitude"] = ""
    // 经度
    // dto["longitude"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

