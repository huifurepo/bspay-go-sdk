/**
 * 直付通分账关系绑定解绑 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantDirectZftReceiverConfigRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantDirectZftReceiverConfigRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantDirectZftReceiverConfigRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付ID
        HuifuId:"6666000103886183",
        // 开发者的应用ID
        AppId:"2021002171607880",
        // 分账开关
        SplitFlag:"1",
        // 分账接收方列表
        ZftSplitReceiverList:getZftSplitReceiverList(),
        // 状态
        Status:"0",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantDirectZftReceiverConfigRequest(dgReq)
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

func getZftSplitReceiverList() string {
    dto := make(map[string]interface{})
    // 分账接收方方类型
    dto["split_type"] = "loginName"
    // 分账接收方账号
    dto["account"] = "739100190@qq.com"
    // 分账接收方真实姓名新增分账关系时必填。解绑分账关系时非必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：张三&lt;/font&gt;
    dto["name"] = "邵文"
    // 分账关系描述
    dto["memo"] = "M20220820032239499098320"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

