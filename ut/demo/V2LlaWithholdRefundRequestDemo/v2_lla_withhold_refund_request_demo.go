/**
 * 代运营佣金代扣退款 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2LlaWithholdRefundRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2LlaWithholdRefundRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2LlaWithholdRefundRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 原请求日期
        OrgReqDate:"20250822",
        // 原请求流水号org_hf_seq_id与org_req_seq_id二选一必填。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：2021091708126665001&lt;/font&gt;
        OrgReqSeqId:"",
        // 原全局流水号org_hf_seq_id与org_req_seq_id二选一必填。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：00470topo1A221019132207P068ac1362af00000&lt;/font&gt;
        OrgHfSeqId:"00470topotB250827093537P979c0a8408100000",
        // 代运营汇付id
        AgencyHuifuId:"6666000108967194",
        // 退款金额
        TransAmt:"25.00",
        // 设备信息
        TerminalDeviceData:get401cdbafBf0248b9Bd22Fc29c064ec90(),
        // 安全信息
        RiskCheckData:get5f9a6c9a2f274b58B08b898a909edb95(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2LlaWithholdRefundRequest(dgReq)
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
    // 备注
    extendInfoMap["remark"] = "部分退款看看"
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.baidu.com"
    return extendInfoMap
}

func get401cdbafBf0248b9Bd22Fc29c064ec90() string {
    dto := make(map[string]interface{})
    // 交易设备类型
    dto["device_type"] = "4"
    // 交易设备IP
    dto["device_ip"] = "172.31.11.11"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get5f9a6c9a2f274b58B08b898a909edb95() string {
    dto := make(map[string]interface{})
    // 经度
    // dto["longitude"] = ""
    // 纬度
    // dto["latitude"] = ""
    // 基站地址
    // dto["base_station"] = ""
    // IP地址
    dto["ip_addr"] = "172.31.11.11"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

