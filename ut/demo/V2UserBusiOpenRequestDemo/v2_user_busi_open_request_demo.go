/**
 * 用户业务入驻 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2UserBusiOpenRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2UserBusiOpenRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2UserBusiOpenRequest{
        // 汇付ID
        HuifuId:"6666000105765113",
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 渠道商/商户汇付Id
        UpperHuifuId:"6666000003084836",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2UserBusiOpenRequest(dgReq)
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
    // 结算信息配置
    // extendInfoMap["&lt;span class&#x3D;&quot;extend settle_config&quot;&gt;settle_config&lt;/span&gt;"] = ""
    // 结算卡信息
    // extendInfoMap["&lt;span class&#x3D;&quot;extend card_info&quot;&gt;card_info&lt;/span&gt;"] = ""
    // 取现配置列表
    // extendInfoMap["&lt;span class&#x3D;&quot;extend cash_config&quot;&gt;cash_config&lt;/span&gt;"] = ""
    // 文件列表
    // extendInfoMap["&lt;span class&#x3D;&quot;extend file_list&quot;&gt;file_list&lt;/span&gt;"] = ""
    // 延迟入账开关
    // extendInfoMap["delay_flag"] = ""
    return extendInfoMap
}

