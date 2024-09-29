/**
 * 新增归集配置 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeSettleCollectionRuleAddRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeSettleCollectionRuleAddRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeSettleCollectionRuleAddRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 转入方商户号
        InHuifuId:"6666000143571659",
        // 转出方商户号
        OutHuifuId:"6666000152758213",
        // 签约人手机号协议类型为电子协议时必填，必须为法人手机号。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：13911111111&lt;/font&gt;
        SignUserMobileNo:"",
        // 协议文件Id协议类型为纸质协议时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e529&lt;/font&gt;
        FileId:"f80a4c17-d7c5-3e31-9e70-daf2bd6be29e",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeSettleCollectionRuleAddRequest(dgReq)
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
    // 转入方账户号
    extendInfoMap["in_acct_id"] = ""
    // 转出方账户号
    extendInfoMap["out_acct_id"] = ""
    // 转出方账户留存金额
    extendInfoMap["remained_amt"] = "1.01"
    // 协议类型
    extendInfoMap["agreement_type"] = "1"
    // 异步消息接收地址
    extendInfoMap["async_return_url"] = "http://service.testexample.com/to/path"
    return extendInfoMap
}

