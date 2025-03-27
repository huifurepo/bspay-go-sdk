/**
 * 分期交易退款 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePayafteruseInstallmentRefundRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePayafteruseInstallmentRefundRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePayafteruseInstallmentRefundRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000108281250",
        // 申请退款金额
        OrdAmt:"0.01",
        // 原请求日期
        OrgReqDate:"20241010",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradePayafteruseInstallmentRefundRequest(dgReq)
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
    // 分账串
    // extendInfoMap["acct_split_bunch"] = get1313ce8b81ce4f5493a623949b89ffd3()
    // 原请求流水号
    extendInfoMap["org_req_seq_id"] = "20241010test10000111qccrr"
    // 原全局流水号
    // extendInfoMap["org_hf_seq_id"] = ""
    // 交易备注
    // extendInfoMap["remark"] = ""
    // 异步通知地址
    // extendInfoMap["notify_url"] = ""
    return extendInfoMap
}

func get5864fd7678e24596B3ce38dba15bb024() interface{} {
    dto := make(map[string]interface{})
    // 分账金额
    // dto["div_amt"] = "test"
    // 分账接收方商户号
    // dto["huifu_id"] = "test"
    // 分账接收方账户号
    // dto["acct_id"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get1313ce8b81ce4f5493a623949b89ffd3() string {
    dto := make(map[string]interface{})
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""
    // 分账明细
    // dto["acct_infos"] = get5864fd7678e24596B3ce38dba15bb024()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

