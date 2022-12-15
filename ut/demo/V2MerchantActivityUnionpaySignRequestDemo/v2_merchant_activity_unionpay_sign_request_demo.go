/**
 * 银联活动报名 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantActivityUnionpaySignRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantActivityUnionpaySignRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantActivityUnionpaySignRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付客户Id
        HuifuId:"6666000103391467",
        // 活动编号
        ActivityId:"306530984470249472",
        // 银联活动商户号
        MerNo:"MH20220707155729AFIIZ",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantActivityUnionpaySignRequest(dgReq)
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
    // 报名补充说明
    extendInfoMap["remark"] = "报名补充说明"
    // 报名文本材料
    extendInfoMap["enlist_txt_makings"] = getEnlistTxtMakings()
    // 报名图片材料
    extendInfoMap["enlist_img_makings"] = getEnlistImgMakings()
    return extendInfoMap
}

func getEnlistImgMakings() string {
    dto := make(map[string]interface{})
    // 活动材料编号
    dto["makings_id"] = "18"
    // 活动材料类型
    dto["makings_type"] = "IMG"
    // 活动材料名称
    dto["makings_name"] = "门头照片"
    // 材料值
    dto["makings_value"] = "42204258-967e-373c-88d2-1afa4c7bb8ef"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getEnlistTxtMakings() string {
    dto := make(map[string]interface{})
    // 活动材料编号
    dto["makings_id"] = "17"
    // 活动材料类型
    dto["makings_type"] = "TXT"
    // 活动材料名称
    dto["makings_name"] = "银联云闪付商户号"
    // 材料值
    dto["makings_value"] = "82339SP5411019L"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

