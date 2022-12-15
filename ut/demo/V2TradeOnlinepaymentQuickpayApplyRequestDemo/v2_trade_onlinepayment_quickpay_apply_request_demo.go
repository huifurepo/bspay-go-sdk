/**
 * 快捷支付申请 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeOnlinepaymentQuickpayApplyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeOnlinepaymentQuickpayApplyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeOnlinepaymentQuickpayApplyRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000119640000",
        // 订单金额
        TransAmt:"1980.00",
        // 绑卡id
        CardBindId:"10032850000",
        // 异步通知地址
        NotifyUrl:"http://tianyi.demo.test.cn/core/extend/BsPaySdk/notify_quick.php",
        // 用户客户号
        UserHuifuId:"6666000121370000",
        // 安全信息
        RiskCheckData:getRiskCheckData(),
        // 设备数据
        TerminalDeviceData:getTerminalDeviceData(),
        // 银行扩展字段
        ExtendPayData:getExtendPayData(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeOnlinepaymentQuickpayApplyRequest(dgReq)
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
    // 订单类型
    // extendInfoMap["order_type"] = ""
    // 备注
    // extendInfoMap["remark"] = ""
    // 订单失效时间
    // extendInfoMap["time_expire"] = ""
    // 营销补贴信息
    // extendInfoMap["combinedpay_data"] = getCombinedpayData()
    // 分账对象
    // extendInfoMap["acct_split_bunch"] = getAcctSplitBunchRucan()
    return extendInfoMap
}

func getCombinedpayData() string {
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

func getTerminalDeviceData() string {
    dto := make(map[string]interface{})
    // 交易设备ip
    dto["device_ip"] = "106.33.180.238"
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

func getRiskCheckData() string {
    dto := make(map[string]interface{})
    // ip地址
    dto["ip_addr"] = "106.33.180.238"
    // 基站地址
    // dto["base_atation"] = ""
    // 纬度
    // dto["latitude"] = ""
    // 经度
    // dto["longitude"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getAcctInfos() interface{} {
    dto := make(map[string]interface{})
    // 分账金额
    // dto["div_amt"] = ""
    // 被分账方ID
    // dto["huifu_id"] = ""
    // 账户号
    // dto["acct_id"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getExtendPayData() string {
    dto := make(map[string]interface{})
    // 业务种类
    // dto["biz_tp"] = "test"
    // 商品简称
    // dto["goods_short_name"] = "test"
    // 网关支付受理渠道
    dto["gw_chnnl_tp"] = "99"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getAcctSplitBunchRucan() string {
    dto := make(map[string]interface{})
    // 分账信息列表
    // dto["acct_infos"] = getAcctInfos()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

