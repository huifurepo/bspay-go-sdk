/**
 * 完结支付分订单 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePayscoreServiceorderCompleteRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePayscoreServiceorderCompleteRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePayscoreServiceorderCompleteRequest{
        // 汇付商户号
        HuifuId:"6666000108854952",
        // 汇付订单号
        // OutOrderNo:"test",
        // 完结金额
        // OrdAmt:"test",
        // 服务时间
        // TimeRange:getTimeRange(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradePayscoreServiceorderCompleteRequest(dgReq)
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
    // 创建服务订单返回的汇付全局流水号
    // extendInfoMap["org_hf_seq_id"] = ""
    // 服务订单创建请求流水号
    // extendInfoMap["org_req_seq_id"] = ""
    // 后付费项目
    // extendInfoMap["post_payments"] = getPostPayments()
    // 商户优惠
    // extendInfoMap["post_discounts"] = getPostDiscounts()
    // 服务位置
    // extendInfoMap["location"] = getLocation()
    // 完结服务时间
    // extendInfoMap["complete_time"] = ""
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

