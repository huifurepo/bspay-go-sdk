/**
 * 灵工微信代发 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeLgwxSurrogateRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeLgwxSurrogateRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeLgwxSurrogateRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 出款方商户号
        HuifuId:"6666000107755175",
        // 支付金额(元)
        CashAmt:"0.11",
        // 代发模式
        SalaryModleType:"1",
        // 落地公司商户号
        BmemberId:"396117173653968220",
        // 子商户应用ID
        SubAppid:"123213",
        // 异步通知地址
        NotifyUrl:"virgo://http://www.gangcai.com",
        // 分账明细
        AcctSplitBunch:get2cc87980007348a7A86e461ee467b2db(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeLgwxSurrogateRequest(dgReq)
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
    // 出款方账户号
    extendInfoMap["acct_id"] = "C02418374"
    return extendInfoMap
}

func get2cc87980007348a7A86e461ee467b2db() string {
    dto := make(map[string]interface{})
    // 用户号
    dto["huifu_id"] = "6666000107979716"
    // 分配金额(元)
    dto["div_amt"] = "0.11"
    // 微信openid
    dto["sub_openid"] = "13232"
    // 转账备注
    dto["remark"] = "灵工代发1"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

