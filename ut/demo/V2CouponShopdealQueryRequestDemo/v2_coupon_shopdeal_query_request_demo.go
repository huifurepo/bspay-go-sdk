/**
 * 美团非餐饮获取团购信息 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2CouponShopdealQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2CouponShopdealQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2CouponShopdealQueryRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付商户号
        HuifuId:"6666000106057033",
        // 门店绑定流水号
        BindId:"b924ef97a4c14bbe88c72794fa4ce505",
        // 页码
        Offset:"1",
        // 页大小
        Limit:"100",
        // 售卖平台
        Source:"2",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2CouponShopdealQueryRequest(dgReq)
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
    // 操作人
    extendInfoMap["operator_id"] = ""
    return extendInfoMap
}

