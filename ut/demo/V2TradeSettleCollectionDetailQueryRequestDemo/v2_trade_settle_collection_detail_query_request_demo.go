/**
 * 资金归集明细查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeSettleCollectionDetailQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeSettleCollectionDetailQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeSettleCollectionDetailQueryRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 开始日期
        BeginDate:"20240711",
        // 结束日期
        EndDate:"20240718",
        // 转出方商户号转出方商户号和转入方商户号必填一个；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123124&lt;/font&gt;
        OutHuifuId:"",
        // 转入方商户号转出方商户号和转入方商户号必填一个；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123124&lt;/font&gt;
        InHuifuId:"6666000146178696",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeSettleCollectionDetailQueryRequest(dgReq)
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
    // 分页页码
    extendInfoMap["page_num"] = "0"
    // 分页条数
    extendInfoMap["page_size"] = "5"
    // 归集状态
    extendInfoMap["status"] = ""
    // 转出方账户号
    extendInfoMap["out_acct_id"] = ""
    // 转入方账户号
    extendInfoMap["in_acct_id"] = ""
    return extendInfoMap
}

