/**
 * 美团卡券核销 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2CouponMeituanConsumeRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2CouponMeituanConsumeRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2CouponMeituanConsumeRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付商户号
        HuifuId:"6666000106057033",
        // 门店绑定流水号
        BindId:"9c2d91f68ba045a998df46ffe395a9ca",
        // 核销券
        ReceiptCodeInfos:get1220ae0e014040c9Af4d128bcbe56d8a(),
        // 登录账号
        AppShopAccount:"123",
        // 登录用户名
        AppShopAccountName:"12345",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2CouponMeituanConsumeRequest(dgReq)
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
    return extendInfoMap
}

func get1220ae0e014040c9Af4d128bcbe56d8a() string {
    dto := make(map[string]interface{})
    // 券码
    dto["receipt_code"] = "5729740654"
    // 核销张数
    dto["num"] = "1"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

