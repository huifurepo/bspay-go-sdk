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
        BankCardNoCrypt:"b9LE5RccVVLChrHgo9lvpLB1XIyJlEeETa1APmkRQ35z06zJ8zD7cnqypNSnA8iK3uAYVDJtCfrz1Hqu1qTCdu5eVWkjBYaAUtuy1ZD4HkEkqbY9/z5lN4jdDyF8xlzonfxhxzm3OM1fWRoYl39Te+pW71ag0SSbQGu6yhWzFD9mBllbj2RR5fWm9BZVtJTLmitIO/HZfirXkRiCPHBjosQJm2bCrVSuzxqJgqmB9Cp1ADIB+f7fG1/G8RElkJ5zyqhDyinlB5b2+fy3hoyuPqB44GCSLEeOF8V0C9uMNNVor1DwvPRLYleNSw43lW4mFx4PhWhjKrWg2NPfbe0mkQ==",
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

