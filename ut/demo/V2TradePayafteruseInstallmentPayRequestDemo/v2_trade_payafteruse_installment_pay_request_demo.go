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
        RiskCheckData:get217763f244c74d4f9f6e0976741d47af(),
        // 支付宝扩展参数集合
        AlipayData:getB1c390dfB189490aAeb41d9cf8494335(),
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
    // extendInfoMap["acct_split_bunch"] = get4c088b857a3c4ba38a93A3e19ae54eed()
    // 设备信息
    extendInfoMap["terminal_device_info"] = getC9018f09676e4a9d8ff20e82c776a8d2()
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.baidu.com"
    return extendInfoMap
}

func getBb0eaeb0451a4e16Ba1095925f7b0cc5() interface{} {
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

func get4c088b857a3c4ba38a93A3e19ae54eed() string {
    dto := make(map[string]interface{})
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""
    // 分账明细
    // dto["acct_infos"] = getBb0eaeb0451a4e16Ba1095925f7b0cc5()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get217763f244c74d4f9f6e0976741d47af() string {
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

func getC9018f09676e4a9d8ff20e82c776a8d2() string {
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

func getD00a3da6E2774af9A886Ce75da002699() interface{} {
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

func get7450295cC01444418648417b36ccfa18() string {
    dto := make(map[string]interface{})
    // 业务主单号
    dto["trade_component_order_id"] = "2024101001502300000002570023887054"
    // 分期履约单号
    dto["tc_installment_order_id"] = "2024101001502301900002570004797346"
    // 系统商编号
    // dto["sys_service_provider_id"] = ""
    // 行业场景类型
    // dto["business_type"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getB1c390dfB189490aAeb41d9cf8494335() string {
    dto := make(map[string]interface{})
    // 业务扩展参数
    dto["extend_params"] = get7450295cC01444418648417b36ccfa18()
    // 支付宝的店铺编号
    dto["alipay_store_id"] = ""
    // 商户操作员编号
    // dto["operator_id"] = ""
    // 商户业务信息
    // dto["ali_business_params"] = ""
    // 订单包含的商品列表信息
    // dto["goods_detail"] = getD00a3da6E2774af9A886Ce75da002699()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

