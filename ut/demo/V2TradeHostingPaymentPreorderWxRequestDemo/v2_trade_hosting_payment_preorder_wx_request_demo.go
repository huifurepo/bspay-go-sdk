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
        MiniappData:get3af138f408f64a088f39A98432ac9e49(),
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
    // 是否延迟交易
    extendInfoMap["delay_acct_flag"] = "Y"
    // 是否合单支付
    // extendInfoMap["merge_pay_flag"] = ""
    // 子单信息列表
    // extendInfoMap["sub_orders_request"] = get9143825f4e4c4d89A43a9b7761ced267()
    // 是否拆单支付
    // extendInfoMap["split_pay_flag"] = ""
    // 拆单支付参数集合
    // extendInfoMap["split_pay_data"] = get3aed07d204df4d00Bad4A1b45ed2ae5c()
    // 分账对象
    extendInfoMap["acct_split_bunch"] = get1a27caaf4d0f4d43Bcca1b32c877da44()
    // 统一收银台扩展参数集合
    // extendInfoMap["hosting_data"] = get072275545b724f46809255f85a93bc26()
    // 交易失效时间
    // extendInfoMap["time_expire"] = ""
    // 业务信息
    // extendInfoMap["biz_info"] = getB9a20bfd2a8c4dc3Af5278ecf486ea99()
    // 交易异步通知地址
    extendInfoMap["notify_url"] = "https://callback.service.com/xx"
    // 微信参数集合
    // extendInfoMap["wx_data"] = getD4b9debf592144be846eB7eb65b13804()
    // 设备信息
    // extendInfoMap["terminal_device_data"] = get69889c1cC4cf4c948329Ab0313a35373()
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

func get9143825f4e4c4d89A43a9b7761ced267() string {
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

func get3aed07d204df4d00Bad4A1b45ed2ae5c() string {
    dto := make(map[string]interface{})
    // 商户贴息标记
    // dto["fq_mer_discount_flag"] = ""
    // 商户业务信息
    // dto["ali_business_params"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get9b9a6c95A13f4c6b9fd69a2f41f9d2c3() interface{} {
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

func get1a27caaf4d0f4d43Bcca1b32c877da44() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = get9b9a6c95A13f4c6b9fd69a2f41f9d2c3()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get072275545b724f46809255f85a93bc26() string {
    dto := make(map[string]interface{})
    // 项目号
    // dto["project_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get3af138f408f64a088f39A98432ac9e49() string {
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

func get663201c25e1b4c678b0c241434e58ad0() interface{} {
    dto := make(map[string]interface{})
    // 指定支付者
    // dto["limit_payer"] = ""
    // 微信实名验证
    // dto["real_name_flag"] = ""

    return dto;
}

func get7b416ac5A2624f95Bb779a1ae47c3c22() interface{} {
    dto := make(map[string]interface{})
    // 姓名
    // dto["name"] = ""
    // 证件类型
    // dto["cert_type"] = ""
    // 证件号
    // dto["cert_no"] = ""

    return dto;
}

func getB9a20bfd2a8c4dc3Af5278ecf486ea99() string {
    dto := make(map[string]interface{})
    // 付款人验证（微信）
    // dto["payer_check_wx"] = get663201c25e1b4c678b0c241434e58ad0()
    // 个人付款人信息
    // dto["person_payer"] = get7b416ac5A2624f95Bb779a1ae47c3c22()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get38f2595f939e450dA5693c68155dc8a6() interface{} {
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

func getE5b7f3ebE834421fBa1a4347d5559700() interface{} {
    dto := make(map[string]interface{})
    // 单品列表
    // dto["goods_detail"] = get38f2595f939e450dA5693c68155dc8a6()
    // 订单原价(元)
    // dto["cost_price"] = ""
    // 商品小票ID
    // dto["receipt_id"] = ""

    return dto;
}

func getD16c2f5aB29142e3B8a04ead1c6e5af8() interface{} {
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

func getC04e503eE2fe4166Aa9b0a5dd323b83f() interface{} {
    dto := make(map[string]interface{})
    // 门店信息
    // dto["store_info"] = getD16c2f5aB29142e3B8a04ead1c6e5af8()

    return dto;
}

func getD4b9debf592144be846eB7eb65b13804() string {
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
    // dto["detail"] = getE5b7f3ebE834421fBa1a4347d5559700()
    // 设备号
    // dto["device_info"] = ""
    // 订单优惠标记
    // dto["goods_tag"] = ""
    // 实名支付
    // dto["identity"] = ""
    // 开发票入口开放标识
    // dto["receipt"] = ""
    // 场景信息
    // dto["scene_info"] = getC04e503eE2fe4166Aa9b0a5dd323b83f()
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

func get69889c1cC4cf4c948329Ab0313a35373() string {
    dto := make(map[string]interface{})
    // 汇付机具号
    // dto["devs_id"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

