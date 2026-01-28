/**
 * 微信支付宝预授权完成 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePreauthpayRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePreauthpayRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePreauthpayRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000108854952",
        // 原交易请求日期
        OrgReqDate:"20221031",
        // 交易金额
        TransAmt:"0.02",
        // 商品描述
        GoodsDesc:"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567",
        // 安全信息
        RiskCheckData:get82649f29E7ed468bB63dCac763d0c52e(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradePreauthpayRequest(dgReq)
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
    // 外部订单号
    extendInfoMap["out_ord_id"] = "12345678901234567890123456789012"
    // 原授权号
    extendInfoMap["org_auth_no"] = ""
    // 原预授权交易请求流水号
    extendInfoMap["org_req_seq_id"] = ""
    // 预授权汇付全局流水号
    extendInfoMap["pre_auth_hf_seq_id"] = "0029000topB221031163126P798c0a8305400000"
    // 备注
    extendInfoMap["remark"] = "123451111"
    // 批次号
    // extendInfoMap["batch_id"] = ""
    // 商户操作员号
    // extendInfoMap["mer_oper_id"] = ""
    // 扩展域
    // extendInfoMap["mer_priv"] = ""
    // 设备信息
    extendInfoMap["terminal_device_data"] = get4c12547917694abaBa1cBb866c7d5361()
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.baidu.com"
    return extendInfoMap
}

func get82649f29E7ed468bB63dCac763d0c52e() string {
    dto := make(map[string]interface{})
    // 基站地址
    dto["base_station"] = "192.168.1.1"
    // ip地址
    dto["ip_addr"] = "180.167.105.130"
    // 纬度
    dto["latitude"] = "33.3"
    // 经度
    dto["longitude"] = "33.3"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get4c12547917694abaBa1cBb866c7d5361() string {
    dto := make(map[string]interface{})
    // 商户终端序列号
    dto["app_version"] = ""
    // 交易设备GPS
    dto["device_gps"] = ""
    // 交易设备ICCID
    dto["device_icc_id"] = ""
    // 交易设备IMEI
    dto["device_imei"] = ""
    // 交易设备IMSI
    dto["device_imsi"] = ""
    // 交易设备IP
    dto["device_ip"] = "10.10.0.1"
    // 交易设备MAC
    dto["device_mac"] = ""
    // 设备类型
    dto["device_type"] = "1"
    // 交易设备WIFIMAC
    dto["device_wifi_mac"] = ""
    // 汇付机具号
    dto["devs_id"] = ""
    // ICCID
    dto["icc_id"] = ""
    // 商户终端实时经纬度信息
    dto["location"] = "+32.10520/+118.80593"
    // 商户交易设备IP
    dto["mer_device_ip"] = ""
    // 商户设备类型
    dto["mer_device_type"] = "01"
    // 移动国家代码
    dto["mobile_country_cd"] = ""
    // 移动网络号码
    dto["mobile_net_num"] = ""
    // 商户终端入网认证编号
    dto["network_license"] = "P3111"
    // 商户终端序列号
    dto["serial_num"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

