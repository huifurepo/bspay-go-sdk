/**
 * 不明来账处理 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePaymentZxeUnknownincomeDisposeRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePaymentZxeUnknownincomeDisposeRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePaymentZxeUnknownincomeDisposeRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000109133323",
        // 银行侧交易流水号参照异步通知里的bank_serial_no；&lt;br/&gt;“银行侧交易流水号”和“来账银行账号，来账账户名称，交易金额，交易日期”二选一必填。
        BankSerialNo:"FRSC202409252NEA000121452600000",
        // 来账银行账号需要密文传输，使用汇付RSA公钥加密(加密前64位，加密后最长2048位），参见[参考文档](https://paas.huifu.com/open/doc/guide/#/api_jiami_jiemi)；示例值：Ly+fnExeyPOTzfOtgRRur77nJB9TAe4PGgK9M……fc6XJXZss&#x3D;“银行侧交易流水号”和“来账银行账号，来账账户名称，交易金额，交易日期”二选一必填。
        // PayAcct:"test",
        // 来账账户名称“银行侧交易流水号”和“来账银行账号，来账账户名称，交易金额，交易日期”二选一必填。
        // PayAcctName:"test",
        // 交易金额“银行侧交易流水号”和“来账银行账号，来账账户名称，交易金额，交易日期”二选一必填。
        // TransAmt:"test",
        // 交易日期“银行侧交易流水号”和“来账银行账号，来账账户名称，交易金额，交易日期”二选一必填。
        // TransDate:"test",
        // 操作类型
        OperateType:"0",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradePaymentZxeUnknownincomeDisposeRequest(dgReq)
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
    // 交易时间
    // extendInfoMap["trans_time"] = ""
    // 异步通知地址
    extendInfoMap["notify_url"] = "https://mock.uutool.cn/fat69kri74k"
    return extendInfoMap
}

