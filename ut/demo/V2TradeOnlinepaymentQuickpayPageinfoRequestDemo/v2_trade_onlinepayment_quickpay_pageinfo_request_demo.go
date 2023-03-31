/**
 * 快捷支付页面版接口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeOnlinepaymentQuickpayPageinfoRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeOnlinepaymentQuickpayPageinfoRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeOnlinepaymentQuickpayPageinfoRequest{
        // 业务请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000108854952",
        // 订单金额
        TransAmt:"1.02",
        // 设备信息
        TerminalDeviceData:getTerminalDeviceData(),
        // 银行扩展信息
        ExtendPayData:getExtendPayData(),
        // 安全信息
        RiskCheckData:getRiskCheckData(),
        // 异步通知地址
        NotifyUrl:"http://www.chinapnr.com",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeOnlinepaymentQuickpayPageinfoRequest(dgReq)
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
    // 用户客户号
    extendInfoMap["user_huifu_id"] = ""
    // 订单类型
    extendInfoMap["order_type"] = "P"
    // 请求类型
    extendInfoMap["request_type"] = "P"
    // 订单失效时间
    extendInfoMap["time_expire"] = ""
    // 备注
    extendInfoMap["remark"] = "remark快捷支付接口"
    // 分账串
    // extendInfoMap["acct_split_bunch"] = getAcctSplitBunchRucan()
    // 页面跳转地址
    extendInfoMap["front_url"] = "http://www.chinapnr.com"
    // 延时标记
    // extendInfoMap["delay_acct_flag"] = ""
    // 手续费扣款标志
    // extendInfoMap["fee_flag"] = ""
    return extendInfoMap
}

func getAcctInfos() interface{} {
    dto := make(map[string]interface{})
    // 被分账对象ID
    // dto["huifu_id"] = "test"
    // 分账金额
    // dto["div_amt"] = "test"
    // 账户号
    // dto["acct_id"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getAcctSplitBunchRucan() string {
    dto := make(map[string]interface{})
    // 分账明细
    // dto["acct_infos"] = getAcctInfos()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getTerminalDeviceData() string {
    dto := make(map[string]interface{})
    // 设备类型
    dto["device_type"] = "1"
    // 交易设备IP
    dto["device_ip"] = "127.0.0.1"
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

func getExtendPayData() string {
    dto := make(map[string]interface{})
    // 网关支付受理渠道
    dto["gw_chnnl_tp"] = "02"
    // 商品简称
    dto["goods_short_name"] = "01"
    // 业务种类
    dto["biz_tp"] = "123451"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getRiskCheckData() string {
    dto := make(map[string]interface{})
    // ip地址
    dto["ip_addr"] = "127.0.0.1"
    // 基站地址
    // dto["base_atation"] = ""
    // 纬度
    // dto["latitude"] = ""
    // 经度
    // dto["longitude"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

