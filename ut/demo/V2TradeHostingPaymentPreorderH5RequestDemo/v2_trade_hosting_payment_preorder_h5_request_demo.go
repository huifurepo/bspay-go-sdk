/**
 * H5、PC预下单接口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeHostingPaymentPreorderH5RequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeHostingPaymentPreorderH5RequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeHostingPaymentPreorderH5Request{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000109133323",
        // 交易金额
        TransAmt:"0.10",
        // 商品描述
        GoodsDesc:"个人电脑",
        // 预下单类型
        PreOrderType:"1",
        // 统一收银台扩展参数集合
        HostingData:get6f5b57c568774d0299c228e6f4091b08(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeHostingPaymentPreorderH5Request(dgReq)
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
    // 收款汇付账户号
    // extendInfoMap["acct_id"] = ""
    // 是否延迟交易
    extendInfoMap["delay_acct_flag"] = "N"
    // 是否支持切换支付方式
    // extendInfoMap["multi_pay_way_flag"] = ""
    // 分账对象
    extendInfoMap["acct_split_bunch"] = get433cd49d16be4a7eA70989990c2bd61f()
    // 交易失效时间
    // extendInfoMap["time_expire"] = ""
    // 业务信息
    extendInfoMap["biz_info"] = get4bee3e14F79f426398f14a8383e0b723()
    // 交易异步通知地址
    extendInfoMap["notify_url"] = "https://callback.service.com/xx"
    // 使用类型
    // extendInfoMap["usage_type"] = ""
    // 交易类型
    // extendInfoMap["trans_type"] = ""
    // 微信参数集合
    // extendInfoMap["wx_data"] = get0b9ec56122554c89A8406dee0d4ca398()
    // 支付宝参数集合
    // extendInfoMap["alipay_data"] = get4d35eb8132844ea4Bb28C502db6951a0()
    // 抖音参数集合
    // extendInfoMap["dy_data"] = get44c38738Bf2649b8Aa8e22e6865eec2c()
    // 银联参数集合
    // extendInfoMap["unionpay_data"] = get679d482aA4c8474fB43e42c534237b76()
    // 设备信息
    // extendInfoMap["terminal_device_data"] = get7259d42290ee414d9275411e468ce9bc()
    // 大额支付参数集合
    // extendInfoMap["largeamt_data"] = getFa26e3a2D80d410d80ad19246e8889a4()
    // 手续费场景标识
    // extendInfoMap["fee_sign"] = ""
    return extendInfoMap
}

func get418f01ceC78549519daa0ebc9927f853() interface{} {
    dto := make(map[string]interface{})
    // 分账金额
    dto["div_amt"] = "0.08"
    // 分账接收方ID
    dto["huifu_id"] = "6666000111546360"
    // 收款汇付账户号
    // dto["acct_id"] = ""
    // 分账百分比%
    // dto["percentage_div"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get433cd49d16be4a7eA70989990c2bd61f() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = get418f01ceC78549519daa0ebc9927f853()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get6f5b57c568774d0299c228e6f4091b08() string {
    dto := make(map[string]interface{})
    // 项目标题
    dto["project_title"] = "收银台标题"
    // 项目号
    dto["project_id"] = "PROJECTID2023101225142567"
    // 请求类型P:PC页面版，默认：P；M:H5页面版；指定交易类型时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：M&lt;/font&gt;
    // dto["request_type"] = "test"
    // 商户私有信息
    dto["private_info"] = "商户私有信息test"
    // 回调地址
    dto["callback_url"] = "https://paas.huifu.com"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get7c9c87a845a14c12B2a1066740ef6880() interface{} {
    dto := make(map[string]interface{})
    // 是否提供校验身份信息
    dto["need_check_info"] = "T"
    // 允许的最小买家年龄
    dto["min_age"] = "12"
    // 是否强制校验付款人身份信息
    dto["fix_buyer"] = "F"

    return dto;
}

func getF01e18fa954f4a5799f09110d2c54147() interface{} {
    dto := make(map[string]interface{})
    // 指定支付者
    dto["limit_payer"] = "ADULT"
    // 微信实名验证
    dto["real_name_flag"] = "Y"

    return dto;
}

func getCe9b512aFb8b498498c6Ba586d5f74a8() interface{} {
    dto := make(map[string]interface{})
    // 姓名
    dto["name"] = "张三"
    // 证件类型
    dto["cert_type"] = "IDENTITY_CARD"
    // 证件号
    dto["cert_no"] = "OLOxrl809XmVIMmPRTIyIpJHxi4HFMJNmxGz1nhZAtps9xPEVDysU3UZPBZfsNFLFcXEMENPsJ+tymbYt42dNQ+76hyEtx+qz0BQhU8SKV8H5itrI4kaXpaJf+sV8OewrJkhDidPdz01vqMEDQRhaMnNwl8snHZxDdpu7HVUz5JwsLYfBbZP2Q4CZpVWS3NunQahZ8zHQ+8EhvYVH1vE7f/M6nJt1/4GoHz9C/UnuYudV0S2uBhlywbX+YkRGNMl/oz5+UNeMDA2jqhtTreSD4+w7S82L53rqKsAbosodOPo4OAMZk4/0QOH5LDbqFByVM97mzHfw7z+Tx/dsXJ8QQ=="
    // 手机号
    dto["mobile"] = "15012345678"

    return dto;
}

func get4bee3e14F79f426398f14a8383e0b723() string {
    dto := make(map[string]interface{})
    // 付款人验证（支付宝）
    dto["payer_check_ali"] = get7c9c87a845a14c12B2a1066740ef6880()
    // 付款人验证（微信）
    dto["payer_check_wx"] = getF01e18fa954f4a5799f09110d2c54147()
    // 个人付款人信息
    dto["person_payer"] = getCe9b512aFb8b498498c6Ba586d5f74a8()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get27ce73b3192f4eddBabf54dec9fbd475() interface{} {
    dto := make(map[string]interface{})
    // 商品编码
    // dto["goods_id"] = ""
    // 商品名称
    // dto["goods_name"] = ""
    // 商品单价(元)
    // dto["price"] = ""
    // 商品数量
    // dto["quantity"] = ""
    // 微信侧商品编码
    // dto["wxpay_goods_id"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get94e673aa9bb84ea3999d9dca7a93b964() interface{} {
    dto := make(map[string]interface{})
    // 单品列表
    // dto["goods_detail"] = get27ce73b3192f4eddBabf54dec9fbd475()
    // 订单原价(元)
    // dto["cost_price"] = ""
    // 商品小票ID
    // dto["receipt_id"] = ""

    return dto;
}

func get5fb97a0cFa3c4cccB4cbC81e9b8143a3() interface{} {
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

func get84eb4ca8D6d64ded899f0127354634ee() interface{} {
    dto := make(map[string]interface{})
    // 门店信息
    // dto["store_info"] = get5fb97a0cFa3c4cccB4cbC81e9b8143a3()

    return dto;
}

func get0b9ec56122554c89A8406dee0d4ca398() string {
    dto := make(map[string]interface{})
    // 附加数据
    // dto["attach"] = ""
    // 商品详情
    // dto["detail"] = get94e673aa9bb84ea3999d9dca7a93b964()
    // 订单优惠标记
    // dto["goods_tag"] = ""
    // 开发票入口开放标识
    // dto["receipt"] = ""
    // 场景信息
    // dto["scene_info"] = get84eb4ca8D6d64ded899f0127354634ee()
    // 单品优惠标识
    // dto["promotion_flag"] = ""
    // 新增商品ID
    // dto["product_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get952ea60e512e4ed3Bfc9483b4166042b() interface{} {
    dto := make(map[string]interface{})
    // 卡类型
    // dto["card_type"] = ""
    // 支付宝点餐场景类型
    // dto["food_order_type"] = ""
    // 花呗分期数
    // dto["hb_fq_num"] = ""
    // 花呗卖家手续费百分比
    // dto["hb_fq_seller_percent"] = ""
    // 行业数据回流信息
    // dto["industry_reflux_info"] = ""
    // 信用卡分期资产方式
    // dto["fq_channels"] = ""
    // 停车场id
    // dto["parking_id"] = ""
    // 系统商编号
    // dto["sys_service_provider_id"] = ""

    return dto;
}

func getAd71e726Cf2c4965B057Cc4743c67c75() interface{} {
    dto := make(map[string]interface{})
    // 商品的编号
    // dto["goods_id"] = "test"
    // 商品名称
    // dto["goods_name"] = "test"
    // 商品单价(元)
    // dto["price"] = "test"
    // 商品数量
    // dto["quantity"] = "test"
    // 商品描述信息
    // dto["body"] = ""
    // 商品类目树
    // dto["categories_tree"] = ""
    // 商品类目
    // dto["goods_category"] = ""
    // 商品的展示地址
    // dto["show_url"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get4d35eb8132844ea4Bb28C502db6951a0() string {
    dto := make(map[string]interface{})
    // 支付宝的店铺编号
    // dto["alipay_store_id"] = ""
    // 业务扩展参数
    // dto["extend_params"] = get952ea60e512e4ed3Bfc9483b4166042b()
    // 订单包含的商品列表信息
    // dto["goods_detail"] = getAd71e726Cf2c4965B057Cc4743c67c75()
    // 商户原始订单号
    // dto["merchant_order_no"] = ""
    // 商户操作员编号
    // dto["operator_id"] = ""
    // 销售产品码
    // dto["product_code"] = ""
    // 卖家支付宝用户号
    // dto["seller_id"] = ""
    // 商户门店编号
    // dto["store_id"] = ""
    // 订单标题
    // dto["subject"] = ""
    // 商家门店名称
    // dto["store_name"] = ""
    // 商户业务信息
    // dto["ali_business_params"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getBfa1f6ac21c4482984bb567a87d31805() interface{} {
    dto := make(map[string]interface{})
    // 场景类型
    // dto["type"] = "test"
    // 应用名称
    // dto["app_name"] = ""
    // 网站URL
    // dto["app_url"] = ""
    // iOS平台BundleID
    // dto["bundle_id"] = ""
    // Android平台PackageName
    // dto["package_name"] = ""

    return dto;
}

func get7de2e7c3715a4a2a897e5cc6eb82d037() interface{} {
    dto := make(map[string]interface{})
    // 用户终端IP
    // dto["payer_client_ip"] = "test"

    return dto;
}

func get44c38738Bf2649b8Aa8e22e6865eec2c() string {
    dto := make(map[string]interface{})
    // 子商户应用ID
    // dto["sub_appid"] = "test"
    // H5场景信息
    // dto["h5_info"] = getBfa1f6ac21c4482984bb567a87d31805()
    // 场景信息
    // dto["scene_info"] = get7de2e7c3715a4a2a897e5cc6eb82d037()
    // 优惠标记
    // dto["coupon_info"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get6b0066137be34e0d913fA123d87b6ecd() interface{} {
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

func getD2810b7188724ceeAffdE9d1d6c9a3ff() string {
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

func get679d482aA4c8474fB43e42c534237b76() string {
    dto := make(map[string]interface{})
    // 收款方附加数据
    // dto["addn_data"] = ""
    // 地区信息
    // dto["area_info"] = ""
    // 前台通知地址
    // dto["front_url"] = ""
    // 收款方附言
    // dto["payee_comments"] = ""
    // 收款方信息
    // dto["payee_info"] = get6b0066137be34e0d913fA123d87b6ecd()
    // 银联分配的服务商机构标识码
    // dto["pnr_ins_id_cd"] = ""
    // 请求方自定义域
    // dto["req_reserved"] = ""
    // 终端信息
    // dto["term_info"] = ""
    // 服务商信息
    // dto["pid_info"] = getD2810b7188724ceeAffdE9d1d6c9a3ff()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get7259d42290ee414d9275411e468ce9bc() string {
    dto := make(map[string]interface{})
    // 汇付机具号
    // dto["devs_id"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getFa26e3a2D80d410d80ad19246e8889a4() string {
    dto := make(map[string]interface{})
    // 付款方名称
    // dto["certificate_name"] = ""
    // 付款方银行卡号
    // dto["bank_card_no"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

