/**
 * 银行大额未入账流水列表查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeOnlinepaymentTransferBankblotterQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeOnlinepaymentTransferBankblotterQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeOnlinepaymentTransferBankblotterQueryRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000003100615",
        // 原请求流水号
        OrgReqSeqId:"2021091708126665001",
        // 原请求日期
        OrgReqDate:"20231215",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeOnlinepaymentTransferBankblotterQueryRequest(dgReq)
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
    // 实际付款方银行卡号
    // extendInfoMap["bank_card_no"] = ""
    // 实际付款方姓名
    extendInfoMap["certificate_name"] = "沈显龙"
    // 实际付款日期
    // extendInfoMap["trans_date"] = ""
    // 交易金额
    // extendInfoMap["trans_amt"] = ""
    // 收款账号/打款备注
    // extendInfoMap["bank_remark"] = ""
    return extendInfoMap
}

