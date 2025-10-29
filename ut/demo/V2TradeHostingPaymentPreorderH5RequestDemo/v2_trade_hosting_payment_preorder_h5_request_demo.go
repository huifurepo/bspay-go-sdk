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
        HostingData:get1e8acadc6fdc437d924301f529cd63d1(),
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
    // 收银台样式
    // extendInfoMap["style_id"] = ""
    // 是否延迟交易
    extendInfoMap["delay_acct_flag"] = "N"
    // 分账对象
    extendInfoMap["acct_split_bunch"] = get6aaa622dC05741489a24535e5b82b99e()
    // 交易失效时间
    // extendInfoMap["time_expire"] = ""
    // 业务信息
    extendInfoMap["biz_info"] = get37e20504E1214db9Ac571bcde55dda3b()
    // 交易异步通知地址
    extendInfoMap["notify_url"] = "https://callback.service.com/xx"
    // 使用类型
    // extendInfoMap["usage_type"] = ""
    // 交易类型
    // extendInfoMap["trans_type"] = ""
    // 微信参数集合
    // extendInfoMap["wx_data"] = getD8a35cd51dbb4b1799a5078dd16d794b()
    // 支付宝参数集合
    // extendInfoMap["alipay_data"] = get7462d7ef7358431b9dd40716cba62e82()
    // 银联参数集合
    // extendInfoMap["unionpay_data"] = get64e25a801c014bcfB722A3161947b191()
    // 设备信息
    // extendInfoMap["terminal_device_data"] = getD177b96cB0f94a3e8595F04fa644ac8b()
    // 大额支付参数集合
    // extendInfoMap["largeamt_data"] = getF76d1606Cdbc4b0c800c20a5de0f311c()
    return extendInfoMap
}

func get71ab5f815f144fe9B7e388460db1238a() interface{} {
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

func get6aaa622dC05741489a24535e5b82b99e() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = get71ab5f815f144fe9B7e388460db1238a()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get1e8acadc6fdc437d924301f529cd63d1() string {
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

func getD403b82e5aa44b21B39974b6192422a3() interface{} {
    dto := make(map[string]interface{})
    // 是否提供校验身份信息
    dto["need_check_info"] = "T"
    // 允许的最小买家年龄
    dto["min_age"] = "12"
    // 是否强制校验付款人身份信息
    dto["fix_buyer"] = "F"

    return dto;
}

func get2c5f3e3a08b74869869e4eec2287cb45() interface{} {
    dto := make(map[string]interface{})
    // 指定支付者
    dto["limit_payer"] = "ADULT"
    // 微信实名验证
    dto["real_name_flag"] = "Y"

    return dto;
}

func get28437a3d21f54b2a899dBc310e632e50() interface{} {
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

func get37e20504E1214db9Ac571bcde55dda3b() string {
    dto := make(map[string]interface{})
    // 付款人验证（支付宝）
    dto["payer_check_ali"] = getD403b82e5aa44b21B39974b6192422a3()
    // 付款人验证（微信）
    dto["payer_check_wx"] = get2c5f3e3a08b74869869e4eec2287cb45()
    // 个人付款人信息
    dto["person_payer"] = get28437a3d21f54b2a899dBc310e632e50()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getDf3ec4aaC53e4cb3B569Cd8e44e16021() interface{} {
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

func getD432172e2d674608A5f5Edcb2d60860b() interface{} {
    dto := make(map[string]interface{})
    // 单品列表
    // dto["goods_detail"] = getDf3ec4aaC53e4cb3B569Cd8e44e16021()
    // 订单原价(元)
    // dto["cost_price"] = ""
    // 商品小票ID
    // dto["receipt_id"] = ""

    return dto;
}

func getEdfbb33e522d41a38ea0376b4257003b() interface{} {
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

func getD36cb51963164276Bb1fE85625c1d8fd() interface{} {
    dto := make(map[string]interface{})
    // 门店信息
    // dto["store_info"] = getEdfbb33e522d41a38ea0376b4257003b()

    return dto;
}

func getD8a35cd51dbb4b1799a5078dd16d794b() string {
    dto := make(map[string]interface{})
    // 附加数据
    // dto["attach"] = ""
    // 商品详情
    // dto["detail"] = getD432172e2d674608A5f5Edcb2d60860b()
    // 订单优惠标记
    // dto["goods_tag"] = ""
    // 开发票入口开放标识
    // dto["receipt"] = ""
    // 场景信息
    // dto["scene_info"] = getD36cb51963164276Bb1fE85625c1d8fd()
    // 单品优惠标识
    // dto["promotion_flag"] = ""
    // 新增商品ID
    // dto["product_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getA67aa70243cd42e898ce68f05c9b7f49() interface{} {
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

func getD2e6935496b042919191Fc96db8735f1() interface{} {
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

func get7462d7ef7358431b9dd40716cba62e82() string {
    dto := make(map[string]interface{})
    // 支付宝的店铺编号
    // dto["alipay_store_id"] = ""
    // 业务扩展参数
    // dto["extend_params"] = getA67aa70243cd42e898ce68f05c9b7f49()
    // 订单包含的商品列表信息
    // dto["goods_detail"] = getD2e6935496b042919191Fc96db8735f1()
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

func getE077e443C0d2414bA82337c4b549fc3d() interface{} {
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

func get64e25a801c014bcfB722A3161947b191() string {
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
    // dto["payee_info"] = getE077e443C0d2414bA82337c4b549fc3d()
    // 银联分配的服务商机构标识码
    // dto["pnr_ins_id_cd"] = ""
    // 请求方自定义域
    // dto["req_reserved"] = ""
    // 终端信息
    // dto["term_info"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getD177b96cB0f94a3e8595F04fa644ac8b() string {
    dto := make(map[string]interface{})
    // 汇付机具号
    // dto["devs_id"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getF76d1606Cdbc4b0c800c20a5de0f311c() string {
    dto := make(map[string]interface{})
    // 付款方名称
    // dto["certificate_name"] = ""
    // 付款方银行卡号
    // dto["bank_card_no"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

