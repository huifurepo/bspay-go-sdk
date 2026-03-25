/**
 * 扫码交易退款 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V3TradePaymentScanpayRefundRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V3TradePaymentScanpayRefundRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V3TradePaymentScanpayRefundRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000108854952",
        // 申请退款金额
        OrdAmt:"0.01",
        // 原交易请求日期
        OrgReqDate:"20221107",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V3TradePaymentScanpayRefundRequest(dgReq)
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
    extendInfoMap["org_hf_seq_id"] = "002900TOP3B221107142320P992ac139c0c00000"
    // 原交易微信支付宝的商户单号
    // extendInfoMap["org_party_order_id"] = ""
    // 原交易请求流水号
    // extendInfoMap["org_req_seq_id"] = ""
    // 分账对象
    // extendInfoMap["acct_split_bunch"] = get641df3f4C94748d48e4635007f755f47()
    // 微信拓展参数集合
    // extendInfoMap["wx_data"] = get81730c8fEec745f89ed55ddca843c35d()
    // 数字货币拓展参数集合
    // extendInfoMap["digital_currency_data"] = get3f660ca48e1249a8870bB36cddbf53df()
    // 抖音拓展参数集合
    // extendInfoMap["dy_data"] = get7380711a77c943abB65dFd0262d2a339()
    // 补贴支付信息
    // extendInfoMap["combinedpay_data"] = getB48bed7aB6174f51Af5d8d93d208f004()
    // 补贴支付手续费承担方信息
    // extendInfoMap["combinedpay_data_fee_info"] = get337cdd659f544810B6ee091fb21950fd()
    // 备注
    // extendInfoMap["remark"] = ""
    // 是否垫资退款
    // extendInfoMap["loan_flag"] = ""
    // 垫资承担者
    // extendInfoMap["loan_undertaker"] = ""
    // 垫资账户类型
    // extendInfoMap["loan_acct_type"] = ""
    // 安全信息
    // extendInfoMap["risk_check_data"] = get857f190d4de14c95812b472db2d4546d()
    // 设备信息
    // extendInfoMap["terminal_device_data"] = get5b739bb834934f4cBf091bfe6f6b1751()
    // 异步通知地址
    // extendInfoMap["notify_url"] = ""
    // 银联参数集合
    // extendInfoMap["unionpay_data"] = get2c17c847649d4d3a84bdFff0a7b9bb33()
    return extendInfoMap
}

func getA259171491144b858892C854abdc3633() interface{} {
    dto := make(map[string]interface{})
    // 分账金额
    // dto["div_amt"] = "test"
    // 分账接收方ID
    // dto["huifu_id"] = "test"
    // 垫资金额
    // dto["part_loan_amt"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get641df3f4C94748d48e4635007f755f47() string {
    dto := make(map[string]interface{})
    // 分账信息列表
    // dto["acct_infos"] = getA259171491144b858892C854abdc3633()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get1c5125870a39427f96d928170154a26d() interface{} {
    dto := make(map[string]interface{})
    // 商品编码
    // dto["goods_id"] = "test"
    // 商品退款金额
    // dto["refund_amt"] = "test"
    // 商品退款数量
    // dto["refund_quantity"] = "test"
    // 商品单价
    // dto["price"] = "test"
    // 微信支付商品编码
    // dto["wxpay_goods_id"] = ""
    // 商品名称
    // dto["goods_name"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getB1dda134445a4afe9f3b1c23e906b0d0() interface{} {
    dto := make(map[string]interface{})
    // 商品详情列表
    // dto["goods_detail"] = get1c5125870a39427f96d928170154a26d()

    return dto;
}

func get81730c8fEec745f89ed55ddca843c35d() string {
    dto := make(map[string]interface{})
    // 退款商品详情
    // dto["detail"] = getB1dda134445a4afe9f3b1c23e906b0d0()
    // 退款原因
    // dto["refund_desc"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get3f660ca48e1249a8870bB36cddbf53df() string {
    dto := make(map[string]interface{})
    // 退款原因
    // dto["refund_desc"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get7380711a77c943abB65dFd0262d2a339() string {
    dto := make(map[string]interface{})
    // 退款原因
    // dto["refund_desc"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getB48bed7aB6174f51Af5d8d93d208f004() string {
    dto := make(map[string]interface{})
    // 汇付商户号
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

func get337cdd659f544810B6ee091fb21950fd() string {
    dto := make(map[string]interface{})
    // 补贴支付手续费承担方汇付编号
    // dto["huifu_id"] = ""
    // 补贴支付手续费承担方账户号
    // dto["acct_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get857f190d4de14c95812b472db2d4546d() string {
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

func get5b739bb834934f4cBf091bfe6f6b1751() string {
    dto := make(map[string]interface{})
    // 设备类型
    // dto["device_type"] = ""
    // 交易设备IP
    // dto["device_ip"] = ""
    // 交易设备MAC
    // dto["device_mac"] = ""
    // 交易设备IMEI
    // dto["device_imei"] = ""
    // 交易设备IMSI
    // dto["device_imsi"] = ""
    // 交易设备ICCID
    // dto["device_icc_id"] = ""
    // 交易设备WIFIMAC
    // dto["device_wifi_mac"] = ""
    // 交易设备GPS
    // dto["device_gps"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get2c17c847649d4d3a84bdFff0a7b9bb33() string {
    dto := make(map[string]interface{})
    // 收款方附加数据
    // dto["addn_data"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

