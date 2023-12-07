/**
 * 银行大额转账 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeOnlinepaymentTransferAccountRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeOnlinepaymentTransferAccountRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeOnlinepaymentTransferAccountRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 收款方商户号
        HuifuId:"6666000003100616",
        // 交易金额
        TransAmt:"10.00",
        // 商品描述
        GoodsDesc:"商品描述001",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeOnlinepaymentTransferAccountRequest(dgReq)
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
    // 收款方用户客户号
    extendInfoMap["user_huifu_id"] = "6666000103905031"
    // 收款方账户号
    extendInfoMap["acct_id"] = "A01199826"
    // 付款方名称
    extendInfoMap["certificate_name"] = "沈显龙"
    // 付款方银行卡号
    extendInfoMap["bank_card_no"] = "Xmjm1RB4AAOaFYQ+PgjBlpugXbd8VAYAGB3J2zrbLfC42Bh5xiB47OOV1EdXyGpBq4H8je7mB/MlyEEs6O8PX6aoI4QHumr8VglrLM6uzbVNCIc3S5RPSmi2M+9+EdIQ6nlWd5+XQ7RJXX5Uvnegn74XzQBcN1d4gd04buwKbLpUPV3tWd1qjQwEE8w4gwEtH3L5AP75Mynz+wHFrUKJF3BTiW2/zJlcq5GJomOl06GEW52AZkXwn6U2suP3a0ySd0Rxbf1yQ1lj3SP56NeeEzuBaFLQWB7mEqJfZF3pE9MHNfi6tR1xwLdcxt98bdIqlteKdNAmgfQzcS13UcwH+w=="
    // 订单类型
    // extendInfoMap["order_type"] = ""
    // 备注
    extendInfoMap["remark"] = "标记123"
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.huifu.com/getResp"
    return extendInfoMap
}

