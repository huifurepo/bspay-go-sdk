/**
 * 卡bin信息查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeCardbinQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeCardbinQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeCardbinQueryRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 银行卡号密文
        BankCardNoCrypt:"VDFveDULoK6qhWuN34P8XF7tuZow4eg74zEPKjNVwSjQTW572jqmYjpRDbEtX0f56hMQUfFv5wFRjvDIY+yTl+SFKgBmoNlPDm9B3mDjOa8er5FEnI5QWgvyuqSxHWJf2eIjU7OHt3q3/0ZsbetNzIAiaAZdicn9Sawsr9yZ8ZOdBmhIPO5tqJyDkoKO5Tpj0VyZMJ5ugSDx/2XGSmX3dHQ1JKlZ/7rovB+WchA9BzZQzSTcvmmdV6mzuyTPWRzxfhyVPAVLzh5XhKyl6SMKmzOM1zSNMPtzDhTaUG4XLSM5A0+Tqt3O4tSi16f5zn2csgwbN6TMd5jrXNzC8UTpdQ==",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeCardbinQueryRequest(dgReq)
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
    return extendInfoMap
}

