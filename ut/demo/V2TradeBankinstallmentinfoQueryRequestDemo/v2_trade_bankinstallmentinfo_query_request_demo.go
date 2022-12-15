/**
 * 银行卡分期支持银行查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeBankinstallmentinfoQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeBankinstallmentinfoQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeBankinstallmentinfoQueryRequest{
        // 页码
        PageNum:"3",
        // 每页条数
        PageSize:"1",
        // 产品号
        // ProductId:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeBankinstallmentinfoQueryRequest(dgReq)
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
    // 银行编码
    extendInfoMap["bank_code"] = ""
    // 银行名称
    extendInfoMap["bank_name"] = ""
    // 是否启用
    extendInfoMap["bank_enable"] = ""
    return extendInfoMap
}

