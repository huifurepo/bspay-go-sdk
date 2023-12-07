/**
 * 创建支付分订单 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePayscoreServiceorderCreateRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePayscoreServiceorderCreateRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePayscoreServiceorderCreateRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户申请单号
        ReqSeqId:tool.GetReqSeqId(),
        // 汇付商户号
        HuifuId:"6666000108854952",
        // 服务信息
        // ServiceIntroduction:"test",
        // 服务风险金
        // RiskFund:getRiskFund(),
        // 服务时间
        // TimeRange:getTimeRange(),
        // 商户回调地址
        // NotifyUrl:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradePayscoreServiceorderCreateRequest(dgReq)
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
    // 服务ID
    // extendInfoMap["service_id"] = ""
    // 子商户公众号ID
    // extendInfoMap["sub_appid"] = ""
    // 场景类型
    // extendInfoMap["trade_scene"] = ""
    // 费率类型
    // extendInfoMap["pay_scene"] = ""
    // 从业机构公众号下的用户标识
    // extendInfoMap["openid"] = ""
    // 子商户公众号下的用户标识
    // extendInfoMap["sub_openid"] = ""
    // 后付费项目
    // extendInfoMap["post_payments"] = getPostPayments()
    // 商户优惠
    // extendInfoMap["post_discounts"] = getPostDiscounts()
    // 服务位置
    // extendInfoMap["location"] = getLocation()
    // 附加数据
    // extendInfoMap["attach"] = ""
    return extendInfoMap
}

func getPostPayments() string {
    dto := make(map[string]interface{})
    // 付费名称
    // dto["name"] = ""
    // 付费金额
    // dto["amount"] = ""
    // 付费说明
    // dto["description"] = ""
    // 付费数量
    // dto["count"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getPostDiscounts() string {
    dto := make(map[string]interface{})
    // 优惠名称
    // dto["name"] = ""
    // 优惠金额
    // dto["amount"] = ""
    // 优惠说明
    // dto["description"] = ""
    // 优惠数量
    // dto["count"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getRiskFund() string {
    dto := make(map[string]interface{})
    // 风险名称
    // dto["name"] = ""
    // 风险金额
    // dto["amount"] = ""
    // 风险说明
    // dto["description"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getTimeRange() string {
    dto := make(map[string]interface{})
    // 服务开始时间
    // dto["start_time"] = ""
    // 服务结束时间
    // dto["end_time"] = ""
    // 服务开始时间备注
    // dto["start_time_remark"] = ""
    // 服务结束时间备注
    // dto["end_time_remark"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getLocation() string {
    dto := make(map[string]interface{})
    // 服务开始地点
    // dto["start_location"] = ""
    // 服务结束地点
    // dto["end_location"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

