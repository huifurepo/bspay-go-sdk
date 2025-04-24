/**
 * 交易确认退款 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePaymentDelaytransConfirmrefundRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePaymentDelaytransConfirmrefundRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePaymentDelaytransConfirmrefundRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000109133323",
        // 原交易请求日期
        OrgReqDate:"20240513",
        // 原交易请求流水号
        OrgReqSeqId:"20240513105825239x0lp7ldbus4sji",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradePaymentDelaytransConfirmrefundRequest(dgReq)
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
    // 分账对象
    extendInfoMap["acct_split_bunch"] = get52147df86c694d86B9305a431b675f29()
    // 是否垫资退款
    // extendInfoMap["loan_flag"] = ""
    // 垫资承担者
    // extendInfoMap["loan_undertaker"] = ""
    // 垫资账户类型
    // extendInfoMap["loan_acct_type"] = ""
    // 备注
    // extendInfoMap["remark"] = ""
    return extendInfoMap
}

func get8b3eddeaB8474c7f905e858f0050be27() interface{} {
    dto := make(map[string]interface{})
    // 分账接收方ID
    dto["huifu_id"] = "6666000109133323"
    // 分账金额(元)
    dto["div_amt"] = "0.01"
    // 垫资金额(元)
    // dto["part_loan_amt"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get52147df86c694d86B9305a431b675f29() string {
    dto := make(map[string]interface{})
    // 分账明细
    dto["acct_infos"] = get8b3eddeaB8474c7f905e858f0050be27()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

