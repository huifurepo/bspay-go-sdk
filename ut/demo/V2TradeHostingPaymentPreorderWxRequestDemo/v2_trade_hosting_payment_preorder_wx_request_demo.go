/**
 * 微信小程序预下单接口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeHostingPaymentPreorderWxRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeHostingPaymentPreorderWxRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeHostingPaymentPreorderWxRequest{
        // 预下单类型
        PreOrderType:"3",
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000109133323",
        // 交易金额
        TransAmt:"0.13",
        // 商品描述
        GoodsDesc:"app跳微信消费",
        // 微信小程序扩展参数集合
        MiniappData:getAbbf77a95ef247ddB2716b934b7b2823(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeHostingPaymentPreorderWxRequest(dgReq)
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
    extendInfoMap["delay_acct_flag"] = "Y"
    // 是否拆单支付
    // extendInfoMap["split_pay_flag"] = ""
    // 拆单支付参数集合
    // extendInfoMap["split_pay_data"] = get4c305391Aa1346e5A8a48882209e79ba()
    // 分账对象
    extendInfoMap["acct_split_bunch"] = get8604df30Ffb6463bB0186219b4af33ac()
    // 交易失效时间
    // extendInfoMap["time_expire"] = ""
    // 业务信息
    // extendInfoMap["biz_info"] = getC912bca127b84615A483D8a5dbea2527()
    // 交易异步通知地址
    extendInfoMap["notify_url"] = "https://callback.service.com/xx"
    // 回调地址
    // extendInfoMap["callback_url"] = ""
    // 微信参数集合
    // extendInfoMap["wx_data"] = getEa9815a59b9b4d76Ae6eE3fba17fcbad()
    // 设备信息
    // extendInfoMap["terminal_device_data"] = getF2095d8eD2184bbe85305961af571c39()
    // 手续费场景标识
    // extendInfoMap["fee_sign"] = ""
    return extendInfoMap
}

func get4c305391Aa1346e5A8a48882209e79ba() string {
    dto := make(map[string]interface{})
    // 商户贴息标记
    // dto["fq_mer_discount_flag"] = ""
    // 商户业务信息
    // dto["ali_business_params"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get57e00fb9F2d443e6Bb83Ddd2bc3af3ab() interface{} {
    dto := make(map[string]interface{})
    // 分账金额
    dto["div_amt"] = "0.01"
    // 分账接收方ID
    dto["huifu_id"] = "6666000109133323"
    // 收款汇付账户号
    // dto["acct_id"] = ""
    // 分账百分比%
    // dto["percentage_div"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get8604df30Ffb6463bB0186219b4af33ac() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = get57e00fb9F2d443e6Bb83Ddd2bc3af3ab()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getAbbf77a95ef247ddB2716b934b7b2823() string {
    dto := make(map[string]interface{})
    // 是否生成scheme_code
    dto["need_scheme"] = "Y"
    // 应用ID
    dto["seq_id"] = "APP_2022100912694428"
    // 私有信息
    dto["private_info"] = "oppsHosting://"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getFb96f5b0306f4eec832718831cca3255() interface{} {
    dto := make(map[string]interface{})
    // 指定支付者
    // dto["limit_payer"] = ""
    // 微信实名验证
    // dto["real_name_flag"] = ""

    return dto;
}

func get035d6c1761b04f31B451Bd98650869c7() interface{} {
    dto := make(map[string]interface{})
    // 姓名
    // dto["name"] = ""
    // 证件类型
    // dto["cert_type"] = ""
    // 证件号
    // dto["cert_no"] = ""

    return dto;
}

func getC912bca127b84615A483D8a5dbea2527() string {
    dto := make(map[string]interface{})
    // 付款人验证（微信）
    // dto["payer_check_wx"] = getFb96f5b0306f4eec832718831cca3255()
    // 个人付款人信息
    // dto["person_payer"] = get035d6c1761b04f31B451Bd98650869c7()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get5e329288D3e748caBf19A18d37ce5f74() interface{} {
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

func getE12c2d6877154b20B97846493d31285a() interface{} {
    dto := make(map[string]interface{})
    // 单品列表
    // dto["goods_detail"] = get5e329288D3e748caBf19A18d37ce5f74()
    // 订单原价(元)
    // dto["cost_price"] = ""
    // 商品小票ID
    // dto["receipt_id"] = ""

    return dto;
}

func get38166442Ea5340bcB8442e1a6f47b865() interface{} {
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

func get183762aeA75d4f91A0bfC24e963166dd() interface{} {
    dto := make(map[string]interface{})
    // 门店信息
    // dto["store_info"] = get38166442Ea5340bcB8442e1a6f47b865()

    return dto;
}

func getEa9815a59b9b4d76Ae6eE3fba17fcbad() string {
    dto := make(map[string]interface{})
    // 子商户应用ID
    // dto["sub_appid"] = ""
    // 子商户用户标识
    // dto["sub_openid"] = ""
    // 附加数据
    // dto["attach"] = ""
    // 商品描述
    // dto["body"] = ""
    // 商品详情
    // dto["detail"] = getE12c2d6877154b20B97846493d31285a()
    // 设备号
    // dto["device_info"] = ""
    // 订单优惠标记
    // dto["goods_tag"] = ""
    // 实名支付
    // dto["identity"] = ""
    // 开发票入口开放标识
    // dto["receipt"] = ""
    // 场景信息
    // dto["scene_info"] = get183762aeA75d4f91A0bfC24e963166dd()
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

func getF2095d8eD2184bbe85305961af571c39() string {
    dto := make(map[string]interface{})
    // 汇付机具号
    // dto["devs_id"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

