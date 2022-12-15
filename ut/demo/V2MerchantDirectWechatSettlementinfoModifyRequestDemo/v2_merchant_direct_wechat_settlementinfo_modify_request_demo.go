/**
 * 微信直连-修改微信结算帐号 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantDirectWechatSettlementinfoModifyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantDirectWechatSettlementinfoModifyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantDirectWechatSettlementinfoModifyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付ID
        HuifuId:"6666000003098550",
        // 开发者的应用ID
        AppId:"wxd2da4051c9e32b86",
        // 微信商户号
        MchId:"1552470931",
        // 特约商户号
        SubMchid:"10888880",
        // 账户类型
        AccountType:"ACCOUNT_TYPE_BUSINESS",
        // 开户银行
        AccountBank:"农业银行",
        // 开户银行省市编码
        BankAddressCode:"310100",
        // 银行账号
        AccountNumber:"6235012141000002900",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantDirectWechatSettlementinfoModifyRequest(dgReq)
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
    // 开户银行全称（含支行）
    extendInfoMap["bank_name"] = "中国农业银行股份有限公司上海马当路支行"
    // 开户银行联行号
    extendInfoMap["bank_branch_id"] = "103290040169"
    return extendInfoMap
}

