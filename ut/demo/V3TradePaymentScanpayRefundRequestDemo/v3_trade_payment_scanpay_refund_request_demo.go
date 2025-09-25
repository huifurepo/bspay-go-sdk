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
    // extendInfoMap["acct_split_bunch"] = get8e93a7827a1d46e888a74d04308badd6()
    // 聚合正扫微信拓展参数集合
    // extendInfoMap["wx_data"] = get5b74ce8bEc2c45a58286447e7e6621cb()
    // 数字货币扩展参数集合
    // extendInfoMap["digital_currency_data"] = get4c1bd5af86214ed6974bC5ff5d64ca9c()
    // 补贴支付信息
    // extendInfoMap["combinedpay_data"] = getCb401253Bd744c95Aa62C8e630c58087()
    // 补贴支付手续费承担方信息
    // extendInfoMap["combinedpay_data_fee_info"] = get9f7f55ed457e450dBfb1863a3c2af720()
    // 备注
    // extendInfoMap["remark"] = ""
    // 是否垫资退款
    // extendInfoMap["loan_flag"] = ""
    // 垫资承担者
    // extendInfoMap["loan_undertaker"] = ""
    // 垫资账户类型
    // extendInfoMap["loan_acct_type"] = ""
    // 安全信息
    // extendInfoMap["risk_check_data"] = get550ccaf0Be804a21A6c09ddc20cc2fcd()
    // 设备信息
    // extendInfoMap["terminal_device_data"] = get83812a8eE7ad4bd39aa3A624b5f7f788()
    // 异步通知地址
    // extendInfoMap["notify_url"] = ""
    // 银联参数集合
    // extendInfoMap["unionpay_data"] = get622f9aa969c64703B9883e664feb9a03()
    return extendInfoMap
}

func getC12608a1C366480b84d8Efe811b82566() interface{} {
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

func get8e93a7827a1d46e888a74d04308badd6() string {
    dto := make(map[string]interface{})
    // 分账信息列表
    // dto["acct_infos"] = getC12608a1C366480b84d8Efe811b82566()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get3cb43b9c37414e9c8d6eC25984bf20c0() interface{} {
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

func get54250eda110c479683a834054605a71e() interface{} {
    dto := make(map[string]interface{})
    // 商品详情列表
    // dto["goods_detail"] = get3cb43b9c37414e9c8d6eC25984bf20c0()

    return dto;
}

func get5b74ce8bEc2c45a58286447e7e6621cb() string {
    dto := make(map[string]interface{})
    // 退款商品详情
    // dto["detail"] = get54250eda110c479683a834054605a71e()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get4c1bd5af86214ed6974bC5ff5d64ca9c() string {
    dto := make(map[string]interface{})
    // 退款原因
    // dto["refund_desc"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getCb401253Bd744c95Aa62C8e630c58087() string {
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

func get9f7f55ed457e450dBfb1863a3c2af720() string {
    dto := make(map[string]interface{})
    // 补贴支付手续费承担方汇付编号
    // dto["huifu_id"] = ""
    // 补贴支付手续费承担方账户号
    // dto["acct_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get550ccaf0Be804a21A6c09ddc20cc2fcd() string {
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

func get83812a8eE7ad4bd39aa3A624b5f7f788() string {
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

func get622f9aa969c64703B9883e664feb9a03() string {
    dto := make(map[string]interface{})
    // 收款方附加数据
    // dto["addn_data"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

