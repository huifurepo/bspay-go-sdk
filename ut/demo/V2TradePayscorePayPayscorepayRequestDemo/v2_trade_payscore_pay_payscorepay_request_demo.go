/**
 * 支付分扣款 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePayscorePayPayscorepayRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePayscorePayPayscorepayRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePayscorePayPayscorepayRequest{
        // 微信扣款单号
        // OutTradeNo:"test",
        // 商品描述
        // GoodsDesc:"test",
        // 商户号
        HuifuId:"6666000108854952",
        // 安全信息
        // RiskCheckData:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradePayscorePayPayscorepayRequest(dgReq)
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
    // 分账对象
    // extendInfoMap["acct_split_bunch"] = getAcctSplitBunch()
    // 扣款登记返回的汇付全局流水号
    // extendInfoMap["deduct_hf_seq_id"] = ""
    // 扣款登记创建请求流水号
    // extendInfoMap["deduct_req_seq_id"] = ""
    // 是否延迟交易
    // extendInfoMap["delay_acct_flag"] = ""
    // 商户回调地址
    // extendInfoMap["notify_url"] = ""
    // 交易备注
    // extendInfoMap["remark"] = ""
    // 请求日期
    extendInfoMap["req_date"] = tool.GetCurrentDate()
    // 请求流水号
    extendInfoMap["req_seq_id"] = tool.GetReqSeqId()
    // 设备信息
    // extendInfoMap["terminal_device_info"] = getTerminalDeviceInfo()
    // 聚合反扫微信参数集合
    // extendInfoMap["wx_data"] = getWxData()
    return extendInfoMap
}

func getAcctInfos() interface{} {
    dto := make(map[string]interface{})
    // 分账金额
    // dto["div_amt"] = "test"
    // 分账商户号
    // dto["huifu_id"] = "test"

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getAcctSplitBunch() string {
    dto := make(map[string]interface{})
    // 分账明细
    // dto["acct_infos"] = getAcctInfos()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getTerminalDeviceInfo() string {
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
    // 商户终端应用程序版
    // dto["app_version"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getWxData() string {
    dto := make(map[string]interface{})
    // 子商户公众账号id
    // dto["sub_appid"] = ""
    // 用户标识
    // dto["openid"] = ""
    // 子商户用户标识
    // dto["sub_openid"] = ""
    // 设备号
    // dto["device_info"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

