/**
 * 全渠道资金付款到账户退款 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2EfpAcctpaymentRefundRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2EfpAcctpaymentRefundRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2EfpAcctpaymentRefundRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付商户号
        HuifuId:"6666000123123123",
        // 原交易全局流水号org_hf_seq_id和org_req_seq_id二选一； &lt;font color&#x3D;&quot;green&quot;&gt;示例值：00470topo1A211015160805P090ac132fef00000&lt;/font&gt;
        OrgHfSeqId:"00470topo1A211015160805P090ac132fef00000",
        // 原交易请求流水号org_hf_seq_id和org_req_seq_id二选一；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：2021091708126665002&lt;/font&gt;
        OrgReqSeqId:"2021091708126665002",
        // 原交易请求日期
        OrgReqDate:"20221022",
        // 退款金额
        RefundAmt:"10.00",
        // 接收方退款对象
        AcctSplitBunch:get5bf3321f6c604d67Ad10C58a4cce226c(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2EfpAcctpaymentRefundRequest(dgReq)
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
    extendInfoMap["remark"] = "备注"
    return extendInfoMap
}

func getE1cff61a45374e07B6a5Fd58a85a8f13() interface{} {
    dto := make(map[string]interface{})
    // 退款金额
    dto["div_amt"] = "1.00"
    // 退款方ID
    dto["huifu_id"] = "6666000123123123"
    // 退款方账户号
    dto["acct_id"] = "F00598600"

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get5bf3321f6c604d67Ad10C58a4cce226c() string {
    dto := make(map[string]interface{})
    // 退账明细
    dto["acct_infos"] = getE1cff61a45374e07B6a5Fd58a85a8f13()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

