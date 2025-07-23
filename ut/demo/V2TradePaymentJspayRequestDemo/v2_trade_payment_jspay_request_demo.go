/**
 * 应用场景 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePaymentJspayRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePaymentJspayRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePaymentJspayRequest{
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
    resp, err := dgSDK.V2TradePaymentJspayRequest(dgReq)
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
    extendInfoMap["wx_data"] = getB0d214acFc1a4c4b8622Afb76b104884()
    // 支付宝参数集合
    extendInfoMap["alipay_data"] = get09e7ceacF4324244A607294193546d43()
    // 银联参数集合
    extendInfoMap["unionpay_data"] = getDfee597c006f47c4A8cfF6e14a275e25()
    // 数字人民币参数集合
    // extendInfoMap["dc_data"] = get0f335b50E40d4bb4A6221be4ff09a58f()
    // 是否延迟交易
    extendInfoMap["delay_acct_flag"] = "N"
    // 手续费扣款标志
    // extendInfoMap["fee_flag"] = ""
    // 分账对象
    extendInfoMap["acct_split_bunch"] = get84c2b24cE1ae49159e2dF71a53affccf()
    // 传入分账遇到优惠的处理规则
    extendInfoMap["term_div_coupon_type"] = "0"
    // 补贴支付信息
    // extendInfoMap["combinedpay_data"] = getEb4cdb8bC9f54609Ab9aF0e76a92714f()
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
    extendInfoMap["risk_check_data"] = get3e7fb662113b4e729d8f45c0fd4a7d56()
    // 设备信息
    extendInfoMap["terminal_device_data"] = get8bf3c08e2f5d42729c86A001a763012e()
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.baidu.com"
    // 手续费补贴信息
    // extendInfoMap["trans_fee_allowance_info"] = get4e009335Db3041baB0c07d67fb26b98f()
    return extendInfoMap
}

func getD756780c707240fcB7c82e5f66908987() interface{} {
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

func get66b51648E4474320Bb51Dfbbbedef6df() interface{} {
    dto := make(map[string]interface{})
    // 单品列表
    dto["goods_detail"] = getD756780c707240fcB7c82e5f66908987()
    // 订单原价(元)
    dto["cost_price"] = "43.00"
    // 商品小票ID
    dto["receipt_id"] = "20220628132043853798"

    return dto;
}

func getDf44d149730b4956A33b6d295dcbe68a() interface{} {
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

func getA8ce15ddD1b2466c989dCb3ea5a67156() interface{} {
    dto := make(map[string]interface{})
    // 门店信息
    // dto["store_info"] = getDf44d149730b4956A33b6d295dcbe68a()

    return dto;
}

func getB0d214acFc1a4c4b8622Afb76b104884() string {
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
    dto["detail"] = get66b51648E4474320Bb51Dfbbbedef6df()
    // 设备号
    // dto["device_info"] = ""
    // 订单优惠标记
    // dto["goods_tag"] = ""
    // 实名支付
    // dto["identity"] = ""
    // 开发票入口开放标识
    // dto["receipt"] = ""
    // 场景信息
    dto["scene_info"] = getA8ce15ddD1b2466c989dCb3ea5a67156()
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

func getD51b32a5361a40a9B45563cae70cfaa3() interface{} {
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

func get3de19a00Ae944bbf860fC41b28921756() interface{} {
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

func getF1ff50db26294c9e8cac01164bdcb29f() interface{} {
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

func get09e7ceacF4324244A607294193546d43() string {
    dto := make(map[string]interface{})
    // 支付宝的店铺编号
    dto["alipay_store_id"] = ""
    // 买家的支付宝唯一用户号
    dto["buyer_id"] = "208870269990XXXX"
    // 买家支付宝账号
    dto["buyer_logon_id"] = "string"
    // 业务扩展参数
    dto["extend_params"] = getD51b32a5361a40a9B45563cae70cfaa3()
    // 订单包含的商品列表信息
    dto["goods_detail"] = get3de19a00Ae944bbf860fC41b28921756()
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
    // dto["ext_user_info"] = getF1ff50db26294c9e8cac01164bdcb29f()
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

func get001d01e99c2541b9A44e5b47e81850f2() interface{} {
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

func getDfee597c006f47c4A8cfF6e14a275e25() string {
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
    // dto["payee_info"] = get001d01e99c2541b9A44e5b47e81850f2()
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

func get0f335b50E40d4bb4A6221be4ff09a58f() string {
    dto := make(map[string]interface{})
    // 数字货币银行编号
    // dto["digital_bank_no"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getBf0c7c93805e4ea2865503c7572b96b7() interface{} {
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

func get84c2b24cE1ae49159e2dF71a53affccf() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = getBf0c7c93805e4ea2865503c7572b96b7()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getEb4cdb8bC9f54609Ab9aF0e76a92714f() string {
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

func get3e7fb662113b4e729d8f45c0fd4a7d56() string {
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

func get8bf3c08e2f5d42729c86A001a763012e() string {
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

func get4e009335Db3041baB0c07d67fb26b98f() string {
    dto := make(map[string]interface{})
    // 补贴手续费金额
    // dto["allowance_fee_amt"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

