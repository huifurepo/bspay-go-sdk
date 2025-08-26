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
    // extendInfoMap["acct_split_bunch"] = getB5e442b2D80a4cc6Be15F54bc7f0d17e()
    // 聚合正扫微信拓展参数集合
    // extendInfoMap["wx_data"] = getFe7352c8769647e3A81bB9d0415b2034()
    // 数字货币扩展参数集合
    // extendInfoMap["digital_currency_data"] = getE92256b7F6d84f2dAf8fA8a5dc08ab1b()
    // 补贴支付信息
    // extendInfoMap["combinedpay_data"] = get7c4f3ed720384dc39917F5f5df4cbb4f()
    // 补贴支付手续费承担方信息
    // extendInfoMap["combinedpay_data_fee_info"] = get4fac276fC02d4f44983a92b84cef9c23()
    // 备注
    // extendInfoMap["remark"] = ""
    // 是否垫资退款
    // extendInfoMap["loan_flag"] = ""
    // 垫资承担者
    // extendInfoMap["loan_undertaker"] = ""
    // 垫资账户类型
    // extendInfoMap["loan_acct_type"] = ""
    // 安全信息
    // extendInfoMap["risk_check_data"] = getDe1e2ccc3b334ad48bae4fb14fcdcf7e()
    // 设备信息
    // extendInfoMap["terminal_device_data"] = getB6899ac3F1a7474e8bee66d5c4886338()
    // 异步通知地址
    // extendInfoMap["notify_url"] = ""
    // 银联参数集合
    // extendInfoMap["unionpay_data"] = get0f6e2e25805844c4B41966db228e8658()
    return extendInfoMap
}

func getDe145ae4Bb634a9cAd7623b55067ba9a() interface{} {
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

func getB5e442b2D80a4cc6Be15F54bc7f0d17e() string {
    dto := make(map[string]interface{})
    // 分账信息列表
    // dto["acct_infos"] = getDe145ae4Bb634a9cAd7623b55067ba9a()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get90fa0a94Ad184e85A0065986f8c40499() interface{} {
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

func getF14468ebCe3e436c9bdb9acb8010e899() interface{} {
    dto := make(map[string]interface{})
    // 商品详情列表
    // dto["goods_detail"] = get90fa0a94Ad184e85A0065986f8c40499()

    return dto;
}

func getFe7352c8769647e3A81bB9d0415b2034() string {
    dto := make(map[string]interface{})
    // 退款商品详情
    // dto["detail"] = getF14468ebCe3e436c9bdb9acb8010e899()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getE92256b7F6d84f2dAf8fA8a5dc08ab1b() string {
    dto := make(map[string]interface{})
    // 退款原因
    // dto["refund_desc"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get7c4f3ed720384dc39917F5f5df4cbb4f() string {
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

func get4fac276fC02d4f44983a92b84cef9c23() string {
    dto := make(map[string]interface{})
    // 补贴支付手续费承担方汇付编号
    // dto["huifu_id"] = ""
    // 补贴支付手续费承担方账户号
    // dto["acct_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getDe1e2ccc3b334ad48bae4fb14fcdcf7e() string {
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

func getB6899ac3F1a7474e8bee66d5c4886338() string {
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

func get0f6e2e25805844c4B41966db228e8658() string {
    dto := make(map[string]interface{})
    // 收款方附加数据
    // dto["addn_data"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

