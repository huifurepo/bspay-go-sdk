/**
 * 聚合反扫 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V3TradePaymentMicropayRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V3TradePaymentMicropayRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V3TradePaymentMicropayRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000109133323",
        // 交易金额
        TransAmt:"1.01",
        // 商品描述
        GoodsDesc:"聚合反扫消费",
        // 支付授权码
        AuthCode:"131135212661863252",
        // 安全信息
        RiskCheckData:get662dde05Ae07407489b4112cd0fdc2aa(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V3TradePaymentMicropayRequest(dgReq)
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
    // extendInfoMap["time_expire"] = ""
    // 手续费扣款标志
    // extendInfoMap["fee_flag"] = ""
    // 禁用支付方式
    // extendInfoMap["limit_pay_type"] = ""
    // 是否延迟交易
    // extendInfoMap["delay_acct_flag"] = ""
    // 渠道号
    // extendInfoMap["channel_no"] = ""
    // 补贴支付信息
    // extendInfoMap["combinedpay_data"] = getEe46736f77394ed094ccCd2ec3c2c656()
    // 补贴支付手续费承担方信息
    // extendInfoMap["combinedpay_data_fee_info"] = get73544a6d326b4917B1dd53b756bba2e5()
    // 场景类型
    // extendInfoMap["pay_scene"] = ""
    // 分账对象
    // extendInfoMap["acct_split_bunch"] = get89418fab7825494aB8a9515de2e6e344()
    // 传入分帐遇到优惠的处理规则
    // extendInfoMap["term_div_coupon_type"] = ""
    // 聚合反扫微信参数集合
    // extendInfoMap["wx_data"] = get08682eee3f0d470cBdb4974f20c4baab()
    // 支付宝扩展参数集合
    // extendInfoMap["alipay_data"] = get17528c7e143040a5B75585fd1a68b81a()
    // 银联参数集合
    // extendInfoMap["unionpay_data"] = get94882e5e893949ba9d00290f861b3c08()
    // 设备信息
    // extendInfoMap["terminal_device_info"] = getD4cee1857f73447393c83b7127f0025a()
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.baidu.com"
    // 交易备注
    // extendInfoMap["remark"] = ""
    // 账户号
    // extendInfoMap["acct_id"] = ""
    // 手续费补贴信息
    // extendInfoMap["trans_fee_allowance_info"] = getFefc9eeaF0224944A5dfFd07198b753f()
    // 手续费场景标识
    // extendInfoMap["fee_sign"] = ""
    return extendInfoMap
}

func getEe46736f77394ed094ccCd2ec3c2c656() string {
    dto := make(map[string]interface{})
    // 补贴方汇付商户号
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

func get73544a6d326b4917B1dd53b756bba2e5() string {
    dto := make(map[string]interface{})
    // 补贴支付手续费承担方汇付编号
    // dto["huifu_id"] = ""
    // 补贴支付手续费承担方账户号
    // dto["acct_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get40001115263c442b83f0E61df0b3536a() interface{} {
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

func get89418fab7825494aB8a9515de2e6e344() string {
    dto := make(map[string]interface{})
    // 分账明细
    // dto["acct_infos"] = get40001115263c442b83f0E61df0b3536a()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get89fb549f04044fbf895bAe6d41737b9e() interface{} {
    dto := make(map[string]interface{})
    // 商品编码
    // dto["goods_id"] = "test"
    // 商品数量
    // dto["quantity"] = "test"
    // 商品名称
    // dto["goods_name"] = ""
    // 商品单价
    // dto["price"] = ""
    // 微信侧商品编码
    // dto["wxpay_goods_id"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get88aaace5C2f24af79d10E93b063c01ce() interface{} {
    dto := make(map[string]interface{})
    // 单品列表
    // dto["goods_detail"] = get89fb549f04044fbf895bAe6d41737b9e()
    // 订单原价
    // dto["cost_price"] = ""
    // 商品小票ID
    // dto["receipt_id"] = ""

    return dto;
}

func get8b1d595d82654e91Ad0c70e1a22710d9() interface{} {
    dto := make(map[string]interface{})
    // 门店id
    // dto["id"] = ""
    // 门店名称
    // dto["name"] = ""
    // 门店行政区划码
    // dto["area_code"] = ""
    // 门店详细地址
    // dto["address"] = ""

    return dto;
}

func get1c9201dfD5b341e788fe9cf3b255b9e5() interface{} {
    dto := make(map[string]interface{})
    // 门店信息
    // dto["store_info"] = get8b1d595d82654e91Ad0c70e1a22710d9()

    return dto;
}

func get08682eee3f0d470cBdb4974f20c4baab() string {
    dto := make(map[string]interface{})
    // 收款设备IP直联模式必填字段；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：192.168.2.2&lt;/font&gt;
    // dto["spbill_create_ip"] = "test"
    // 子商户公众账号id
    // dto["sub_appid"] = ""
    // 设备号
    // dto["device_info"] = ""
    // 订单优惠标记
    // dto["goods_tag"] = ""
    // 附加数据
    // dto["attach"] = ""
    // 商品详情
    // dto["detail"] = get88aaace5C2f24af79d10E93b063c01ce()
    // 场景信息
    // dto["scene_info"] = get1c9201dfD5b341e788fe9cf3b255b9e5()
    // 单品优惠标识
    // dto["promotion_flag"] = ""
    // 电子发票入口开放标识
    // dto["receipt"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getB2817307C99242c3B3eb539ff6858d08() interface{} {
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

func get47bb962e164344caB7c759f106d93529() interface{} {
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
    // 支卡通相关参数
    // dto["preset_pay_tool"] = ""

    return dto;
}

func get713f8fec1ccc4070Bcd3A6723a0baaeb() interface{} {
    dto := make(map[string]interface{})
    // 姓名
    // dto["name"] = ""
    // 手机号
    // dto["mobile"] = ""
    // 证件类型
    // dto["cert_type"] = ""
    // 证件号
    // dto["cert_no"] = ""
    // 允许的最小买家年龄
    // dto["min_age"] = ""
    // 是否强制校验付款人身份信息
    // dto["fix_buyer"] = ""
    // 是否强制校验身份信息
    // dto["need_check_info"] = ""

    return dto;
}

func get17528c7e143040a5B75585fd1a68b81a() string {
    dto := make(map[string]interface{})
    // 支付宝的店铺编号
    // dto["alipay_store_id"] = ""
    // 订单包含的商品列表信息
    // dto["goods_detail"] = getB2817307C99242c3B3eb539ff6858d08()
    // 业务扩展参数
    // dto["extend_params"] = get47bb962e164344caB7c759f106d93529()
    // 商户操作员编号
    // dto["operator_id"] = ""
    // 商户门店编号
    // dto["store_id"] = ""
    // 外部指定买家
    // dto["ext_user_info"] = get713f8fec1ccc4070Bcd3A6723a0baaeb()
    // 商户业务信息
    // dto["ali_business_params"] = ""
    // 订单描述
    // dto["body"] = ""
    // 优惠明细参数
    // dto["ali_promo_params"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getFb0fb42a43ca42b08973Bc1fb63b99b4() string {
    dto := make(map[string]interface{})
    // 服务商订单编号
    // dto["pnr_order_id"] = ""
    // 服务商密文
    // dto["pid_sct"] = ""
    // 场景标识
    // dto["trade_scene"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get94882e5e893949ba9d00290f861b3c08() string {
    dto := make(map[string]interface{})
    // 币种
    // dto["currency_code"] = ""
    // 支持发票
    // dto["invoice_st"] = ""
    // 商户类别
    // dto["mer_cat_code"] = ""
    // 服务商机构标识码
    // dto["pnr_ins_id_cd"] = ""
    // 特殊计费信息
    // dto["specfeeinfo"] = ""
    // 终端号
    // dto["term_id"] = ""
    // 收款方附加数据
    // dto["addn_data"] = ""
    // 服务商信息
    // dto["pid_info"] = getFb0fb42a43ca42b08973Bc1fb63b99b4()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get662dde05Ae07407489b4112cd0fdc2aa() string {
    dto := make(map[string]interface{})
    // ip地址
    dto["ip_addr"] = "180.167.105.130"
    // 基站地址
    dto["base_station"] = "192.168.1.1"
    // 纬度
    dto["latitude"] = "33.3"
    // 经度
    dto["longitude"] = "33.3"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getD4cee1857f73447393c83b7127f0025a() string {
    dto := make(map[string]interface{})
    // 商户设备类型
    // dto["mer_device_type"] = "test"
    // 汇付机具号
    // dto["devs_id"] = "test"
    // 设备类型
    // dto["device_type"] = ""
    // 交易设备IP
    // dto["device_ip"] = ""
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

func getFefc9eeaF0224944A5dfFd07198b753f() string {
    dto := make(map[string]interface{})
    // 补贴手续费金额
    // dto["allowance_fee_amt"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

