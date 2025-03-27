/**
 * 钱包密码重置 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2WalletPasswordResetRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2WalletPasswordResetRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2WalletPasswordResetRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000107309462",
        // 钱包用户ID
        UserHuifuId:"6666000107355468",
        // 钱包绑定手机号
        CustMobile:"13771817106",
        // 手机短信验证码
        VerifyNo:"652364",
        // 短信验证流水号
        VerifySeqId:"WALLET0000000054024907",
        // 跳转地址
        FrontUrl:"",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2WalletPasswordResetRequest(dgReq)
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
    // 请求失效时间
    extendInfoMap["time_expired"] = ""
    // 个人证件号码
    // extendInfoMap["cert_no"] = ""
    // 银行卡号
    extendInfoMap["card_no"] = ""
    // 银行卡绑定手机号
    // extendInfoMap["card_mobile"] = ""
    return extendInfoMap
}

