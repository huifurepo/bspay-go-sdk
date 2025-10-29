/**
 * 银联统一在线收银台 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeOnlinepaymentUnionpayRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeOnlinepaymentUnionpayRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeOnlinepaymentUnionpayRequest{
        // 商户号
        HuifuId:"6666000109133323",
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 订单金额
        TransAmt:"0.11",
        // 商品描述
        OrderDesc:"通用性商品1",
        // 安全信息
        RiskCheckData:get72948c317e164ae79d4a93cef8895c95(),
        // 三方支付数据jsonObject&lt;br/&gt;pay_scene&#x3D;U_JSAPI或pay_scene&#x3D;U_MINIAPP时，必填
        // ThirdPayData:getE6adb6f3591a41ec9c6137c27ca2398d(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeOnlinepaymentUnionpayRequest(dgReq)
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
    // 卡号锁定标识
    extendInfoMap["card_number_lock"] = ""
    // 直通模式的银行标识
    extendInfoMap["ebank_en_abbr"] = ""
    // 交易银行卡卡号
    extendInfoMap["pay_card_no"] = ""
    // 支付卡类型
    // extendInfoMap["pay_card_type"] = ""
    // 订单失效时间
    extendInfoMap["time_expire"] = ""
    // 前端跳转地址
    extendInfoMap["front_url"] = "https://www.service.com/getresp"
    // 异步通知地址
    extendInfoMap["notify_url"] = "https://www.service.com/getresp"
    // 备注
    extendInfoMap["remark"] = "merPriv11"
    // 支付场景
    // extendInfoMap["pay_scene"] = ""
    // 签约令牌号
    // extendInfoMap["sign_token_no"] = ""
    // 延时标记
    extendInfoMap["delay_acct_flag"] = "Y"
    // 手续费扣款标志
    // extendInfoMap["fee_flag"] = ""
    // 分账对象
    // extendInfoMap["acct_split_bunch"] = get938f9796E5bd4f85A4cf2ce4a306837d()
    // 设备信息数据
    // extendInfoMap["terminal_device_data"] = get86e6396a54f5492cB3b06678723bed4b()
    return extendInfoMap
}

func get0aa0431065074e9f92016d99005ce03e() interface{} {
    dto := make(map[string]interface{})
    // 分账金额
    // dto["div_amt"] = ""
    // 分账接收方ID
    // dto["huifu_id"] = ""
    // 账户号
    // dto["acct_id"] = ""
    // 分账百分比%
    // dto["percentage_div"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get938f9796E5bd4f85A4cf2ce4a306837d() string {
    dto := make(map[string]interface{})
    // 分账明细
    // dto["acct_infos"] = get0aa0431065074e9f92016d99005ce03e()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get72948c317e164ae79d4a93cef8895c95() string {
    dto := make(map[string]interface{})
    // 基站地址
    dto["base_station"] = "7"
    // ip地址
    dto["ip_addr"] = "172.28.52.52"
    // 纬度
    dto["latitude"] = "4"
    // 经度
    dto["longitude"] = "3"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getE6adb6f3591a41ec9c6137c27ca2398d() string {
    dto := make(map[string]interface{})
    // 小程序id
    // dto["app_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get86e6396a54f5492cB3b06678723bed4b() string {
    dto := make(map[string]interface{})
    // 终端设备号
    // dto["device_id"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

