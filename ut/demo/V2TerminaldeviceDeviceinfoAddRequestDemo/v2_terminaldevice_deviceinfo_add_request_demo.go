/**
 * 新增终端报备 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TerminaldeviceDeviceinfoAddRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TerminaldeviceDeviceinfoAddRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TerminaldeviceDeviceinfoAddRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000104575213",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TerminaldeviceDeviceinfoAddRequest(dgReq)
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
    // 终端信息
    extendInfoMap["terminal_info_list"] = getTerminalInfoList()
    return extendInfoMap
}

func getTerminalInfoList() string {
    dto := make(map[string]interface{})
    // 终端硬件序列号
    dto["sn"] = "433333"
    // 终端21号文编号
    dto["tusn"] = "J434445679"
    // 终端型号代号
    dto["dev_model_code"] = "01"
    // 终端布放地址
    dto["terminal_address"] = "上海额的发"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

