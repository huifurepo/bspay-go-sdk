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
        GoodsDesc:"支付托管消费",
        // 预下单类型
        PreOrderType:"1",
        // 半支付托管扩展参数集合
        HostingData:get97399eabA25d41d38d7135da80c50ab2(),
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
    extendInfoMap["acct_split_bunch"] = get2e15193dC1e74cd19ba2A0db5a62ea59()
    // 交易失效时间
    // extendInfoMap["time_expire"] = ""
    // 业务信息
    extendInfoMap["biz_info"] = getD09496cb54354998Bd3c490f29795d25()
    // 交易异步通知地址
    extendInfoMap["notify_url"] = "https://callback.service.com/xx"
    // 使用类型
    // extendInfoMap["usage_type"] = ""
    // 交易类型
    // extendInfoMap["trans_type"] = ""
    // 微信参数集合
    // extendInfoMap["wx_data"] = getF650ecb742344e21Ab0bD6ab713d2a65()
    // 支付宝参数集合
    // extendInfoMap["alipay_data"] = get59bd3ff360df40989b888707d8bc0733()
    // 抖音参数集合
    // extendInfoMap["dy_data"] = getEa804efc10e44fe28f5d4d7029bad882()
    // 银联参数集合
    // extendInfoMap["unionpay_data"] = getE8d452bf85544f5e890c34e25686a9c4()
    // 设备信息
    // extendInfoMap["terminal_device_data"] = get2bf02cdfFb804336937aA9bf0b2881d0()
    // 大额支付参数集合
    // extendInfoMap["largeamt_data"] = get8917c8c32c6d42859eb752be632db2a8()
    // 手续费场景标识
    // extendInfoMap["fee_sign"] = ""
    return extendInfoMap
}

func get6a2b7220Fb2345e7Bd1eAefffe3510fc() interface{} {
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

func get2e15193dC1e74cd19ba2A0db5a62ea59() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = get6a2b7220Fb2345e7Bd1eAefffe3510fc()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get97399eabA25d41d38d7135da80c50ab2() string {
    dto := make(map[string]interface{})
    // 项目标题
    dto["project_title"] = "收银台标题"
    // 半支付托管项目号
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

func get2767ae1870aa4e34Ac499f850e0460a7() interface{} {
    dto := make(map[string]interface{})
    // 是否提供校验身份信息
    dto["need_check_info"] = "T"
    // 允许的最小买家年龄
    dto["min_age"] = "12"
    // 是否强制校验付款人身份信息
    dto["fix_buyer"] = "F"

    return dto;
}

func get8fe397b23f0d4cadBc91Bb2d9c619b7d() interface{} {
    dto := make(map[string]interface{})
    // 指定支付者
    dto["limit_payer"] = "ADULT"
    // 微信实名验证
    dto["real_name_flag"] = "Y"

    return dto;
}

func get8ee44447B4464f798f5a6ff92175933a() interface{} {
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

func getD09496cb54354998Bd3c490f29795d25() string {
    dto := make(map[string]interface{})
    // 付款人验证（支付宝）
    dto["payer_check_ali"] = get2767ae1870aa4e34Ac499f850e0460a7()
    // 付款人验证（微信）
    dto["payer_check_wx"] = get8fe397b23f0d4cadBc91Bb2d9c619b7d()
    // 个人付款人信息
    dto["person_payer"] = get8ee44447B4464f798f5a6ff92175933a()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getDa98ad1e3e1346d6Ad8e42f48dd578c8() interface{} {
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

func get2c5f1ef951b64c3aB9103133bb0a7ba6() interface{} {
    dto := make(map[string]interface{})
    // 单品列表
    // dto["goods_detail"] = getDa98ad1e3e1346d6Ad8e42f48dd578c8()
    // 订单原价(元)
    // dto["cost_price"] = ""
    // 商品小票ID
    // dto["receipt_id"] = ""

    return dto;
}

func get79d97887Da0d445dBfffC4d306e879c2() interface{} {
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

func get3d1e0df6561a4a5e92cbC06fb1d462cd() interface{} {
    dto := make(map[string]interface{})
    // 门店信息
    // dto["store_info"] = get79d97887Da0d445dBfffC4d306e879c2()

    return dto;
}

func getF650ecb742344e21Ab0bD6ab713d2a65() string {
    dto := make(map[string]interface{})
    // 附加数据
    // dto["attach"] = ""
    // 商品详情
    // dto["detail"] = get2c5f1ef951b64c3aB9103133bb0a7ba6()
    // 订单优惠标记
    // dto["goods_tag"] = ""
    // 开发票入口开放标识
    // dto["receipt"] = ""
    // 场景信息
    // dto["scene_info"] = get3d1e0df6561a4a5e92cbC06fb1d462cd()
    // 单品优惠标识
    // dto["promotion_flag"] = ""
    // 新增商品ID
    // dto["product_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get88742fe10e2844eb8adcD452c0513890() interface{} {
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

func get1cf9c5be08fe406bB7e91c7f01c24e8c() interface{} {
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

func get59bd3ff360df40989b888707d8bc0733() string {
    dto := make(map[string]interface{})
    // 支付宝的店铺编号
    // dto["alipay_store_id"] = ""
    // 业务扩展参数
    // dto["extend_params"] = get88742fe10e2844eb8adcD452c0513890()
    // 订单包含的商品列表信息
    // dto["goods_detail"] = get1cf9c5be08fe406bB7e91c7f01c24e8c()
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

func get36ee098d5f12479595845fdd522d44c5() interface{} {
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

func get3796381aE64d403dA34411b0bb6d0a18() interface{} {
    dto := make(map[string]interface{})
    // 用户终端IP
    // dto["payer_client_ip"] = "test"

    return dto;
}

func getEa804efc10e44fe28f5d4d7029bad882() string {
    dto := make(map[string]interface{})
    // 子商户应用ID
    // dto["sub_appid"] = "test"
    // H5场景信息
    // dto["h5_info"] = get36ee098d5f12479595845fdd522d44c5()
    // 场景信息
    // dto["scene_info"] = get3796381aE64d403dA34411b0bb6d0a18()
    // 优惠标记
    // dto["coupon_info"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get292ec8fe80c9455a962eD6120493a567() interface{} {
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

func getE8d452bf85544f5e890c34e25686a9c4() string {
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
    // dto["payee_info"] = get292ec8fe80c9455a962eD6120493a567()
    // 银联分配的服务商机构标识码
    // dto["pnr_ins_id_cd"] = ""
    // 请求方自定义域
    // dto["req_reserved"] = ""
    // 终端信息
    // dto["term_info"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get2bf02cdfFb804336937aA9bf0b2881d0() string {
    dto := make(map[string]interface{})
    // 汇付机具号
    // dto["devs_id"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get8917c8c32c6d42859eb752be632db2a8() string {
    dto := make(map[string]interface{})
    // 付款方名称
    // dto["certificate_name"] = ""
    // 付款方银行卡号
    // dto["bank_card_no"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

