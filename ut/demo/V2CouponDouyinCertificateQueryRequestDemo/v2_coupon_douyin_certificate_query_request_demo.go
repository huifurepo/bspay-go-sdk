/**
 * 抖音券状态批量查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2CouponDouyinCertificateQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2CouponDouyinCertificateQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2CouponDouyinCertificateQueryRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付商户号
        HuifuId:"6666000132473423",
        // 门店绑定流水号
        BindId:"7123fc6e9337a5",
        // 验券准备接口返回的加密券码encrypted_code和order_id二选一必传，encrypted_code和order_id不能同时传入
        EncryptedCode:"0",
        // 订单id验券准备等接口获得，encrypted_code和order_id二选一必传，encrypted_code和order_id不能同时传入
        // OrderId:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2CouponDouyinCertificateQueryRequest(dgReq)
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
    return extendInfoMap
}

