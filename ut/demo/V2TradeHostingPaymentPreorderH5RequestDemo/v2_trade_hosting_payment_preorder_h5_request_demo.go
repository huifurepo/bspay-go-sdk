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
        HostingData:get3cf8e10b6c8149c992080777dace9fd5(),
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
    extendInfoMap["acct_split_bunch"] = getEacc2d5f2b734de68812B1f0df3f9112()
    // 交易失效时间
    // extendInfoMap["time_expire"] = ""
    // 业务信息
    extendInfoMap["biz_info"] = get670e8f4439404e90B3164f64b638ff68()
    // 交易异步通知地址
    extendInfoMap["notify_url"] = "https://callback.service.com/xx"
    // 使用类型
    // extendInfoMap["usage_type"] = ""
    // 交易类型
    // extendInfoMap["trans_type"] = ""
    // 微信参数集合
    // extendInfoMap["wx_data"] = get29169b0b0e654ba286c2F801a3d2e4da()
    // 支付宝参数集合
    // extendInfoMap["alipay_data"] = get09d7a941E5fa430483eeDbe63b5ac9ff()
    // 抖音参数集合
    // extendInfoMap["dy_data"] = getFa54ed35C1874c9f9a27A9128ca0dddd()
    // 银联参数集合
    // extendInfoMap["unionpay_data"] = get7a67fe4b66e1493cA3bf07dd83f565ed()
    // 设备信息
    // extendInfoMap["terminal_device_data"] = get2844b911B316400b991138262e5350a0()
    // 大额支付参数集合
    // extendInfoMap["largeamt_data"] = get288703c8E841479c812901a133bb2220()
    // 手续费场景标识
    // extendInfoMap["fee_sign"] = ""
    // 是否交易手续费分摊
    // extendInfoMap["fee_split_flag"] = ""
    return extendInfoMap
}

func get6d0e47e51d8e422dA89fA4e48ea89012() interface{} {
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

func getEacc2d5f2b734de68812B1f0df3f9112() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = get6d0e47e51d8e422dA89fA4e48ea89012()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get3cf8e10b6c8149c992080777dace9fd5() string {
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

func getAb8059e294bf4911A1855df75741104e() interface{} {
    dto := make(map[string]interface{})
    // 是否提供校验身份信息
    dto["need_check_info"] = "T"
    // 允许的最小买家年龄
    dto["min_age"] = "12"
    // 是否强制校验付款人身份信息
    dto["fix_buyer"] = "F"

    return dto;
}

func getD6f5ffa76e2549e2Af5aB33ef837631f() interface{} {
    dto := make(map[string]interface{})
    // 指定支付者
    dto["limit_payer"] = "ADULT"
    // 微信实名验证
    dto["real_name_flag"] = "Y"

    return dto;
}

func get7cca7e8d1743479997d44b2a2cc8fb38() interface{} {
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

func get670e8f4439404e90B3164f64b638ff68() string {
    dto := make(map[string]interface{})
    // 付款人验证（支付宝）
    dto["payer_check_ali"] = getAb8059e294bf4911A1855df75741104e()
    // 付款人验证（微信）
    dto["payer_check_wx"] = getD6f5ffa76e2549e2Af5aB33ef837631f()
    // 个人付款人信息
    dto["person_payer"] = get7cca7e8d1743479997d44b2a2cc8fb38()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getDb155469Ca054803Bd74Ac64b3c86b93() interface{} {
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

func get8bd3d39e53ed49d59c325c85b6f3ce55() interface{} {
    dto := make(map[string]interface{})
    // 单品列表
    // dto["goods_detail"] = getDb155469Ca054803Bd74Ac64b3c86b93()
    // 订单原价(元)
    // dto["cost_price"] = ""
    // 商品小票ID
    // dto["receipt_id"] = ""

    return dto;
}

func get108be90e4cf84e3eA4e5Cc292aea34c9() interface{} {
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

func getAeaf534d25e4439c82cb3f4492b3d395() interface{} {
    dto := make(map[string]interface{})
    // 门店信息
    // dto["store_info"] = get108be90e4cf84e3eA4e5Cc292aea34c9()

    return dto;
}

func get29169b0b0e654ba286c2F801a3d2e4da() string {
    dto := make(map[string]interface{})
    // 附加数据
    // dto["attach"] = ""
    // 商品详情
    // dto["detail"] = get8bd3d39e53ed49d59c325c85b6f3ce55()
    // 订单优惠标记
    // dto["goods_tag"] = ""
    // 开发票入口开放标识
    // dto["receipt"] = ""
    // 场景信息
    // dto["scene_info"] = getAeaf534d25e4439c82cb3f4492b3d395()
    // 单品优惠标识
    // dto["promotion_flag"] = ""
    // 新增商品ID
    // dto["product_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getF4a7920453104ce7B594Fa86cdefb284() interface{} {
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

func get2de9e2673d1a473082847e2261690167() interface{} {
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

func get09d7a941E5fa430483eeDbe63b5ac9ff() string {
    dto := make(map[string]interface{})
    // 支付宝的店铺编号
    // dto["alipay_store_id"] = ""
    // 业务扩展参数
    // dto["extend_params"] = getF4a7920453104ce7B594Fa86cdefb284()
    // 订单包含的商品列表信息
    // dto["goods_detail"] = get2de9e2673d1a473082847e2261690167()
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

func get101501b539504c169701B3d726a84cbb() interface{} {
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

func get8316421d73d34c8cA7a6F2c4c538897c() interface{} {
    dto := make(map[string]interface{})
    // 用户终端IP
    // dto["payer_client_ip"] = "test"

    return dto;
}

func getFa54ed35C1874c9f9a27A9128ca0dddd() string {
    dto := make(map[string]interface{})
    // 子商户应用ID
    // dto["sub_appid"] = "test"
    // H5场景信息
    // dto["h5_info"] = get101501b539504c169701B3d726a84cbb()
    // 场景信息
    // dto["scene_info"] = get8316421d73d34c8cA7a6F2c4c538897c()
    // 优惠标记
    // dto["coupon_info"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getA55bee88Bcf644bdAc4fE766ee9ee942() interface{} {
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

func get6495ba8d8d7c468e840517a79b59e2a3() string {
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

func get7a67fe4b66e1493cA3bf07dd83f565ed() string {
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
    // dto["payee_info"] = getA55bee88Bcf644bdAc4fE766ee9ee942()
    // 银联分配的服务商机构标识码
    // dto["pnr_ins_id_cd"] = ""
    // 请求方自定义域
    // dto["req_reserved"] = ""
    // 终端信息
    // dto["term_info"] = ""
    // 服务商信息
    // dto["pid_info"] = get6495ba8d8d7c468e840517a79b59e2a3()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get2844b911B316400b991138262e5350a0() string {
    dto := make(map[string]interface{})
    // 汇付机具号
    // dto["devs_id"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get288703c8E841479c812901a133bb2220() string {
    dto := make(map[string]interface{})
    // 付款方名称
    // dto["certificate_name"] = ""
    // 付款方银行卡号
    // dto["bank_card_no"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

