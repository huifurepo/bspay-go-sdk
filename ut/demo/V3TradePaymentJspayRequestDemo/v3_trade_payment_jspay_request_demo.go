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
    extendInfoMap["wx_data"] = get8f4ad72627694c18Ba86Bd9089993223()
    // 支付宝参数集合
    extendInfoMap["alipay_data"] = get6ff391b708884f29B5bb9c38dbc05eaf()
    // 银联参数集合
    extendInfoMap["unionpay_data"] = getFdd35a17B50c44bb9dae6dc0496e1c44()
    // 数字人民币参数集合
    // extendInfoMap["dc_data"] = get4e4205d2Ef6b4c88B3b39af08eb7b8e7()
    // 是否延迟交易
    extendInfoMap["delay_acct_flag"] = "N"
    // 手续费扣款标志
    // extendInfoMap["fee_flag"] = ""
    // 分账对象
    extendInfoMap["acct_split_bunch"] = get5f2b22187a1340818461B8baa5c50f4b()
    // 传入分账遇到优惠的处理规则
    extendInfoMap["term_div_coupon_type"] = "0"
    // 补贴支付信息
    // extendInfoMap["combinedpay_data"] = get9a739c81E37741baBdb45e347a0faac7()
    // 补贴支付手续费承担方信息
    // extendInfoMap["combinedpay_data_fee_info"] = getC95f39053e964866A0b32ed0912a257f()
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
    extendInfoMap["risk_check_data"] = get69fc765aD5704c85B3ea06703ac7f588()
    // 设备信息
    extendInfoMap["terminal_device_data"] = get7c420dd3D80743789b387e128045019d()
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.baidu.com"
    // 手续费补贴信息
    // extendInfoMap["trans_fee_allowance_info"] = get1a44c52b2db0437e90e089cae551d69e()
    // 手续费场景标识
    // extendInfoMap["fee_sign"] = ""
    return extendInfoMap
}

func get10b097adD5b743408ffe5a8ef0c870a3() interface{} {
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

func get4101ac5472b24fd1A43f3948d0c6409c() interface{} {
    dto := make(map[string]interface{})
    // 单品列表
    dto["goods_detail"] = get10b097adD5b743408ffe5a8ef0c870a3()
    // 订单原价(元)
    dto["cost_price"] = "43.00"
    // 商品小票ID
    dto["receipt_id"] = "20220628132043853798"

    return dto;
}

func get5f6d1a317a1049edBc256f5f7ee3f006() interface{} {
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

func get2518f582Da9747b3Ac3bCc374547d331() interface{} {
    dto := make(map[string]interface{})
    // 门店信息
    // dto["store_info"] = get5f6d1a317a1049edBc256f5f7ee3f006()

    return dto;
}

func get8f4ad72627694c18Ba86Bd9089993223() string {
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
    dto["detail"] = get4101ac5472b24fd1A43f3948d0c6409c()
    // 设备号
    // dto["device_info"] = ""
    // 订单优惠标记
    // dto["goods_tag"] = ""
    // 实名支付
    // dto["identity"] = ""
    // 开发票入口开放标识
    // dto["receipt"] = ""
    // 场景信息
    dto["scene_info"] = get2518f582Da9747b3Ac3bCc374547d331()
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

func get2d16cf16A91c46d48b34083415a8cef6() interface{} {
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

func get2cfdafba93bd41d2A00051d8605d030c() interface{} {
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

func get9d9461e41bff4302847d00f06e7ba9c9() interface{} {
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

func get6ff391b708884f29B5bb9c38dbc05eaf() string {
    dto := make(map[string]interface{})
    // 支付宝的店铺编号
    dto["alipay_store_id"] = ""
    // 买家的支付宝唯一用户号
    dto["buyer_id"] = "208870269990XXXX"
    // 买家支付宝账号
    dto["buyer_logon_id"] = "string"
    // 业务扩展参数
    dto["extend_params"] = get2d16cf16A91c46d48b34083415a8cef6()
    // 订单包含的商品列表信息
    dto["goods_detail"] = get2cfdafba93bd41d2A00051d8605d030c()
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
    // dto["ext_user_info"] = get9d9461e41bff4302847d00f06e7ba9c9()
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

func get388555dcFa0b49999b1f21e5f7824e11() interface{} {
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

func get3222dc83Df9f474189488cd1eb19fa7c() string {
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

func getFdd35a17B50c44bb9dae6dc0496e1c44() string {
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
    // dto["payee_info"] = get388555dcFa0b49999b1f21e5f7824e11()
    // 银联分配的服务商机构标识码
    // dto["pnr_ins_id_cd"] = ""
    // 请求方自定义域
    // dto["req_reserved"] = ""
    // 终端信息
    // dto["term_info"] = ""
    // 银联用户标识
    // dto["user_id"] = ""
    // 服务商信息
    // dto["pid_info"] = get3222dc83Df9f474189488cd1eb19fa7c()
    // 银联支付标识
    // dto["app_up_identifier"] = ""
    // 交易支付渠道
    // dto["pay_channel"] = ""
    // 第三方公众账号ID
    // dto["third_part_app_id"] = ""
    // 第三方用户标识
    // dto["third_part_open_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get4e4205d2Ef6b4c88B3b39af08eb7b8e7() string {
    dto := make(map[string]interface{})
    // 数字货币银行编号
    // dto["digital_bank_no"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getAbc106c0D4db4dc882d204d73b0e9c60() interface{} {
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

func get5f2b22187a1340818461B8baa5c50f4b() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = getAbc106c0D4db4dc882d204d73b0e9c60()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get9a739c81E37741baBdb45e347a0faac7() string {
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

func getC95f39053e964866A0b32ed0912a257f() string {
    dto := make(map[string]interface{})
    // 补贴支付手续费承担方汇付编号
    // dto["huifu_id"] = ""
    // 补贴支付手续费承担方账户号
    // dto["acct_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get69fc765aD5704c85B3ea06703ac7f588() string {
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

func get7c420dd3D80743789b387e128045019d() string {
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

func get1a44c52b2db0437e90e089cae551d69e() string {
    dto := make(map[string]interface{})
    // 补贴手续费金额
    // dto["allowance_fee_amt"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

