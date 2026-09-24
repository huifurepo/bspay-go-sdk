/**
 * 发货信息录入 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V3TradeHostingPaymentShippinginfoUploadRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V3TradeHostingPaymentShippinginfoUploadRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V3TradeHostingPaymentShippinginfoUploadRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000103334211",
        // 订单信息
        OrderKey:get7c40df27Fc134b61B0c7Edce6d8f621b(),
        // 物流模式
        LogisticsType:1,
        // 发货模式
        DeliveryMode:1,
        // 否已全部发货完成分拆发货模式时必填，用于标识分拆发货模式下是否已全部发货完成，只有全部发货完成的情况下才会向用户推送发货完成通知。&lt;br/&gt;枚举值: true，false&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：true&lt;/font&gt;
        // IsAllDelivered:"test",
        // 上传时间
        UploadTime:"20260831101212",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V3TradeHostingPaymentShippinginfoUploadRequest(dgReq)
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
    // 物流信息列表
    extendInfoMap["shipping_list"] = get12b9768650294e6cAafb18f939d1ae02()
    // 支付者
    extendInfoMap["payer"] = getD4934e1e62bc4c09B037875617dbf057()
    return extendInfoMap
}

func get7c40df27Fc134b61B0c7Edce6d8f621b() string {
    dto := make(map[string]interface{})
    // 订单单号类型
    dto["order_number_type"] = 2
    // 原支付交易对应的微信订单号
    dto["transaction_id"] = "fake-transid-20221214190427-1"
    // 支付下单商户的商户号
    // dto["mchid"] = ""
    // 商户系统内部订单号
    // dto["out_trade_no"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getC48dfa9850104f4c83c360f466f5fb81() interface{} {
    dto := make(map[string]interface{})
    // 寄件人联系方式
    dto["consignor_contact"] = "+86-177****1234"
    // 收件人联系方式
    // dto["receiver_contact"] = ""

    return dto;
}

func get12b9768650294e6cAafb18f939d1ae02() string {
    dto := make(map[string]interface{})
    // 商品信息
    dto["item_desc"] = "微信气泡狗集线器*1"
    // 物流单号
    dto["tracking_no"] = "fake-trackingno-2022121419042711"
    // 物流公司编码，快递公司ID
    dto["express_company"] = "STO"
    // 联系方式
    dto["contact"] = getC48dfa9850104f4c83c360f466f5fb81()

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getD4934e1e62bc4c09B037875617dbf057() string {
    dto := make(map[string]interface{})
    // 用户标识
    dto["openid"] = "ogqztkPsejM9MQAFfwCQSCi4oNg3"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

