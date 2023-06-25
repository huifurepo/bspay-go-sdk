/**
 * 手机WAP支付 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeOnlinepaymentWappayRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeOnlinepaymentWappayRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeOnlinepaymentWappayRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000103124174",
        // 交易金额
        TransAmt:"300.01",
        // 网联扩展数据
        ExtendPayData:getExtendPayData(),
        // 安全信息
        RiskCheckData:getRiskCheckData(),
        // 设备信息
        TerminalDeviceData:getTerminalDeviceData(),
        // 页面跳转地址
        FrontUrl:"http://www.baidu.com",
        // 异步通知地址
        NotifyUrl:"virgo://http://192.168.25.213:30030/sspc/onlineAsync",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeOnlinepaymentWappayRequest(dgReq)
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
    extendInfoMap["acct_split_bunch"] = getAcctSplitBunchRucan()
    // 银行卡号
    extendInfoMap["bank_card_no"] = "6222021102043040313"
    // 延时标记
    extendInfoMap["delay_acct_flag"] = "N"
    // 交易有效期
    extendInfoMap["time_expire"] = "20220406210038"
    // 分期期数
    extendInfoMap["instalments_num"] = "03"
    // 页面失败跳转地址
    extendInfoMap["front_fail_url"] = "http://www.baidu.com"
    // 备注
    extendInfoMap["remark"] = ""
    return extendInfoMap
}

func getAcctInfos() interface{} {
    dto := make(map[string]interface{})
    // 支付金额
    // dto["div_amt"] = ""
    // 被分账方ID
    // dto["huifu_id"] = ""
    // 账户号
    // dto["acct_id"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getAcctSplitBunchRucan() string {
    dto := make(map[string]interface{})
    // 分账信息列表
    dto["acct_infos"] = getAcctInfos()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getExtendPayData() string {
    dto := make(map[string]interface{})
    // 商品简称
    dto["goods_short_name"] = "一般商品"
    // 网关支付受理渠道
    dto["gw_chnnl_tp"] = "01"
    // 业务种类
    dto["biz_tp"] = "123456"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getRiskCheckData() string {
    dto := make(map[string]interface{})
    // ip地址
    dto["ip_addr"] = "111"
    // 基站地址
    // dto["base_station"] = ""
    // 纬度
    // dto["latitude"] = ""
    // 经度
    // dto["longitude"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getTerminalDeviceData() string {
    dto := make(map[string]interface{})
    // 交易设备ip
    dto["device_ip"] = "127.0.0.1"
    // 设备类型
    dto["device_type"] = "1"
    // 交易设备gps
    // dto["device_gps"] = ""
    // 交易设备iccid
    // dto["device_icc_id"] = ""
    // 交易设备imei
    // dto["device_imei"] = ""
    // 交易设备imsi
    // dto["device_imsi"] = ""
    // 交易设备mac
    // dto["device_mac"] = ""
    // 交易设备wifimac
    // dto["device_wifi_mac"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

