/**
 * 应用场景 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V3TradePaymentJspayRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V3TradePaymentJspayRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V3TradePaymentJspayRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000109133323",
        // 商品描述
        GoodsDesc:"hibs自动化-通用版验证",
        // 交易类型
        TradeType:"A_NATIVE",
        // 交易金额
        TransAmt:"0.10",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V3TradePaymentJspayRequest(dgReq)
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
    // 账户号
    // extendInfoMap["acct_id"] = ""
    // 交易有效期
    extendInfoMap["time_expire"] = "20250518235959"
    // 微信参数集合
    extendInfoMap["wx_data"] = get34dff1117b3d4b27Bb3a3008d1385a3b()
    // 支付宝参数集合
    extendInfoMap["alipay_data"] = get28071c3dCc3a4138Bb93B53b837d1cdf()
    // 银联参数集合
    extendInfoMap["unionpay_data"] = getEa0b07b621664848Be1e04aaca89faac()
    // 数字人民币参数集合
    // extendInfoMap["dc_data"] = get2e3e0476D8db4f9581a0Aed52599fcf0()
    // 是否延迟交易
    extendInfoMap["delay_acct_flag"] = "N"
    // 手续费扣款标志
    // extendInfoMap["fee_flag"] = ""
    // 分账对象
    extendInfoMap["acct_split_bunch"] = getFf4d5a9c4eaa4b21A7506e25c22c0d53()
    // 传入分账遇到优惠的处理规则
    extendInfoMap["term_div_coupon_type"] = "0"
    // 补贴支付信息
    // extendInfoMap["combinedpay_data"] = getC26c8071286f429eB3aeC03ea7829128()
    // 禁用信用卡标记
    extendInfoMap["limit_pay_type"] = "NO_CREDIT"
    // 商户贴息标记
    extendInfoMap["fq_mer_discount_flag"] = "N"
    // 渠道号
    extendInfoMap["channel_no"] = ""
    // 场景类型
    extendInfoMap["pay_scene"] = "02"
    // 备注
    extendInfoMap["remark"] = "string"
    // 安全信息
    extendInfoMap["risk_check_data"] = get4eb45e7eD9844eac90378cd4e494a1f0()
    // 设备信息
    extendInfoMap["terminal_device_data"] = get09c5df58509a4737941606d5a11a639e()
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.baidu.com"
    // 手续费补贴信息
    // extendInfoMap["trans_fee_allowance_info"] = getC805c7554acd47419b994955bef7d73f()
    return extendInfoMap
}

func get990d443c980a428e805bEf11f3022ab6() interface{} {
    dto := make(map[string]interface{})
    // 商品编码
    dto["goods_id"] = "6934572310301"
    // 商品名称
    dto["goods_name"] = "太龙双黄连口服液"
    // 商品单价(元)
    dto["price"] = "43.00"
    // 商品数量
    dto["quantity"] = "1"
    // 微信侧商品编码
    dto["wxpay_goods_id"] = "12235413214070356458058"

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getA413a09d56b3469cB0ec46dda8bd1aa8() interface{} {
    dto := make(map[string]interface{})
    // 单品列表
    dto["goods_detail"] = get990d443c980a428e805bEf11f3022ab6()
    // 订单原价(元)
    dto["cost_price"] = "43.00"
    // 商品小票ID
    dto["receipt_id"] = "20220628132043853798"

    return dto;
}

func get2812fc764a1c43bbB5a711e1eb2ec245() interface{} {
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

func get6427e4d1C5b04033A009Ed75137a47c1() interface{} {
    dto := make(map[string]interface{})
    // 门店信息
    // dto["store_info"] = get2812fc764a1c43bbB5a711e1eb2ec245()

    return dto;
}

func get34dff1117b3d4b27Bb3a3008d1385a3b() string {
    dto := make(map[string]interface{})
    // 子商户应用ID
    dto["sub_appid"] = "wxdfe9a5d141f96685"
    // 子商户用户标识
    dto["sub_openid"] = "o8jhotzittQSetZ-N0Yj4Hz91Rqc"
    // 附加数据
    // dto["attach"] = ""
    // 商品描述
    // dto["body"] = ""
    // 商品详情
    dto["detail"] = getA413a09d56b3469cB0ec46dda8bd1aa8()
    // 设备号
    // dto["device_info"] = ""
    // 订单优惠标记
    // dto["goods_tag"] = ""
    // 实名支付
    // dto["identity"] = ""
    // 开发票入口开放标识
    // dto["receipt"] = ""
    // 场景信息
    dto["scene_info"] = get6427e4d1C5b04033A009Ed75137a47c1()
    // 终端ip
    // dto["spbill_create_ip"] = ""
    // 单品优惠标识
    // dto["promotion_flag"] = ""
    // 新增商品ID
    // dto["product_id"] = ""
    // 指定支付者
    // dto["limit_payer"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getAdb91d85B07f4236Ace45638f9b662b8() interface{} {
    dto := make(map[string]interface{})
    // 卡类型
    dto["card_type"] = ""
    // 支付宝点餐场景类型
    dto["food_order_type"] = "qr_order"
    // 花呗分期数
    dto["hb_fq_num"] = ""
    // 花呗卖家手续费百分比
    dto["hb_fq_seller_percent"] = ""
    // 行业数据回流信息
    dto["industry_reflux_info"] = "string"
    // 信用卡分期资产方式
    // dto["fq_channels"] = ""
    // 停车场id
    dto["parking_id"] = "123wsx"
    // 系统商编号
    dto["sys_service_provider_id"] = "1111111"

    return dto;
}

func get2bcdeef090ce43a0Ae4eAc0081d748ea() interface{} {
    dto := make(map[string]interface{})
    // 商品的编号
    dto["goods_id"] = "12312321"
    // 商品名称
    dto["goods_name"] = "汇付"
    // 商品单价(元)
    dto["price"] = "43.00"
    // 商品数量
    dto["quantity"] = "20"
    // 商品描述信息
    dto["body"] = ""
    // 商品类目树
    dto["categories_tree"] = "string"
    // 商品的展示地址
    dto["show_url"] = ""
    // 商品类目
    dto["goods_category"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get6b063f1564c248ecA7c7063379f288a6() interface{} {
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

func get28071c3dCc3a4138Bb93B53b837d1cdf() string {
    dto := make(map[string]interface{})
    // 买家的支付宝唯一用户号
    dto["buyer_id"] = "2088702699908257"
    // 支付宝的店铺编号
    dto["alipay_store_id"] = ""
    // 买家支付宝账号
    dto["buyer_logon_id"] = "string"
    // 业务扩展参数
    dto["extend_params"] = getAdb91d85B07f4236Ace45638f9b662b8()
    // 订单包含的商品列表信息
    dto["goods_detail"] = get2bcdeef090ce43a0Ae4eAc0081d748ea()
    // 商户原始订单号
    dto["merchant_order_no"] = "string"
    // 商户操作员编号
    dto["operator_id"] = "123213213"
    // 销售产品码
    dto["product_code"] = "string"
    // 卖家支付宝用户号
    dto["seller_id"] = "string"
    // 商户门店编号
    dto["store_id"] = ""
    // 外部指定买家
    // dto["ext_user_info"] = get6b063f1564c248ecA7c7063379f288a6()
    // 订单标题
    // dto["subject"] = ""
    // 商家门店名称
    // dto["store_name"] = ""
    // 小程序应用的appid
    // dto["op_app_id"] = ""
    // 商户业务信息
    // dto["ali_business_params"] = ""
    // 订单描述
    // dto["body"] = ""
    // 优惠明细参数
    // dto["ali_promo_params"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get5a98aa4f669240eeB5c328ec1d900bd5() interface{} {
    dto := make(map[string]interface{})
    // 商户类别
    // dto["mer_cat_code"] = ""
    // 二级商户代码
    // dto["sub_id"] = ""
    // 二级商户名称
    // dto["sub_name"] = ""
    // 终端号
    // dto["term_id"] = ""

    return dto;
}

func getEa0b07b621664848Be1e04aaca89faac() string {
    dto := make(map[string]interface{})
    // 二维码
    // dto["qr_code"] = ""
    // 收款方附加数据
    // dto["addn_data"] = ""
    // 地区信息
    // dto["area_info"] = ""
    // 持卡人ip
    // dto["customer_ip"] = ""
    // 前台通知地址
    // dto["front_url"] = ""
    // 订单描述
    // dto["order_desc"] = ""
    // 收款方附言
    // dto["payee_comments"] = ""
    // 收款方信息
    // dto["payee_info"] = get5a98aa4f669240eeB5c328ec1d900bd5()
    // 银联分配的服务商机构标识码
    // dto["pnr_ins_id_cd"] = ""
    // 请求方自定义域
    // dto["req_reserved"] = ""
    // 终端信息
    // dto["term_info"] = ""
    // 银联用户标识
    // dto["user_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get2e3e0476D8db4f9581a0Aed52599fcf0() string {
    dto := make(map[string]interface{})
    // 数字货币银行编号
    // dto["digital_bank_no"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get640c6e89E59a4de5Ad6fEa579780754c() interface{} {
    dto := make(map[string]interface{})
    // 分账接收方ID
    dto["huifu_id"] = "6666000109133323"
    // 分账金额
    dto["div_amt"] = "0.10"
    // 账户号
    // dto["acct_id"] = ""
    // 分账百分比%
    // dto["percentage_div"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getFf4d5a9c4eaa4b21A7506e25c22c0d53() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = get640c6e89E59a4de5Ad6fEa579780754c()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getC26c8071286f429eB3aeC03ea7829128() string {
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

func get4eb45e7eD9844eac90378cd4e494a1f0() string {
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

func get09c5df58509a4737941606d5a11a639e() string {
    dto := make(map[string]interface{})
    // 商户设备类型
    // dto["mer_device_type"] = "test"
    // 汇付机具号
    dto["devs_id"] = "SPINTP357338300264411"
    // 设备类型
    dto["device_type"] = "1"
    // 交易设备IP
    dto["device_ip"] = "10.10.0.1"
    // 交易设备MAC
    dto["device_mac"] = ""
    // 交易设备IMEI
    dto["device_imei"] = ""
    // 交易设备IMSI
    dto["device_imsi"] = ""
    // 交易设备ICCID
    dto["device_icc_id"] = ""
    // 交易设备WIFIMAC
    dto["device_wifi_mac"] = ""
    // 交易设备GPS
    dto["device_gps"] = "192.168.0.0"
    // 商户终端应用程序版本
    // dto["app_version"] = ""
    // SIM卡卡号
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
    // 商户终端序列号
    // dto["serial_num"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getC805c7554acd47419b994955bef7d73f() string {
    dto := make(map[string]interface{})
    // 补贴手续费金额
    // dto["allowance_fee_amt"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

