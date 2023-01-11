/**
 * 商户分账配置 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantSplitConfigRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantSplitConfigRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantSplitConfigRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 商户汇付Id
        HuifuId:"6666000105582434",
        // 分账规则来源
        RuleOrigin:"02",
        // 分账开关
        DivFlag:"Y",
        // 最大分账比例
        ApplyRatio:"100",
        // 生效类型
        StartType:"0",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantSplitConfigRequest(dgReq)
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
    // 分账明细
    extendInfoMap["acct_split_bunch_list"] = getAcctSplitBunchList()
    // 交易手续费外扣开关
    extendInfoMap["out_fee_flag"] = "1"
    // 交易手续费外扣时的账户类型
    extendInfoMap["out_fee_acct_type"] = "01"
    // 交易手续费外扣汇付ID
    extendInfoMap["out_fee_huifuid"] = "6666000105582434"
    // 固定手续费
    extendInfoMap["split_fee_rate"] = "10.89"
    // 百分比手续费
    extendInfoMap["per_amt"] = "99"
    // 异步地址
    extendInfoMap["async_return_url"] = "http://192.168.85.157:30031/sspm/testVirgo"
    return extendInfoMap
}

func getAcctSplitBunchList() string {
    dto := make(map[string]interface{})
    // 分账比例
    dto["fee_rate"] = "100"
    // 汇付Id
    dto["huifu_id"] = "6666000105582434"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

