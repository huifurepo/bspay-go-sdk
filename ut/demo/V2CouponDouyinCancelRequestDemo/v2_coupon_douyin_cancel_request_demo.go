/**
 * 抖音卡券撤销 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2CouponDouyinCancelRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2CouponDouyinCancelRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2CouponDouyinCancelRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付商户号
        HuifuId:"6666000107767088",
        // 门店绑定流水号
        BindId:"88fd7c9b63e84a259dfe3eecb811fce8",
        // 加密抖音券码
        EncryptedCode:"5584192259",
        // 核销标识
        VerifyId:"3435",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2CouponDouyinCancelRequest(dgReq)
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
    // 机具id
    // extendInfoMap["device_id"] = ""
    // 操作人id
    // extendInfoMap["operator_id"] = ""
    // 操作人姓名
    // extendInfoMap["operator_name"] = ""
    // 取消核销总次数
    // extendInfoMap["times_card_cancel_count"] = ""
    // 撤销核销幂等
    // extendInfoMap["cancel_token"] = ""
    return extendInfoMap
}

