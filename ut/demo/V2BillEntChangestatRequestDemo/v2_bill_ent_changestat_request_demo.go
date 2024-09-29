/**
 * 企业账单状态变更 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2BillEntChangestatRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2BillEntChangestatRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2BillEntChangestatRequest{
        // 请求流水号
        // ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        // ReqDate:tool.GetCurrentDate(),
        // 商户号
        // HuifuId:"test",
        // 账单编号
        BillNo:"ZD2024082686348233",
        // 变更状态
        // BillStat:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2BillEntChangestatRequest(dgReq)
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
    // 变更原因
    extendInfoMap["remark"] = "测试"
    return extendInfoMap
}

