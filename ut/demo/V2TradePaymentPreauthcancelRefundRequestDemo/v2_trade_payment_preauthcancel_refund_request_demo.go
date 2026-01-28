/**
 * 微信支付宝预授权撤销 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePaymentPreauthcancelRefundRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePaymentPreauthcancelRefundRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePaymentPreauthcancelRefundRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 客户号
        HuifuId:"6666000108854952",
        // 原交易请求日期
        OrgReqDate:"20221031",
        // 撤销金额
        OrdAmt:"0.02",
        // 风控信息
        RiskCheckInfo:get9a82ccbd1b2c4e84A88aEb45c5ae6519(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradePaymentPreauthcancelRefundRequest(dgReq)
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
    extendInfoMap["out_ord_Id"] = ""
    // 原授权号
    extendInfoMap["org_auth_no"] = ""
    // 原交易请求流水号
    extendInfoMap["org_req_seq_id"] = ""
    // 原预授权全局流水号
    extendInfoMap["pre_auth_hf_seq_id"] = "0029000topB221031162644P959c0a8305400000"
    // 批次号
    extendInfoMap["batch_id"] = ""
    // 商品描述
    extendInfoMap["good_desc"] = ""
    // 备注
    extendInfoMap["remark"] = ""
    // 交易发起时间
    extendInfoMap["send_time"] = ""
    // 是否人工介入
    extendInfoMap["is_manual_process"] = "Y"
    // 汇付机具号
    extendInfoMap["devs_id"] = "SPINTP366020000360401"
    // 商户操作员号
    // extendInfoMap["mer_oper_id"] = ""
    // 扩展域
    // extendInfoMap["mer_priv"] = ""
    // 设备信息
    extendInfoMap["terminal_device_info"] = get35049a53934b425f87a06f9e3d91d5d5()
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.baidu.com"
    return extendInfoMap
}

func get9a82ccbd1b2c4e84A88aEb45c5ae6519() string {
    dto := make(map[string]interface{})
    // 基站地址
    dto["base_station"] = "192.168.1.1"
    // ip地址
    dto["ip_addr"] = "192.168.1.1"
    // 纬度
    dto["latitude"] = "33.3"
    // 经度
    dto["longitude"] = "33.3"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get35049a53934b425f87a06f9e3d91d5d5() string {
    dto := make(map[string]interface{})
    // 交易设备GPS
    dto["device_gps"] = "192.168.0.0"
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
    dto["devs_id"] = "SPINTP366061800405501"
    // 操作员号
    dto["mer_oper_id"] = ""
    // 逻辑终端号
    dto["pnr_dev_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

