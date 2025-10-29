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
        MiniappData:get89e1260aC2a2430286f681c1d345d987(),
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
    // 分账对象
    extendInfoMap["acct_split_bunch"] = getAbb36b9851ec4c89B724991f36bbb537()
    // 交易失效时间
    // extendInfoMap["time_expire"] = ""
    // 业务信息
    // extendInfoMap["biz_info"] = get1e91233c3b514d21901161ff2c17461e()
    // 交易异步通知地址
    extendInfoMap["notify_url"] = "https://callback.service.com/xx"
    // 回调地址
    // extendInfoMap["callback_url"] = ""
    // 微信参数集合
    // extendInfoMap["wx_data"] = get0cccdfed4c4840adAe13C1b4160ea624()
    // 设备信息
    // extendInfoMap["terminal_device_data"] = getE8d01f1d074c4d33Ab52E96c0ae4aead()
    return extendInfoMap
}

func getFc087f7305f4441d8d44B944f81342e0() interface{} {
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

func getAbb36b9851ec4c89B724991f36bbb537() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = getFc087f7305f4441d8d44B944f81342e0()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get89e1260aC2a2430286f681c1d345d987() string {
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

func getF9c587ff49624ecdA2d766be46176bbb() interface{} {
    dto := make(map[string]interface{})
    // 指定支付者
    // dto["limit_payer"] = ""
    // 微信实名验证
    // dto["real_name_flag"] = ""

    return dto;
}

func get240df619021e46c982d4C5e7eb417f52() interface{} {
    dto := make(map[string]interface{})
    // 姓名
    // dto["name"] = ""
    // 证件类型
    // dto["cert_type"] = ""
    // 证件号
    // dto["cert_no"] = ""

    return dto;
}

func get1e91233c3b514d21901161ff2c17461e() string {
    dto := make(map[string]interface{})
    // 付款人验证（微信）
    // dto["payer_check_wx"] = getF9c587ff49624ecdA2d766be46176bbb()
    // 个人付款人信息
    // dto["person_payer"] = get240df619021e46c982d4C5e7eb417f52()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get53ae6690B1634528B5ec19860aff0a5e() interface{} {
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

func get5363f4df67b44a51B0d30b90fb0242dc() interface{} {
    dto := make(map[string]interface{})
    // 单品列表
    // dto["goods_detail"] = get53ae6690B1634528B5ec19860aff0a5e()
    // 订单原价(元)
    // dto["cost_price"] = ""
    // 商品小票ID
    // dto["receipt_id"] = ""

    return dto;
}

func get5dfa6188Ce324403A8ef7d1054e8d2b9() interface{} {
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

func get8e7acadb53e34c8c849561ed6576e6b1() interface{} {
    dto := make(map[string]interface{})
    // 门店信息
    // dto["store_info"] = get5dfa6188Ce324403A8ef7d1054e8d2b9()

    return dto;
}

func get0cccdfed4c4840adAe13C1b4160ea624() string {
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
    // dto["detail"] = get5363f4df67b44a51B0d30b90fb0242dc()
    // 设备号
    // dto["device_info"] = ""
    // 订单优惠标记
    // dto["goods_tag"] = ""
    // 实名支付
    // dto["identity"] = ""
    // 开发票入口开放标识
    // dto["receipt"] = ""
    // 场景信息
    // dto["scene_info"] = get8e7acadb53e34c8c849561ed6576e6b1()
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

func getE8d01f1d074c4d33Ab52E96c0ae4aead() string {
    dto := make(map[string]interface{})
    // 汇付机具号
    // dto["devs_id"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

