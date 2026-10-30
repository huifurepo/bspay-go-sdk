/**
 * 支付宝小程序预下单接口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeHostingPaymentPreorderAliRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeHostingPaymentPreorderAliRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeHostingPaymentPreorderAliRequest{
        // 商户号
        HuifuId:"6666000109133323",
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 预下单类型
        PreOrderType:"2",
        // 交易金额
        TransAmt:"0.10",
        // 商品描述
        GoodsDesc:"app跳支付宝消费",
        // app扩展参数集合
        AppData:get071c7f8fCe5341708f20F1cf750b3308(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeHostingPaymentPreorderAliRequest(dgReq)
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
    // 是否合单支付
    // extendInfoMap["merge_pay_flag"] = ""
    // 子单信息列表
    // extendInfoMap["sub_orders_request"] = getE35bb59f055046438dfa5e2f822c68c1()
    // 分账对象
    extendInfoMap["acct_split_bunch"] = getB030a02bA5c8472c9cff00ed4cacb3a7()
    // 统一收银台扩展参数集合
    // extendInfoMap["hosting_data"] = get7e0c32a9Cea440a1B8839321e39ed0fc()
    // 交易失效时间
    // extendInfoMap["time_expire"] = ""
    // 业务信息
    // extendInfoMap["biz_info"] = get7398d9ecC6114d4eA0977142e0c11757()
    // 异步通知地址
    extendInfoMap["notify_url"] = "https://callback.service.com/xx"
    // 支付宝参数集合
    // extendInfoMap["alipay_data"] = get948f4498D5bb46b7979cA19d0f6c16df()
    // 设备信息
    // extendInfoMap["terminal_device_data"] = getBf7b33fb979847f9890eCe3044270146()
    // 手续费场景标识
    // extendInfoMap["fee_sign"] = ""
    // 是否交易手续费分摊
    // extendInfoMap["fee_split_flag"] = ""
    // 手续费扣款标志
    // extendInfoMap["fee_flag"] = ""
    // 渠道号
    // extendInfoMap["channel_no"] = ""
    // 场景类型
    // extendInfoMap["pay_scene"] = ""
    return extendInfoMap
}

func getE35bb59f055046438dfa5e2f822c68c1() string {
    dto := make(map[string]interface{})
    // 子单请求流水号
    // dto["sub_req_seq_id"] = "test"
    // 子单商户号
    // dto["sub_huifu_id"] = "test"
    // 子单交易金额
    // dto["sub_trans_amt"] = "test"
    // 子单账户号
    // dto["sub_acct_id"] = ""
    // 是否延迟交易
    // dto["delay_acct_flag"] = ""
    // 子单备注
    // dto["sub_remark"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get26f8c4beCab64ba1834fDd7fa82cf1a7() interface{} {
    dto := make(map[string]interface{})
    // 分账金额
    dto["div_amt"] = "0.08"
    // 分账接收方ID
    dto["huifu_id"] = "6666000109133323"
    // 收款汇付账户号
    // dto["acct_id"] = ""
    // 分账百分比%
    // dto["percentage_div"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getB030a02bA5c8472c9cff00ed4cacb3a7() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = get26f8c4beCab64ba1834fDd7fa82cf1a7()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get7e0c32a9Cea440a1B8839321e39ed0fc() string {
    dto := make(map[string]interface{})
    // 项目号
    // dto["project_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get071c7f8fCe5341708f20F1cf750b3308() string {
    dto := make(map[string]interface{})
    // 小程序返回码
    dto["app_schema"] = "app跳转链接"
    // 支付宝小程序ID
    // dto["appid"] = ""
    // 私有信息
    // dto["private_info"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get91f6489b0f5446d19da30042b5e26453() interface{} {
    dto := make(map[string]interface{})
    // 是否提供校验身份信息
    // dto["need_check_info"] = ""
    // 允许的最小买家年龄
    // dto["min_age"] = ""
    // 是否强制校验付款人身份信息
    // dto["fix_buyer"] = ""

    return dto;
}

func get2effa40bA18e472dB4cc13ac8f232ff4() interface{} {
    dto := make(map[string]interface{})
    // 姓名
    // dto["name"] = ""
    // 证件类型
    // dto["cert_type"] = ""
    // 证件号
    // dto["cert_no"] = ""
    // 手机号
    // dto["mobile"] = ""

    return dto;
}

func get7398d9ecC6114d4eA0977142e0c11757() string {
    dto := make(map[string]interface{})
    // 付款人验证（支付宝）
    // dto["payer_check_ali"] = get91f6489b0f5446d19da30042b5e26453()
    // 个人付款人信息
    // dto["person_payer"] = get2effa40bA18e472dB4cc13ac8f232ff4()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get10722936D7694c069cef864fc36d8120() interface{} {
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

func get3281c6153ee141bf92265e1128648634() interface{} {
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

func get948f4498D5bb46b7979cA19d0f6c16df() string {
    dto := make(map[string]interface{})
    // 支付宝的店铺编号
    // dto["alipay_store_id"] = ""
    // 业务扩展参数
    // dto["extend_params"] = get10722936D7694c069cef864fc36d8120()
    // 订单包含的商品列表信息
    // dto["goods_detail"] = get3281c6153ee141bf92265e1128648634()
    // 商户原始订单号
    // dto["merchant_order_no"] = ""
    // 商户操作员编号
    // dto["operator_id"] = ""
    // 产品码
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
    // 商户贴息标记
    // dto["fq_mer_discount_flag"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getBf7b33fb979847f9890eCe3044270146() string {
    dto := make(map[string]interface{})
    // 汇付机具号
    // dto["devs_id"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

