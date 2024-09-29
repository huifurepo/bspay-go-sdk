/**
 * 修改付款人信息 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2BillEntPayerUpdateRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2BillEntPayerUpdateRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2BillEntPayerUpdateRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000003100615",
        // 付款人
        PayerId:"P2024082786373612",
        // 付款人名称
        PayerName:"史凡",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2BillEntPayerUpdateRequest(dgReq)
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
    // 付款人邮箱地址
    extendInfoMap["payer_email"] = "1111@163.com"
    // 付款人手机号码
    extendInfoMap["payer_mobile_no"] = "17611111111"
    return extendInfoMap
}

