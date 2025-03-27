/**
 * 全渠道资金提现申请 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2EfpEncashRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2EfpEncashRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2EfpEncashRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户汇付id
        HuifuId:"6666000108422302",
        // 交易金额.单位:元，2位小数
        CashAmt:"0.01",
        // 取现卡序列号
        TokenNo:"10001933531",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2EfpEncashRequest(dgReq)
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
    // 原请求流水号
    extendInfoMap["efp_req_seq_id"] = "20241128164919defaultit687891821"
    // 原请求日期
    extendInfoMap["efp_req_date"] = "20241128"
    // 备注
    extendInfoMap["remark"] = "saas取现"
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.service.com/getresp"
    return extendInfoMap
}

