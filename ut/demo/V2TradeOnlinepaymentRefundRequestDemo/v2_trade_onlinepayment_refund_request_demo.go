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
        TerminalDeviceData:get35d4a53eD6c1411a9ced250654fb32bc(),
        // 安全信息条件必填，当为银行大额支付时可不填，jsonObject格式
        RiskCheckData:getF25a2614657744d4Bc59741b0034991d(),
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
    // extendInfoMap["acct_split_bunch"] = getB2063bd6B0444da8A946Df04df4862fd()
    // 备注
    // extendInfoMap["remark"] = ""
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.baidu.com"
    // 补贴支付信息
    // extendInfoMap["combinedpay_data"] = getF6ac280047d84b54Af3cFc21f6942cf0()
    // 大额转账支付账户信息数据
    // extendInfoMap["bank_info_data"] = get289bf626B24d46cdA273Fdaa77310b55()
    return extendInfoMap
}

func get1e71bbb43d3d4de7A4a590b7e2ca4d2c() interface{} {
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

func getB2063bd6B0444da8A946Df04df4862fd() string {
    dto := make(map[string]interface{})
    // 分账信息列表
    // dto["acct_infos"] = get1e71bbb43d3d4de7A4a590b7e2ca4d2c()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get35d4a53eD6c1411a9ced250654fb32bc() string {
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
    // 终端设备号
    // dto["device_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getF25a2614657744d4Bc59741b0034991d() string {
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

func getF6ac280047d84b54Af3cFc21f6942cf0() string {
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

func get289bf626B24d46cdA273Fdaa77310b55() string {
    dto := make(map[string]interface{})
    // 银行编号
    // dto["bank_code"] = ""
    // 付款方账户类型
    // dto["card_acct_type"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

