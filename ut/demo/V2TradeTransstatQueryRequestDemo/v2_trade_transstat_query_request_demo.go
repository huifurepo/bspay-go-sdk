/**
 * 批量交易状态查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeTransstatQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeTransstatQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeTransstatQueryRequest{
        // 商户号
        HuifuId:"6666000109133323",
        // 页码
        PageNo:"1",
        // 页大小
        PageSize:"4",
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeTransstatQueryRequest(dgReq)
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
    // 请求订单
    extendInfoMap["reqseqid_list"] = "[\"295700155481522176\",\"20221108104817E93140\",\"20221108104800E93135\",\"20221108112153E93750\",\"20221108133737E96102\"]"
    return extendInfoMap
}

