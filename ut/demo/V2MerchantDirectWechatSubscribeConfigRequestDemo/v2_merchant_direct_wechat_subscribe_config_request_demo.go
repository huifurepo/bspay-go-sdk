/**
 * 微信直连-微信关注配置 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantDirectWechatSubscribeConfigRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantDirectWechatSubscribeConfigRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantDirectWechatSubscribeConfigRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付ID
        HuifuId:"6666000003099420",
        // 开发者的应用ID
        AppId:"wx3767c5bd01df5061",
        // 商户号
        MchId:"1552470931",
        // 特约商户号
        SubMchid:"10888880",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantDirectWechatSubscribeConfigRequest(dgReq)
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
    // 绑定APPID配置
    extendInfoMap["bind_app_id_conf_list"] = getBindAppIdConfList()
    // 关注配置
    extendInfoMap["subscribe_conf_list"] = getSubscribeConfList()
    // 支付目录配置
    extendInfoMap["pay_path_conf_list"] = getPayPathConfList()
    return extendInfoMap
}

func getBindAppIdConfList() string {
    dto := make(map[string]interface{})
    // 关联APPID
    dto["sub_appid"] = "oQOa46X2FxRqEy6F4YmwIRCrA7Mk"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getSubscribeConfList() string {
    dto := make(map[string]interface{})
    // 关联APPID
    dto["sub_appid"] = "wx5934540532"
    // 推荐关注APPID服务商的公众号APPID；与receipt_appid二选一；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx5934540532&lt;/font&gt;
    dto["subscribe_appid"] = "oQOa46X2FxRqEy6F4YmwIRCrA7Mk"
    // 支付凭证推荐小程序appid需为通过微信认证的小程序appid，且认证主体与服务商主体一致；与subscribe_appid二选一；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：wx852a790f100000fe&lt;/font&gt;
    dto["receipt_appid"] = "wx852a790f100000fe"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getPayPathConfList() string {
    dto := make(map[string]interface{})
    // 授权目录
    dto["jsapi_path"] = "http://www.dsf.com/init"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

