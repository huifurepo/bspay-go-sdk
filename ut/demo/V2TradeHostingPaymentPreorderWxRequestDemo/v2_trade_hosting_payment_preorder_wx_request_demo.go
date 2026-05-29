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
        MiniappData:getFb8d217747794a7a8b44B97d1707840b(),
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
    // extendInfoMap["split_pay_data"] = get33b93d702a424823A9c104c962b0a4d3()
    // 分账对象
    extendInfoMap["acct_split_bunch"] = get41a5e8fcA0a0401aA6f2279cf764b391()
    // 统一收银台扩展参数集合
    // extendInfoMap["hosting_data"] = get74d0313eFbd74bf49eef5f0e71adaf93()
    // 交易失效时间
    // extendInfoMap["time_expire"] = ""
    // 业务信息
    // extendInfoMap["biz_info"] = getB15a1cee66424b8dA524Fb3dcc211aa9()
    // 交易异步通知地址
    extendInfoMap["notify_url"] = "https://callback.service.com/xx"
    // 微信参数集合
    // extendInfoMap["wx_data"] = get3b4a86c0Fb9546d298f37ea8e161c2a0()
    // 设备信息
    // extendInfoMap["terminal_device_data"] = getEc9037987819440fB638B05a05de8932()
    // 手续费场景标识
    // extendInfoMap["fee_sign"] = ""
    // 是否交易手续费分摊
    // extendInfoMap["fee_split_flag"] = ""
    return extendInfoMap
}

func get33b93d702a424823A9c104c962b0a4d3() string {
    dto := make(map[string]interface{})
    // 商户贴息标记
    // dto["fq_mer_discount_flag"] = ""
    // 商户业务信息
    // dto["ali_business_params"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get94fb584696764557940b87b68de6e256() interface{} {
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

func get41a5e8fcA0a0401aA6f2279cf764b391() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = get94fb584696764557940b87b68de6e256()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get74d0313eFbd74bf49eef5f0e71adaf93() string {
    dto := make(map[string]interface{})
    // 项目号
    // dto["project_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getFb8d217747794a7a8b44B97d1707840b() string {
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

func get9e5c93bb1c864987907cB96cc9c1bdbc() interface{} {
    dto := make(map[string]interface{})
    // 指定支付者
    // dto["limit_payer"] = ""
    // 微信实名验证
    // dto["real_name_flag"] = ""

    return dto;
}

func get1584a9c420a6463dA01010944ee35fd9() interface{} {
    dto := make(map[string]interface{})
    // 姓名
    // dto["name"] = ""
    // 证件类型
    // dto["cert_type"] = ""
    // 证件号
    // dto["cert_no"] = ""

    return dto;
}

func getB15a1cee66424b8dA524Fb3dcc211aa9() string {
    dto := make(map[string]interface{})
    // 付款人验证（微信）
    // dto["payer_check_wx"] = get9e5c93bb1c864987907cB96cc9c1bdbc()
    // 个人付款人信息
    // dto["person_payer"] = get1584a9c420a6463dA01010944ee35fd9()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getC17acafc5b524dfbB569D428985d86ea() interface{} {
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

func getD8eff6a7D39741bb81a38ef29fdb841f() interface{} {
    dto := make(map[string]interface{})
    // 单品列表
    // dto["goods_detail"] = getC17acafc5b524dfbB569D428985d86ea()
    // 订单原价(元)
    // dto["cost_price"] = ""
    // 商品小票ID
    // dto["receipt_id"] = ""

    return dto;
}

func getFf2bd73bFe6b4e9eAe5553471ba84021() interface{} {
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

func get620d86e01a934650A7e99dccba659467() interface{} {
    dto := make(map[string]interface{})
    // 门店信息
    // dto["store_info"] = getFf2bd73bFe6b4e9eAe5553471ba84021()

    return dto;
}

func get3b4a86c0Fb9546d298f37ea8e161c2a0() string {
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
    // dto["detail"] = getD8eff6a7D39741bb81a38ef29fdb841f()
    // 设备号
    // dto["device_info"] = ""
    // 订单优惠标记
    // dto["goods_tag"] = ""
    // 实名支付
    // dto["identity"] = ""
    // 开发票入口开放标识
    // dto["receipt"] = ""
    // 场景信息
    // dto["scene_info"] = get620d86e01a934650A7e99dccba659467()
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

func getEc9037987819440fB638B05a05de8932() string {
    dto := make(map[string]interface{})
    // 汇付机具号
    // dto["devs_id"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

