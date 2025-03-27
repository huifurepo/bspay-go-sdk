/**
 * 服务单创建 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePayafteruseCreditbizorderCreateRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePayafteruseCreditbizorderCreateRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePayafteruseCreditbizorderCreateRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000108281250",
        // 订单总金额
        TransAmt:"0.01",
        // 支付宝用户ID
        BuyerId:"2088000000000000",
        // 订单标题
        Title:"测试001",
        // 订单类型
        MerchantBizType:"INDIRECT_CHARGE_WITHHOLD",
        // 订单详情地址
        Path:"https://www.baidu.com/",
        // 芝麻信用服务ID
        ZmServiceId:"2024081500001003000081751200",
        // 商品详细信息
        ItemInfos:get864f5a50D5064cea9423A5f3ca9e73a7(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradePayafteruseCreditbizorderCreateRequest(dgReq)
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
    // 异步通知地址
    extendInfoMap["notify_url"] = "https://mock.uutool.cn/4pga0jqv8vv0"
    // 追踪ID
    extendInfoMap["source_id"] = "MjA4ODcwMjY5OTkwODI1N3wyMDIxMDAzMTUwNjM4MTE2fDE3Mjg1NDk3OTU0OTl8ZmFsc2V8VE9LRU5fSVNfTlVMTA=="
    // 支付宝交易号
    extendInfoMap["trade_no"] = "2024092722001408251414114417"
    // 代扣协议签约场景
    extendInfoMap["deduct_sign_scene"] = "INDUSTRY|XIANXIANG_BIKE_CHARGE"
    // 芝麻信用外部类⽬
    extendInfoMap["zm_category_id"] = "credit_pay_for_battery_charging"
    // 是否不需要核身
    // extendInfoMap["no_need_verify_identity"] = ""
    // 开通成功后跳转地址
    // extendInfoMap["acceptance_jump_url"] = ""
    return extendInfoMap
}

func get1efd507a9385411f80f998b1c37876d9() interface{} {
    dto := make(map[string]interface{})
    // 总分期数
    dto["period_num"] = 1
    // 每期最大金额
    dto["period_max_price"] = 0.30
    // 每期金额
    // dto["period_price"] = ""

    return dto;
}

func get864f5a50D5064cea9423A5f3ca9e73a7() string {
    dto := make(map[string]interface{})
    // 商户商品ID
    dto["out_item_id"] = "1234567"
    // 商品名称
    dto["goods_name"] = "快充"
    // 商品数量
    dto["item_cnt"] = "1"
    // 商品单价
    dto["sale_price"] = "0.30"
    // 商品的编号
    dto["goods_id"] = "Ldkc00001"
    // 商品分期信息
    dto["item_installment_info"] = get1efd507a9385411f80f998b1c37876d9()

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

