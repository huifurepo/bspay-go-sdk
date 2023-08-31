/**
 * 微信商户配置 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantBusiConfigRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantBusiConfigRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantBusiConfigRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付客户Id
        HuifuId:"6666000108854952",
        // 业务开通类型
        FeeType:"02",
        // 公众号支付Appid条件必填，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx3767c5bd01df5061&lt;/font&gt; ；wx_woa_app_id 、wx_woa_path和 wx_applet_app_id三者不能同时为空
        WxWoaAppId:"wx3767c5bd01df5061",
        // 微信公众号授权目录条件必填，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：https://paas.huifu.com/shouyintai/demo/h5/&lt;/font&gt;；wx_woa_app_id 、wx_woa_path和 wx_applet_app_id三者不能同时为空
        WxWoaPath:"https://paas.huifu.com/shouyin/demo/h5/",
        // 微信小程序APPID条件必填，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx8523175fea790f10&lt;/font&gt; ；wx_woa_app_id 、wx_woa_path和 wx_applet_app_id三者不能同时为空
        WxAppletAppId:"wx8523175fea790f10",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantBusiConfigRequest(dgReq)
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
    // 微信公众号APPID对应的秘钥
    extendInfoMap["wx_woa_secret"] = "64afb60bef3a22ac282aa7880cdaca98"
    // 微信小程序APPID对应的秘钥
    extendInfoMap["wx_applet_secret"] = "1323a4165a662d6e4f9f51b3f7a58e3f"
    // 渠道号
    extendInfoMap["bank_channel_no"] = "JQF00001"
    // 异步消息接收地址
    // extendInfoMap["async_return_url"] = ""
    return extendInfoMap
}

