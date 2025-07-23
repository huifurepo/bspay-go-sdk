/**
 * 抖音卡券核销 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2CouponDouyinConsumeRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2CouponDouyinConsumeRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2CouponDouyinConsumeRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付商户号
        HuifuId:"6666000107767088",
        // 门店绑定流水号
        BindId:"88fd7c9b63e84a259dfe3eecb811fce8",
        // 加密抖音券码列表
        EncryptedCodes:"[\"2343\",\"5462\"]",
        // 校验标识
        VerifyToken:"EfdAdS3",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2CouponDouyinConsumeRequest(dgReq)
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
    // 核销额外参数
    // extendInfoMap["verify_extra"] = get11f3d825Aad642ec886f1867af6dfb8e()
    return extendInfoMap
}

func get11f3d825Aad642ec886f1867af6dfb8e() interface{} {
    dto := make(map[string]interface{})
    // 开台时间（秒）
    // dto["biz_time"] = "test"
    // 实际抵扣金额（分））
    // dto["actual_deduction_amount"] = "test"

    return dto;
}

