/**
 * 服务商终端列表查询 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TerminaldeviceManageQueryRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TerminaldeviceManageQueryRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TerminaldeviceManageQueryRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TerminaldeviceManageQueryRequest(dgReq)
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
    // 渠道商号
    extendInfoMap["upper_huifu_id"] = "6666000104633228"
    // 终端号
    // extendInfoMap["deviceId"] = ""
    // 绑定状态
    extendInfoMap["is_bind"] = "Y"
    // 当前页码
    extendInfoMap["page_num"] = "1"
    // 每页条数
    extendInfoMap["page_size"] = "1"
    return extendInfoMap
}

