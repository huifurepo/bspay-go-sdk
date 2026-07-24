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
        MiniappData:getD9703d52B92d48238f2f8efaffc683a8(),
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
    // 是否拆单支付
    // extendInfoMap["split_pay_flag"] = ""
    // 拆单支付参数集合
    // extendInfoMap["split_pay_data"] = getA9b4116c348246a09256Cb92ff1468a1()
    // 分账对象
    extendInfoMap["acct_split_bunch"] = getB56e94c601784737Bd351f6915482693()
    // 统一收银台扩展参数集合
    // extendInfoMap["hosting_data"] = get1e3a25fa77f145358744D9607c3d5c6d()
    // 交易失效时间
    // extendInfoMap["time_expire"] = ""
    // 业务信息
    // extendInfoMap["biz_info"] = get2b8f9b9dDd5e4bbd8ebb3b814144f1c8()
    // 交易异步通知地址
    extendInfoMap["notify_url"] = "https://callback.service.com/xx"
    // 微信参数集合
    // extendInfoMap["wx_data"] = getA5c774deD00546ab866d5fe1d69040e5()
    // 设备信息
    // extendInfoMap["terminal_device_data"] = getA090526f7e75421a9c50E1cf7710783d()
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

func getA9b4116c348246a09256Cb92ff1468a1() string {
    dto := make(map[string]interface{})
    // 商户贴息标记
    // dto["fq_mer_discount_flag"] = ""
    // 商户业务信息
    // dto["ali_business_params"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get148c81120b084d9fBd7f36f2c211ff9b() interface{} {
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

func getB56e94c601784737Bd351f6915482693() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = get148c81120b084d9fBd7f36f2c211ff9b()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get1e3a25fa77f145358744D9607c3d5c6d() string {
    dto := make(map[string]interface{})
    // 项目号
    // dto["project_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getD9703d52B92d48238f2f8efaffc683a8() string {
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

func get1c48675d4a224883A357C3cbbfa7b30f() interface{} {
    dto := make(map[string]interface{})
    // 指定支付者
    // dto["limit_payer"] = ""
    // 微信实名验证
    // dto["real_name_flag"] = ""

    return dto;
}

func getFcc910e32d31434eBa686cdfe1c93169() interface{} {
    dto := make(map[string]interface{})
    // 姓名
    // dto["name"] = ""
    // 证件类型
    // dto["cert_type"] = ""
    // 证件号
    // dto["cert_no"] = ""

    return dto;
}

func get2b8f9b9dDd5e4bbd8ebb3b814144f1c8() string {
    dto := make(map[string]interface{})
    // 付款人验证（微信）
    // dto["payer_check_wx"] = get1c48675d4a224883A357C3cbbfa7b30f()
    // 个人付款人信息
    // dto["person_payer"] = getFcc910e32d31434eBa686cdfe1c93169()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getFd43fd8f56b04805994c660c8ab819f6() interface{} {
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

func get008abb5d9f31431aA8531bc2d077d9cf() interface{} {
    dto := make(map[string]interface{})
    // 单品列表
    // dto["goods_detail"] = getFd43fd8f56b04805994c660c8ab819f6()
    // 订单原价(元)
    // dto["cost_price"] = ""
    // 商品小票ID
    // dto["receipt_id"] = ""

    return dto;
}

func get984fc4b21ed145b396b00461b6b1f98d() interface{} {
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

func get749cbf43F8bd452688d19b44ecf68848() interface{} {
    dto := make(map[string]interface{})
    // 门店信息
    // dto["store_info"] = get984fc4b21ed145b396b00461b6b1f98d()

    return dto;
}

func getA5c774deD00546ab866d5fe1d69040e5() string {
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
    // dto["detail"] = get008abb5d9f31431aA8531bc2d077d9cf()
    // 设备号
    // dto["device_info"] = ""
    // 订单优惠标记
    // dto["goods_tag"] = ""
    // 实名支付
    // dto["identity"] = ""
    // 开发票入口开放标识
    // dto["receipt"] = ""
    // 场景信息
    // dto["scene_info"] = get749cbf43F8bd452688d19b44ecf68848()
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

func getA090526f7e75421a9c50E1cf7710783d() string {
    dto := make(map[string]interface{})
    // 汇付机具号
    // dto["devs_id"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

