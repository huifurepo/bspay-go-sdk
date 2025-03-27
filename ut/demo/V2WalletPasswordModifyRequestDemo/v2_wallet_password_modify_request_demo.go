/**
 * 钱包密码修改 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2WalletPasswordModifyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2WalletPasswordModifyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2WalletPasswordModifyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000107309462",
        // 钱包用户ID
        UserHuifuId:"6666000107355468",
        // 手机短信验证码
        VerifyNo:"011363",
        // 短信验证流水号
        VerifySeqId:"WALLET0000000054024907",
        // 跳转地址
        FrontUrl:"https://www.huifu.com/products-services/",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2WalletPasswordModifyRequest(dgReq)
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
    return extendInfoMap
}

