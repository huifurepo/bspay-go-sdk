/**
 * 灵工支付 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2FlexibleTradeRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2FlexibleTradeRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2FlexibleTradeRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 出款方商户号
        OutHuifuId:"6666000108903745",
        // 出款方账户号
        OutAcctId:"C03117654",
        // 交易阶段操作类型
        StageOperationType:"FIRST_STAGE",
        // 前段交易流水号** 当交易阶段操作类型为02时，该字段必填。填写的是交易阶段操作类型为01时交易已完成的交易全局流水号。 &lt;font color&#x3D;&quot;green&quot;&gt;示例值：20250620112533115566896&lt;/font&gt;
        PhaseHfSeqId:"",
        // 支付金额
        OrdAmt:"20",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2FlexibleTradeRequest(dgReq)
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
    extendInfoMap["remark"] = ""
    // 分账对象
    extendInfoMap["acct_split_bunch"] = getD5cc6c3fD3854f9fB3eeF736df9fbbf8()
    return extendInfoMap
}

func get3fc17817Caf445dc8f13A2c315f6d1e8() interface{} {
    dto := make(map[string]interface{})
    // 分账金额
    dto["div_amt"] = "20.00"
    // 分账接收方ID
    dto["huifu_id"] = "6666000108898793"
    // 账户号
    dto["acct_id"] = "C03113649"

    return dto;
}

func getD5cc6c3fD3854f9fB3eeF736df9fbbf8() interface{} {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_info"] = get3fc17817Caf445dc8f13A2c315f6d1e8()

    return dto;
}

