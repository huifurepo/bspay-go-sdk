/**
 * 跳转电子小票页面 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeElectronReceiptsJumpinfoRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeElectronReceiptsJumpinfoRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeElectronReceiptsJumpinfoRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 商户汇付Id
        HuifuId:"6666000103334211",
        // 原请求日期
        OrgReqDate:"20230517",
        // 原请求流水号原请求流水号、原交易返回的全局流水号至少要送其中一项；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：2021091708126665001&lt;/font&gt;
        OrgReqSeqId:"20230517111710E83514",
        // 汇付全局流水号原请求流水号、原交易返回的全局流水号至少要送其中一项；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：00290TOP1GR210919004230P853ac13262200000&lt;/font&gt;
        OrgHfSeqId:"0036000topB230517111710P034c0a8304100000",
        // 票据信息
        ReceiptData:getReceiptDataRucan(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeElectronReceiptsJumpinfoRequest(dgReq)
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

func getJumpInfo() interface{} {
    dto := make(map[string]interface{})
    // 小票跳转信息小程序AppID
    dto["merchant_app_id"] = "wxcaced8415a866378"
    // 小票跳转信息小程序路径
    dto["merchant_path"] = "pages/cashier/paySuccess"

    return dto;
}

func getWxReceiptDataRucan() interface{} {
    dto := make(map[string]interface{})
    // 跳转信息
    dto["jump_info"] = getJumpInfo()

    return dto;
}

func getReceiptDataRucan() string {
    dto := make(map[string]interface{})
    // 三方通道类型
    dto["third_channel_type"] = "T"
    // 微信票据信息
    dto["wx_receipt_data"] = getWxReceiptDataRucan()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

