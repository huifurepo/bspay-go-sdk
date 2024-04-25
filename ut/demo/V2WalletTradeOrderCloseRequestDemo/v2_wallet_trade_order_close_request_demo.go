/**
 * 钱包关单 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2WalletTradeOrderCloseRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2WalletTradeOrderCloseRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2WalletTradeOrderCloseRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000108134725",
        // 原交易全局流水号org_hf_seq_id，org_req_seq_id二选一；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：00470topo1A221019132207P068ac1362af00000&lt;/font&gt;
        OrgHfSeqId:"01060000000240402102610P904ac1f123400000",
        // 原交易请求流水号org_hf_seq_id，org_req_seq_id二选一；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：rQ2021121311173944134649875651&lt;/font&gt;
        // OrgReqSeqId:"test",
        // 原交易请求日期
        OrgReqDate:"20240402",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2WalletTradeOrderCloseRequest(dgReq)
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

