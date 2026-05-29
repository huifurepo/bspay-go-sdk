/**
 * 防断链入驻 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantAtpreventApplyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantAtpreventApplyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantAtpreventApplyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户汇付Id
        HuifuId:"6666000108460751",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantAtpreventApplyRequest(dgReq)
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
    // 异步通知地址
    extendInfoMap["async_url"] = "http://service.example.com/to/path"
    // 微信开通明细
    extendInfoMap["wx_open_list"] = get4ee5a21dCd974031Aaff86bc6a77ad6e()
    // 支付宝开通明细
    extendInfoMap["ali_open_list"] = getAf3f6e6aB58f47f99761E0107d9733be()
    return extendInfoMap
}

func get4ee5a21dCd974031Aaff86bc6a77ad6e() interface{} {
    dto := make(map[string]interface{})
    // 渠道号
    dto["pay_channel_id"] = "10000001"
    // 线上开通数
    dto["online_open_count"] = "1"
    // 线下开通数
    dto["offline_open_count"] = "1"

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getAf3f6e6aB58f47f99761E0107d9733be() interface{} {
    dto := make(map[string]interface{})
    // 渠道号
    dto["pay_channel_id"] = "10000001"
    // 线上开通数
    dto["online_open_count"] = "1"
    // 线下开通数
    dto["offline_open_count"] = "1"

    dtoList := [1]interface{}{dto}
    return dtoList
}

