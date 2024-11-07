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
        MiniappData:getMiniappDataRucan(),
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
    // 收银台样式
    // extendInfoMap["style_id"] = ""
    // 是否延迟交易
    extendInfoMap["delay_acct_flag"] = "Y"
    // 分账对象
    extendInfoMap["acct_split_bunch"] = getAcctSplitBunchRucan()
    // 交易失效时间
    // extendInfoMap["time_expire"] = ""
    // 业务信息
    // extendInfoMap["biz_info"] = getBizInfo()
    // 交易异步通知地址
    extendInfoMap["notify_url"] = "https://callback.service.com/xx"
    // 微信参数集合
    // extendInfoMap["wx_data"] = getWxData()
    return extendInfoMap
}

func getAcctInfosRucan() interface{} {
    dto := make(map[string]interface{})
    // 分账金额
    dto["div_amt"] = "0.01"
    // 分账接收方ID
    dto["huifu_id"] = "6666000109133323"

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getAcctSplitBunchRucan() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = getAcctInfosRucan()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getMiniappDataRucan() string {
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

func getPayCheckWx() interface{} {
    dto := make(map[string]interface{})
    // 指定支付者
    // dto["limit_payer"] = ""
    // 微信实名验证
    // dto["real_name_flag"] = ""

    return dto;
}

func getPersonPayer() interface{} {
    dto := make(map[string]interface{})
    // 姓名
    // dto["name"] = ""
    // 证件类型
    // dto["cert_type"] = ""
    // 证件号
    // dto["cert_no"] = ""

    return dto;
}

func getBizInfo() string {
    dto := make(map[string]interface{})
    // 付款人验证（微信）
    // dto["payer_check_wx"] = getPayCheckWx()
    // 个人付款人信息
    // dto["person_payer"] = getPersonPayer()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getGoodsDetailWxRucan() interface{} {
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

func getDetail() interface{} {
    dto := make(map[string]interface{})
    // 单品列表
    // dto["goods_detail"] = getGoodsDetailWxRucan()
    // 订单原价(元)
    // dto["cost_price"] = ""
    // 商品小票ID
    // dto["receipt_id"] = ""

    return dto;
}

func getStoreInfo() interface{} {
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

func getSceneInfo() interface{} {
    dto := make(map[string]interface{})
    // 门店信息
    // dto["store_info"] = getStoreInfo()

    return dto;
}

func getWxData() string {
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
    // dto["detail"] = getDetail()
    // 设备号
    // dto["device_info"] = ""
    // 订单优惠标记
    // dto["goods_tag"] = ""
    // 实名支付
    // dto["identity"] = ""
    // 开发票入口开放标识
    // dto["receipt"] = ""
    // 场景信息
    // dto["scene_info"] = getSceneInfo()
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

