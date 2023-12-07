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
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000141569791",
        // 扣款登记创建请求流水号deduct_req_seq_id与deduct_hf_seq_id二选一；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：2022012614120615001&lt;/font&gt;
        DeductReqSeqId:"1726841301594394626",
        // 扣款登记返回的汇付全局流水号deduct_req_seq_id与deduct_hf_seq_id二选一；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：00470topo1A211015160805P090ac132fef00000&lt;/font&gt;
        // DeductHfSeqId:"test",
        // 微信扣款单号
        OutTradeNo:"03212311224952047516172",
        // 商品描述
        GoodsDesc:"bp充电",
        // 安全信息
        RiskCheckData:getRiskCheckData(),
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
    // 聚合反扫微信参数集合
    // extendInfoMap["wx_data"] = getWxData()
    // 是否延迟交易
    // extendInfoMap["delay_acct_flag"] = ""
    // 分账对象
    // extendInfoMap["acct_split_bunch"] = getAcctSplitBunch()
    // 设备信息
    // extendInfoMap["terminal_device_info"] = getTerminalDeviceInfo()
    // 交易备注
    // extendInfoMap["remark"] = ""
    // 商户回调地址
    // extendInfoMap["notify_url"] = ""
    return extendInfoMap
}

func getWxData() string {
    dto := make(map[string]interface{})
    // 子商户用户标识
    // dto["sub_openid"] = "test"
    // 子商户公众账号id
    // dto["sub_appid"] = ""
    // 用户标识
    // dto["openid"] = ""
    // 设备号
    // dto["device_info"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
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

func getRiskCheckData() string {
    dto := make(map[string]interface{})
    // ip地址
    dto["ip_address"] = "127.0.0.1"
    // 基站地址
    // dto["base_station"] = ""
    // 纬度
    // dto["latitude"] = ""
    // 经度
    // dto["longitude"] = ""

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

