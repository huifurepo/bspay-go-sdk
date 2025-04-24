/**
 * 分期扣款 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePayafteruseInstallmentPayRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePayafteruseInstallmentPayRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePayafteruseInstallmentPayRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 客户号
        HuifuId:"6666000108281250",
        // 交易金额
        TransAmt:"0.01",
        // 商品描述
        GoodsDesc:"聚合反扫消费",
        // 风控信息
        RiskCheckData:get5f3a85d0F5914afaAcacBad2c1d9d0c1(),
        // 支付宝扩展参数集合
        AlipayData:getDc737f3f10584db087acA3bf4f00e91f(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradePayafteruseInstallmentPayRequest(dgReq)
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
    // 入账账户号
    extendInfoMap["acct_id"] = "C02682415"
    // 交易备注
    extendInfoMap["remark"] = ""
    // 是否延迟交易
    // extendInfoMap["delay_acct_flag"] = ""
    // 分账串
    // extendInfoMap["acct_split_bunch"] = getE3ae9a478aec40059dbc91aa66aa0d21()
    // 设备信息
    extendInfoMap["terminal_device_info"] = getF7961ae67c874fb28754F33b592b7720()
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.baidu.com"
    return extendInfoMap
}

func get160495108cc94a02B915A460403ce173() interface{} {
    dto := make(map[string]interface{})
    // 分账接收方商户号
    // dto["huifu_id"] = "test"
    // 分账金额
    // dto["div_amt"] = ""
    // 分账接收方账户号
    // dto["acct_id"] = ""
    // 分账百分比%
    // dto["percentage_div"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getE3ae9a478aec40059dbc91aa66aa0d21() string {
    dto := make(map[string]interface{})
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""
    // 分账明细
    // dto["acct_infos"] = get160495108cc94a02B915A460403ce173()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get5f3a85d0F5914afaAcacBad2c1d9d0c1() string {
    dto := make(map[string]interface{})
    // IP地址
    // dto["ip_addr"] = "test"
    // 经度
    dto["longitude"] = "1"
    // 纬度
    dto["latitude"] = "2"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getF7961ae67c874fb28754F33b592b7720() string {
    dto := make(map[string]interface{})
    // 商户设备类型
    dto["mer_device_type"] = "01"
    // 商户终端实时经纬度信息
    dto["location"] = "+32.10520/+118.80593"
    // 商户交易设备IP
    dto["mer_device_ip"] = "10.11.12.13"
    // 设备类型
    dto["device_type"] = "1"
    // 交易设备IP
    dto["device_ip"] = "10.10.0.1"
    // 交易设备MAC
    dto["device_mac"] = "F0E1D2C3B4A5"
    // 交易设备GPS
    dto["device_gps"] = "20.346790"
    // 交易设备IMEI
    dto["device_imei"] = "660035730311000126101"
    // 交易设备IMSI
    dto["device_imsi"] = "460030912121001"
    // 交易设备ICCID
    dto["device_icc_id"] = "660035730311000126101"
    // 交易设备WIFIMAC
    dto["device_wifi_mac"] = "968778695A4B"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get8b2931c7A3cf4d4698bfEd6b6cb91b1f() interface{} {
    dto := make(map[string]interface{})
    // 商品的编号
    // dto["goods_id"] = "test"
    // 商品名称
    // dto["goods_name"] = "test"
    // 商品数量
    // dto["quantity"] = "test"
    // 商品单价
    // dto["price"] = "test"
    // 商品类目
    // dto["goods_category"] = ""
    // 商品描述信息
    // dto["body"] = ""
    // 商品的展示地址
    // dto["show_url"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get4449220bC6ae41b49391E552f87d642b() interface{} {
    dto := make(map[string]interface{})
    // 业务主单号
    dto["trade_component_order_id"] = "2024101001502300000002570023887054"
    // 分期履约单号
    dto["tc_installment_order_id"] = "2024101001502301900002570004797346"
    // 系统商编号
    // dto["sys_service_provider_id"] = ""

    return dto;
}

func getDc737f3f10584db087acA3bf4f00e91f() string {
    dto := make(map[string]interface{})
    // 业务扩展参数
    dto["extend_params"] = get4449220bC6ae41b49391E552f87d642b()
    // 支付宝的店铺编号
    dto["alipay_store_id"] = ""
    // 商户操作员编号
    // dto["operator_id"] = ""
    // 商户业务信息
    // dto["ali_business_params"] = ""
    // 订单包含的商品列表信息
    // dto["goods_detail"] = get8b2931c7A3cf4d4698bfEd6b6cb91b1f()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

