/**
 * 快捷支付页面版 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeOnlinepaymentQuickpayFrontpayRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeOnlinepaymentQuickpayFrontpayRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeOnlinepaymentQuickpayFrontpayRequest{
        // 业务请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000109133323",
        // 订单金额
        TransAmt:"0.01",
        // 银行扩展信息
        ExtendPayData:getA5a44257248146f586ec9275ef51ed56(),
        // 设备信息
        TerminalDeviceData:get843a86e56f334304Ac7c38c6524abd82(),
        // 安全信息
        RiskCheckData:getE7f925f2Ab5041919687F439e2fba4ad(),
        // 异步通知地址
        NotifyUrl:"http://www.baidu.com",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeOnlinepaymentQuickpayFrontpayRequest(dgReq)
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
    // extendInfoMap["user_huifu_id"] = ""
    // 订单类型
    extendInfoMap["order_type"] = "P"
    // 订单失效时间
    extendInfoMap["time_expire"] = ""
    // 商品描述
    extendInfoMap["goods_desc"] = "快捷支付接口"
    // 请求类型
    extendInfoMap["request_type"] = "P"
    // 延时标记
    // extendInfoMap["delay_acct_flag"] = ""
    // 分账串
    extendInfoMap["acct_split_bunch"] = get82800b8cE35d42b9B688260754e5deef()
    // 手续费扣款标志
    extendInfoMap["fee_flag"] = "2"
    // 备注
    extendInfoMap["remark"] = "remark快捷支付接口"
    // 页面跳转地址
    extendInfoMap["front_url"] = "http://www.baidu.com"
    return extendInfoMap
}

func getDef0a935D6a343269064019353c3caf5() interface{} {
    dto := make(map[string]interface{})
    // 分账接收方ID
    dto["huifu_id"] = "6666000109133323"
    // 分账金额
    dto["div_amt"] = "0.01"
    // 账户号
    // dto["acct_id"] = ""
    // 分账百分比%
    // dto["percentage_div"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get82800b8cE35d42b9B688260754e5deef() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = getDef0a935D6a343269064019353c3caf5()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getA5a44257248146f586ec9275ef51ed56() string {
    dto := make(map[string]interface{})
    // 商品简称
    dto["goods_short_name"] = "01"
    // 业务种类
    dto["biz_tp"] = "123451"
    // 网关支付受理渠道
    dto["gw_chnnl_tp"] = "02"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get843a86e56f334304Ac7c38c6524abd82() string {
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

func getE7f925f2Ab5041919687F439e2fba4ad() string {
    dto := make(map[string]interface{})
    // ip地址
    dto["ip_addr"] = "127.0.0.1"
    // 基站地址
    // dto["base_station"] = ""
    // 纬度
    // dto["latitude"] = ""
    // 经度
    // dto["longitude"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

