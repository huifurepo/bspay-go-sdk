/**
 * 聚合反扫接口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePaymentMicropayRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePaymentMicropayRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePaymentMicropayRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000018328947",
        // 交易金额
        TransAmt:"0.01",
        // 商品描述
        GoodsDesc:"聚合反扫消费",
        // 支付授权码
        AuthCode:"2884138408701518074",
        // 安全信息
        RiskCheckData:getRiskCheckData(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradePaymentMicropayRequest(dgReq)
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
    // 交易有效期
    extendInfoMap["time_expire"] = "20220918150330"
    // 手续费扣款标志
    // extendInfoMap["fee_flag"] = ""
    // 禁用信用卡标记
    extendInfoMap["limit_pay_type"] = ""
    // 是否延迟交易
    extendInfoMap["delay_acct_flag"] = "Y"
    // 渠道号
    extendInfoMap["channel_no"] = ""
    // 营销补贴信息
    // extendInfoMap["combinedpay_data"] = getCombinedpayData()
    // 场景类型
    extendInfoMap["pay_scene"] = ""
    // 分账对象
    // extendInfoMap["acct_split_bunch"] = getAcctSplitBunch()
    // 传入分帐遇到优惠的处理规则
    extendInfoMap["term_div_coupon_type"] = "3"
    // 聚合反扫微信参数集合
    // extendInfoMap["wx_data"] = getWxData()
    // 支付宝扩展参数集合
    extendInfoMap["alipay_data"] = getAlipayData()
    // 银联参数集合
    // extendInfoMap["unionpay_data"] = getUnionpayData()
    // 设备信息
    extendInfoMap["terminal_device_info"] = getTerminalDeviceInfo()
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.baidu.com"
    // 交易备注
    extendInfoMap["remark"] = ""
    // 账户号
    // extendInfoMap["acct_id"] = ""
    return extendInfoMap
}

func getGoodsDetail() interface{} {
    dto := make(map[string]interface{})
    // 商品的编号
    // dto["goods_id"] = "test"
    // 商品名称
    // dto["goods_name"] = "test"
    // 商品数量
    // dto["quantity"] = "test"
    // 商品单价
    // dto["price"] = "test"
    // 商品类目树
    // dto["categories_tree"] = ""
    // 商品类目
    // dto["goods_category"] = ""
    // 商品描述信息
    // dto["body"] = ""
    // 商品的展示地址
    // dto["show_url"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getStoreInfo() interface{} {
    dto := make(map[string]interface{})
    // 门店id
    // dto["id"] = ""
    // 门店名称
    // dto["name"] = ""
    // 门店行政区划码
    // dto["area_code"] = ""
    // 门店详细地址
    // dto["address"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
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

func getAcctInfosRucan() interface{} {
    dto := make(map[string]interface{})
    // 分账金额
    // dto["div_amt"] = "test"
    // 被分账方ID
    // dto["huifu_id"] = "test"
    // 账户号
    // dto["acct_id"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getTerminalDeviceInfo() string {
    dto := make(map[string]interface{})
    // 商户设备类型
    // dto["mer_device_type"] = "test"
    // 汇付机具号
    // dto["devs_id"] = "test"
    // 设备类型
    dto["device_type"] = "4"
    // 交易设备IP
    dto["device_ip"] = "10.10.0.1"
    // 交易设备MAC
    dto["device_mac"] = "030147441006000182623"
    // 交易设备IMEI
    dto["device_imei"] = "030147441006000182623"
    // 交易设备IMSI
    dto["device_imsi"] = "030147441006000182623"
    // 交易设备ICCID
    dto["device_icc_id"] = "030147441006000182623"
    // 交易设备WIFIMAC
    dto["device_wifi_mac"] = "030147441006000182623"
    // 交易设备GPS
    dto["device_gps"] = "111"
    // 商户终端应用程序版
    // dto["app_version"] = ""
    // 加密随机因子
    // dto["encrypt_rand_num"] = ""
    // SIM 卡卡号
    // dto["icc_id"] = ""
    // 商户终端实时经纬度信息
    // dto["location"] = ""
    // 商户交易设备IP
    // dto["mer_device_ip"] = ""
    // 移动国家代码
    // dto["mobile_country_cd"] = ""
    // 移动网络号码
    // dto["mobile_net_num"] = ""
    // 商户终端入网认证编号
    // dto["network_license"] = ""
    // 密文数据
    // dto["secret_text"] = ""
    // 商户终端序列号
    // dto["serial_num"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getAcctSplitBunch() string {
    dto := make(map[string]interface{})
    // 分账明细
    // dto["acct_infos"] = getAcctInfosRucan()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getRiskCheckData() string {
    dto := make(map[string]interface{})
    // ip地址
    // dto["ip_addr"] = ""
    // 基站地址
    // dto["base_atation"] = ""
    // 纬度
    dto["latitude"] = "2"
    // 经度
    dto["longitude"] = "1"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getUnionpayData() string {
    dto := make(map[string]interface{})
    // 币种
    // dto["currency_code"] = ""
    // 支持发票
    // dto["invoice_st"] = ""
    // 商户类别
    // dto["mer_cat_code"] = ""
    // 银联参数集合
    // dto["pnrins_id_cd"] = ""
    // 特殊计费信息
    // dto["specfeeinfo"] = ""
    // 终端号
    // dto["term_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getSceneInfo() interface{} {
    dto := make(map[string]interface{})
    // 门店信息
    // dto["store_info"] = getStoreInfo()

    return dto;
}

func getWxData() string {
    dto := make(map[string]interface{})
    // 子商户公众账号id
    // dto["sub_appid"] = ""
    // 用户标识
    // dto["openid"] = ""
    // 子商户用户标识
    // dto["sub_openid"] = ""
    // 设备号
    // dto["device_info"] = ""
    // 附加数据
    // dto["attach"] = ""
    // 商品详情
    // dto["detail"] = getDetail()
    // 场景信息
    // dto["scene_info"] = getSceneInfo()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getExtendParams() interface{} {
    dto := make(map[string]interface{})
    // 卡类型
    // dto["card_type"] = ""
    // 支付宝点餐场景类型
    // dto["food_order_type"] = ""
    // 花呗分期数
    // dto["hb_fq_num"] = ""
    // 花呗卖家承担的手续费百分比
    // dto["hb_fq_seller_percent"] = ""
    // 行业数据回流信息
    // dto["industry_reflux_info"] = ""
    // 停车场id
    // dto["parking_id"] = ""
    // 系统商编号
    // dto["sys_service_provider_id"] = ""

    return dto;
}

func getAlipayData() string {
    dto := make(map[string]interface{})
    // 支付宝的店铺编号
    dto["alipay_store_id"] = ""
    // 订单包含的商品列表信息
    // dto["goods_detail"] = getGoodsDetail()
    // 业务扩展参数
    // dto["extend_params"] = getExtendParams()
    // 商户操作员编号
    // dto["operator_id"] = ""
    // 商户门店编号
    dto["store_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getDetail() interface{} {
    dto := make(map[string]interface{})
    // 单品列表
    // dto["goods_detail"] = getGoodsDetailWxRucan()
    // 订单原价
    // dto["cost_price"] = ""
    // 商品小票ID
    // dto["receipt_id"] = ""

    return dto;
}

func getGoodsDetailWxRucan() interface{} {
    dto := make(map[string]interface{})
    // 商品编码
    // dto["goods_id"] = ""
    // 商品名称
    // dto["goods_name"] = ""
    // 商品单价
    // dto["price"] = ""
    // 商品数量
    // dto["quantity"] = ""
    // 微信侧商品编码
    // dto["wxpay_goods_id"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

