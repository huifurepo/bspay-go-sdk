/**
 * 创建修改小票自定义入口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeElectronReceiptsCustomentrancesCreateRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeElectronReceiptsCustomentrancesCreateRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeElectronReceiptsCustomentrancesCreateRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000103334211",
        // 操作类型
        OperateType:"A",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeElectronReceiptsCustomentrancesCreateRequest(dgReq)
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
    // 票据信息
    extendInfoMap["receipt_data"] = getReceiptData()
    return extendInfoMap
}

func getJumpLinkData() interface{} {
    dto := make(map[string]interface{})
    // 商家小程序AppID
    dto["mini_programs_app_id"] = "oBmItsxLKa6pd5dSHK4xRLXTt05M"
    // 商家小程序path
    dto["mini_programs_path"] = "https://wxpaylogo.qpic.cn/wxpaylogo/PiajxSqBRaEIPAeia7ImvtsoMpdQ8uEd23s8VtfKDXa04FZk8kXDeH9Q/0"

    return dto;
}

func getWxReceiptData() interface{} {
    dto := make(map[string]interface{})
    // 品牌ID
    dto["brand_id"] = "1"
    // 自定义入口种类
    dto["custom_entrance_type"] = "MERCHANT_ACTIVITY"
    // 副标题
    dto["sub_title"] = "1"
    // 商品缩略图URL
    dto["goods_thumbnail_url"] = "1"
    // 入口展示开始时间
    dto["start_time"] = "2023-08-17T13:20:00+08:00"
    // 入口展示结束时间
    dto["end_time"] = "2023-08-18T11:20:00+08:00"
    // 自定义入口状态
    dto["custom_entrance_state"] = "ONLINE"
    // 请求业务单据号
    dto["out_request_no"] = "1"
    // 跳转信息
    dto["jump_link_data"] = getJumpLinkData()

    return dto;
}

func getReceiptData() string {
    dto := make(map[string]interface{})
    // 三方通道类型
    dto["third_channel_type"] = "T"
    // 微信票据信息
    dto["wx_receipt_data"] = getWxReceiptData()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

