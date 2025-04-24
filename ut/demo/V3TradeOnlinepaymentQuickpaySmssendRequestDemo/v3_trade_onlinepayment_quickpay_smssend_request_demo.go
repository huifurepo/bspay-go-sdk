/**
 * 快捷短信发送 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V3TradeOnlinepaymentQuickpaySmssendRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V3TradeOnlinepaymentQuickpaySmssendRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V3TradeOnlinepaymentQuickpaySmssendRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000109133323",
        // 用户客户号
        UserHuifuId:"6666000149909053",
        // 绑卡id
        CardBindId:"10049847200",
        // 订单金额
        TransAmt:"10.00",
        // 异步通知地址
        NotifyUrl:"http://tianyi.demo.test.cn/core/extend/BsPaySdk/notify_quick.php",
        // 网联数据
        // NuccData:getA552c831Cd0846b8830f0ce1be990e53(),
        // 设备数据
        TerminalDeviceData:getEb10401d27a0468aA840Fd020391c341(),
        // 安全信息
        RiskCheckData:getBe09da814577476eA71c6d47f2fac2ac(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V3TradeOnlinepaymentQuickpaySmssendRequest(dgReq)
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
    // 分账对象
    // extendInfoMap["acct_split_bunch"] = get0fda4a7803b94a2fB5f17a083f8b015f()
    // 是否延迟交易
    // extendInfoMap["delay_acct_flag"] = ""
    // 账户号
    // extendInfoMap["acct_id"] = ""
    // 手续费扣款标志
    // extendInfoMap["fee_flag"] = ""
    // 补贴支付信息
    // extendInfoMap["combinedpay_data"] = get79bd2ee956e844ee8bb65e2f3ebf8096()
    return extendInfoMap
}

func getAdfb77e661f444deB2ddF0631bb7f4d4() interface{} {
    dto := make(map[string]interface{})
    // 分账接收方ID
    // dto["huifu_id"] = "test"
    // 分账金额
    // dto["div_amt"] = ""
    // 账户号
    // dto["acct_id"] = ""
    // 分账百分比%
    // dto["percentage_div"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get0fda4a7803b94a2fB5f17a083f8b015f() string {
    dto := make(map[string]interface{})
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""
    // 分账信息列表，Array格式
    // dto["acct_infos"] = getAdfb77e661f444deB2ddF0631bb7f4d4()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getA552c831Cd0846b8830f0ce1be990e53() string {
    dto := make(map[string]interface{})
    // 商品简称
    // dto["goods_short_name"] = "test"
    // 网关支付受理渠道
    // dto["gw_chnnl_tp"] = "test"
    // 业务种类
    // dto["biz_tp"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getEb10401d27a0468aA840Fd020391c341() string {
    dto := make(map[string]interface{})
    // 设备类型
    dto["device_type"] = "1"
    // 交易设备ip
    dto["device_ip"] = "106.33.180.238"
    // 交易设备mac
    // dto["device_mac"] = ""
    // 交易设备imei
    // dto["device_imei"] = ""
    // 交易设备imsi
    // dto["device_imsi"] = ""
    // 交易设备iccid
    // dto["device_icc_id"] = ""
    // 交易设备wifimac
    // dto["device_wifi_mac"] = ""
    // 交易设备gps
    // dto["device_gps"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getBe09da814577476eA71c6d47f2fac2ac() string {
    dto := make(map[string]interface{})
    // ip地址
    dto["ip_addr"] = "106.33.180.238"
    // 基站地址
    // dto["base_station"] = ""
    // 纬度
    // dto["latitude"] = ""
    // 经度
    // dto["longitude"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get79bd2ee956e844ee8bb65e2f3ebf8096() string {
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

