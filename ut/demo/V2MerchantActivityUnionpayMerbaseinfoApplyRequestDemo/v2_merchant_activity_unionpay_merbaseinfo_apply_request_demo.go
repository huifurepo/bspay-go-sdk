/**
 * 银联活动商户入驻 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantActivityUnionpayMerbaseinfoApplyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantActivityUnionpayMerbaseinfoApplyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantActivityUnionpayMerbaseinfoApplyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付客户Id
        HuifuId:"6666000103391467",
        // 商户类型
        MerType:"MERCHANT_03",
        // 经营类型
        DealType:"01",
        // 所属行业（MCC）MERCHANT_01-自然人 需要传入，参考[银联MCC编码](https://paas.huifu.com/partners/api/#/csfl/api_csfl_ylmccbm) ；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：5311&lt;/font&gt;
        Mcc:"",
        // 负责人手机号
        LegalMobile:"18900458938",
        // 联系人身份证号
        ContractIdNo:"110101199003071874",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantActivityUnionpayMerbaseinfoApplyRequest(dgReq)
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
    // 文件列表
    extendInfoMap["file_list"] = getFileList()
    return extendInfoMap
}

func getFileList() string {
    dto := make(map[string]interface{})
    // 文件类型
    dto["file_type"] = "F02"
    // 文件jfileID
    dto["file_id"] = "e2113798-ee67-30ac-9981-9dd742510adb"
    // 文件名称
    dto["file_name"] = "身份证正面照片"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

