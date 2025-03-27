/**
 * 红字发票开具接口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2InvoiceRedopenRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2InvoiceRedopenRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2InvoiceRedopenRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付商户号
        HuifuId:"6666000107430944",
        // 原发票号码
        OriIvcNumber:"25317000000132667193",
        // 红冲原因
        RedApplyReason:"02",
        // 红冲申请来源
        RedApplySource:"01",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2InvoiceRedopenRequest(dgReq)
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
    // 备注
    extendInfoMap["resv"] = "备注"
    // 开票结果异步通知地址
    extendInfoMap["callback_url"] = "virgo://http://192.168.85.157:30031/sspm/testVirgo"
    return extendInfoMap
}

