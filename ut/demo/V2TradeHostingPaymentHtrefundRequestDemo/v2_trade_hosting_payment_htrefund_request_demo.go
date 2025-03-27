/**
 * 托管交易退款 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeHostingPaymentHtrefundRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeHostingPaymentHtrefundRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeHostingPaymentHtrefundRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000003100616",
        // 申请退款金额
        OrdAmt:"0.01",
        // 原交易请求日期
        OrgReqDate:"20240229",
        // 安全信息线上交易退款必填，参见线上退款接口；jsonObject字符串
        RiskCheckData:get195595a868964f2bB023E9566fcd0297(),
        // 设备信息线上交易退款必填，参见线上退款接口；jsonObject字符串
        TerminalDeviceData:get8d8843c250f94e9b80a253d37ec6f80a(),
        // 大额转账支付账户信息数据jsonObject格式；银行大额转账支付交易退款申请时必填
        // BankInfoData:getAa3a4591240343e2Bad5D6a0764f06dc(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeHostingPaymentHtrefundRequest(dgReq)
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
    // 原交易全局流水号
    extendInfoMap["org_hf_seq_id"] = ""
    // 原交易微信支付宝的商户单号
    extendInfoMap["org_party_order_id"] = ""
    // 原交易请求流水号
    extendInfoMap["org_req_seq_id"] = "202207099803123123199941"
    // 分账对象
    extendInfoMap["acct_split_bunch"] = get4a68d378Cb6e41dfA9405a589b476160()
    // 备注
    // extendInfoMap["remark"] = ""
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.baidu.com"
    return extendInfoMap
}

func get33a52525B1614d3bBc18Ff7d935b2bca() interface{} {
    dto := make(map[string]interface{})
    // 分账金额
    dto["div_amt"] = "0.12"
    // 分账接收方ID
    dto["huifu_id"] = "6666000003100616"

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get4a68d378Cb6e41dfA9405a589b476160() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = get33a52525B1614d3bBc18Ff7d935b2bca()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get195595a868964f2bB023E9566fcd0297() string {
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

func get8d8843c250f94e9b80a253d37ec6f80a() string {
    dto := make(map[string]interface{})
    // 设备类型
    dto["device_type"] = "4"
    // 交易设备IP
    // dto["device_ip"] = ""
    // 交易设备MAC
    // dto["device_mac"] = ""
    // 交易设备GPS
    // dto["device_gps"] = ""
    // 交易设备IMEI
    // dto["device_imei"] = ""
    // 交易设备IMSI
    // dto["device_imsi"] = ""
    // 交易设备ICCID
    // dto["device_icc_id"] = ""
    // 交易设备WIFIMAC
    // dto["device_wifi_mac"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getAa3a4591240343e2Bad5D6a0764f06dc() string {
    dto := make(map[string]interface{})
    // 省份付款方为对公账户时必填，参见省市地区码；示例值：0013
    // dto["province"] = "test"
    // 地区付款方为对公账户时必填，参见省市地区码；示例值：1301
    // dto["area"] = "test"
    // 银行编号付款方为对公账户时必填，参考：银行编码； 示例值：01040000
    // dto["bank_code"] = "test"
    // 联行号付款方为对公账户时必填，参见：银行支行编码； 示例值：102290026507
    // dto["correspondent_code"] = "test"
    // 付款方账户类型
    // dto["card_acct_type"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

