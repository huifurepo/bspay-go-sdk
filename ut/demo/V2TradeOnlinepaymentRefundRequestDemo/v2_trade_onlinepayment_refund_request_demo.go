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
        HuifuId:"6666000109133323",
        // 退款金额
        OrdAmt:"0.01",
        // 设备信息条件必填，当为银行大额支付时可不填，jsonObject格式
        TerminalDeviceData:get3875e406Fbd34508A04a334eb92da82a(),
        // 安全信息条件必填，当为银行大额支付时可不填，jsonObject格式
        RiskCheckData:getE8e44aaf76484b72B83fE1c7f081b40b(),
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
    extendInfoMap["org_req_date"] = "20240401"
    // 原交易全局流水号
    extendInfoMap["org_hf_seq_id"] = ""
    // 原交易请求流水号
    extendInfoMap["org_req_seq_id"] = "295700155481522176"
    // 分账对象
    // extendInfoMap["acct_split_bunch"] = get3c0b1469C829446a816d42c76304b250()
    // 备注
    // extendInfoMap["remark"] = ""
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.baidu.com"
    // 补贴支付信息
    // extendInfoMap["combinedpay_data"] = get4865eb8f0a3045f9B9ecE003a27b7391()
    // 大额转账支付账户信息数据
    // extendInfoMap["bank_info_data"] = getAd211f65098943119bdbC070a9366c76()
    return extendInfoMap
}

func getAcc12f300d964c559afdD701a3fad84f() interface{} {
    dto := make(map[string]interface{})
    // 商户号
    // dto["huifu_id"] = "test"
    // 支付金额
    // dto["div_amt"] = ""
    // 账户号
    // dto["acct_id"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get3c0b1469C829446a816d42c76304b250() string {
    dto := make(map[string]interface{})
    // 分账信息列表
    // dto["acct_infos"] = getAcc12f300d964c559afdD701a3fad84f()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get3875e406Fbd34508A04a334eb92da82a() string {
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

func getE8e44aaf76484b72B83fE1c7f081b40b() string {
    dto := make(map[string]interface{})
    // 经度
    // dto["longitude"] = "test"
    // 纬度
    // dto["latitude"] = "test"
    // 基站地址
    // dto["base_station"] = "test"
    // ip地址
    dto["ip_addr"] = "172.1.1.1"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get4865eb8f0a3045f9B9ecE003a27b7391() string {
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

func getAd211f65098943119bdbC070a9366c76() string {
    dto := make(map[string]interface{})
    // 银行编号
    // dto["bank_code"] = ""
    // 付款方账户类型
    // dto["card_acct_type"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

