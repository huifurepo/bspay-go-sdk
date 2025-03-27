/**
 * 批量出金交易查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeBatchtranslogQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeBatchtranslogQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeBatchtranslogQueryRequest{
        // 商户号
        HuifuId:"6666000000041651",
        // 开始日期
        BeginDate:"20230315",
        // 结束日期
        EndDate:"20230316",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeBatchtranslogQueryRequest(dgReq)
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
    // 交易类型
    extendInfoMap["batch_trans_type"] = ""
    // 分页页码
    extendInfoMap["page_num"] = "1"
    // 分页条数
    extendInfoMap["page_size"] = "10"
    // 原支付全局流水号
    // extendInfoMap["payment_hf_seq_id"] = ""
    return extendInfoMap
}

