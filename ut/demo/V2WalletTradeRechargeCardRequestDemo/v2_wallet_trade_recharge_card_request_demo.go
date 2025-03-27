/**
 * 钱包绑卡充值下单 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2WalletTradeRechargeCardRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2WalletTradeRechargeCardRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2WalletTradeRechargeCardRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000003100615",
        // 钱包用户ID
        UserHuifuId:"6666000108133109",
        // 订单金额
        TransAmt:"10.00",
        // 微信充值信息微信充值必填
        // WxRechareInfo:get5aab43a511384703Aee9A17b6d476396(),
        // 支付宝充值信息支付宝充值必填
        // AlipayRechargeInfo:get1a9a249fFf56429c8a73E7a93ef80c2b(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2WalletTradeRechargeCardRequest(dgReq)
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
    // 订单失效时间
    extendInfoMap["time_expire"] = ""
    // 备注
    extendInfoMap["remark"] = "remark11"
    // 充值方式
    extendInfoMap["recharge_type"] = "A01"
    // 绑定卡充值信息
    extendInfoMap["bank_recharge_info"] = get0994fa46Ea53463295f1Db4d35a9e9ed()
    // 异步通知地址
    extendInfoMap["notify_url"] = "archer://C_TOPAT_NOTIFY"
    return extendInfoMap
}

func get718333af8b7c4fffAe5a25a0c20f8d7e() string {
    dto := make(map[string]interface{})
    // 业务种类编码
    dto["biz_tp"] = "100099"
    // 商品简称
    dto["goods_short_name"] = "个人电脑"
    // 支付页面类型
    dto["gw_chnnl_tp"] = "02"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getCecd9a03C90045d6881d47baf8ff4e54() string {
    dto := make(map[string]interface{})
    // 交易设备类型
    dto["device_type"] = "3"
    // 交易设备IP
    dto["device_ip"] = "172.31.31.145"
    // 交易设备MAC
    dto["device_mac"] = "F0E1D2C3B4A5"
    // 交易终端设备IMEI
    dto["device_imei"] = "460030912121001"
    // 交易设备IMSI
    dto["device_imsi"] = "460030912121001"
    // 交易设备ICCID
    dto["device_icc_id"] = "898600680113F0123014"
    // 交易设备WIFI MAC
    dto["device_wifi_mac"] = "968778695A4B"
    // 交易设备经纬度
    dto["device_gps"] = "20.346790,-4.654321"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getA619b5347d694307A1a8C1ef90627a60() string {
    dto := make(map[string]interface{})
    // 经度
    dto["longitude"] = "2"
    // 纬度
    dto["latitude"] = "1"
    // 基站地址
    dto["base_station"] = "460001039217563"
    // IP地址
    dto["ip_addr"] = "172.28.52.52"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get0994fa46Ea53463295f1Db4d35a9e9ed() string {
    dto := make(map[string]interface{})
    // 银行卡序列号
    dto["token_no"] = "10000136685"
    // 跳转地址
    dto["front_url"] = "http://www.baidu.com"
    // 网联扩展数据
    dto["extend_pay_data"] = get718333af8b7c4fffAe5a25a0c20f8d7e()
    // 设备信息
    dto["terminal_device_data"] = getCecd9a03C90045d6881d47baf8ff4e54()
    // 安全信息
    dto["risk_check_data"] = getA619b5347d694307A1a8C1ef90627a60()
    // 密码页面类型
    dto["request_type"] = "M"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get5aab43a511384703Aee9A17b6d476396() string {
    dto := make(map[string]interface{})
    // 交易类型
    // dto["trade_type"] = "test"
    // 子商户公众账号ID
    // dto["sub_appid"] = ""
    // 用户标识
    // dto["openid"] = ""
    // 子商户用户标识
    // dto["sub_openid"] = ""
    // 渠道号
    // dto["channel_no"] = ""
    // 场景类型
    // dto["pay_scene"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get1a9a249fFf56429c8a73E7a93ef80c2b() string {
    dto := make(map[string]interface{})
    // 交易类型
    // dto["trade_type"] = "test"
    // 买家的支付宝唯一用户号
    // dto["buyer_id"] = "test"
    // 支付宝的店铺编号
    // dto["alipay_store_id"] = ""
    // 渠道号
    // dto["channel_no"] = ""
    // 场景类型
    // dto["pay_scene"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

