/**
 * 抖音卡券校验 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2CouponDouyinPrepareRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2CouponDouyinPrepareRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2CouponDouyinPrepareRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付商户号
        HuifuId:"6666000107767088",
        // 门店绑定流水号
        BindId:"88fd7c9b63e84a259dfe3eecb811fce8",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2CouponDouyinPrepareRequest(dgReq)
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
    // 抖音团购券码
    extendInfoMap["coupon_code"] = "5729740654"
    // 从二维码解析出来的标识（传参前需要先进行URL编码，注意不要有空格)
    // extendInfoMap["encrypted_data"] = ""
    return extendInfoMap
}

