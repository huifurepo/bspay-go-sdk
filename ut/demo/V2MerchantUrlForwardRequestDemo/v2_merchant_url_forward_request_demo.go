/**
 * 商户统一进件（页面版） - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantUrlForwardRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantUrlForwardRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantUrlForwardRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 渠道商号汇付分配的渠道商号，该商户进件完成后将隶属该渠道商/代理商&lt;br/&gt;sys_id主体为渠道商时填写； &lt;br/&gt;sys_id主体为总部商户时选填；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123123&lt;/font&gt;
        UpperHuifuId:"6666000108290240",
        // 门店号
        StoreId:"store128",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantUrlForwardRequest(dgReq)
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
    // 手机号
    extendInfoMap["phone"] = "13917352618"
    // 跳转地址失效时间
    extendInfoMap["expires"] = "50000"
    // 返回页面URL
    // extendInfoMap["back_page_url"] = ""
    // 异步接收URL
    // extendInfoMap["async_receive_url"] = ""
    // 模版编号
    extendInfoMap["template_id"] = "155625"
    // 商户身份
    // extendInfoMap["head_office_flag"] = ""
    // 上级汇付Id
    // extendInfoMap["head_huifu_id"] = ""
    // 使用上级微信、支付宝商户号发起交易
    // extendInfoMap["use_chains_flag"] = ""
    // 小微经营类型
    // extendInfoMap["micro_biz_type"] = ""
    return extendInfoMap
}

