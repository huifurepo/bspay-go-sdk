/**
 * 发货信息合单录入 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V3TradeHostingPaymentCombinedshippinginfoUploadRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V3TradeHostingPaymentCombinedshippinginfoUploadRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V3TradeHostingPaymentCombinedshippinginfoUploadRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000103334211",
        // 订单信息
        OrderKey:getAa22aa8e4421431bBcbb03bf9ea6b08d(),
        // 上传时间
        UploadTime:"20231112200913",
        // 支付者
        Payer:getAeaf17f6D3934753A5f70c39ebdf7f85(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V3TradeHostingPaymentCombinedshippinginfoUploadRequest(dgReq)
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
    // 子单物流详情
    extendInfoMap["sub_orders"] = get3614704a5b104be1906399e5c7076590()
    return extendInfoMap
}

func getAa22aa8e4421431bBcbb03bf9ea6b08d() string {
    dto := make(map[string]interface{})
    // 订单单号类型
    dto["order_number_type"] = "1"
    // 原支付交易对应的微信订单号
    // dto["transaction_id"] = ""
    // 支付下单商户的商户号
    dto["mchid"] = "fake-mchid-123"
    // 商户系统内部订单号
    dto["out_trade_no"] = "fake-tradeno-20221214190427-0"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getAeaf17f6D3934753A5f70c39ebdf7f85() string {
    dto := make(map[string]interface{})
    // 用户标识
    dto["openid"] = "ogqztkPsejM9MQAFfwCQSCi4oNg3"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getFc965ab6287d405dAb325c90241f7b70() interface{} {
    dto := make(map[string]interface{})
    // 订单单号类型
    dto["order_number_type"] = "1"
    // 原支付交易对应的微信订单号
    // dto["transaction_id"] = ""
    // 支付下单商户的商户号
    dto["mchid"] = "fake-mchid-123"
    // 商户系统内部订单号
    dto["out_trade_no"] = "fake-tradeno-20221214190427-01"

    return dto;
}

func getAfc3fe0753a949889a8b3bac0ed4196a() interface{} {
    dto := make(map[string]interface{})
    // 寄件人联系方式
    dto["consignor_contact"] = "021-**34-12"
    // 收件人联系方式
    // dto["receiver_contact"] = ""

    return dto;
}

func getCc61a8ed21154bdd8231E4e48a2ae31a() interface{} {
    dto := make(map[string]interface{})
    // 商品信息
    dto["item_desc"] = "微信气泡狗零钱包*1"
    // 物流单号
    dto["tracking_no"] = "fake-trackingno-202212141904271"
    // 物流公司编码
    dto["express_company"] = "YD"
    // 联系方式
    dto["contact"] = getAfc3fe0753a949889a8b3bac0ed4196a()

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get3614704a5b104be1906399e5c7076590() string {
    dto := make(map[string]interface{})
    // 需要上传物流详情的子单订单，订单类型与合单订单保持一致
    dto["order_key"] = getFc965ab6287d405dAb325c90241f7b70()
    // 物流模式
    dto["logistics_type"] = "1"
    // 发货模式
    dto["delivery_mode"] = "2"
    // 是否已全部发货完成
    dto["is_all_delivered"] = "true"
    // 子单物流信息列表
    dto["shipping_list"] = getCc61a8ed21154bdd8231E4e48a2ae31a()

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

