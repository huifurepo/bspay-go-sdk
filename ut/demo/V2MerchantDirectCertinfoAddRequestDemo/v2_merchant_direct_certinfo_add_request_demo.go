/**
 * 证书登记 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantDirectCertinfoAddRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantDirectCertinfoAddRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantDirectCertinfoAddRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 渠道商汇付Id
        UpperHuifuId:"6666000103509367",
        // 开通类型
        PayWay:"W",
        // 开发者的应用ID
        AppId:"20220818198665087",
        // 文件列表
        FileList:getD3445b35252c4ebd855b4f1aeb23cfbc(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantDirectCertinfoAddRequest(dgReq)
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
    // 商户号
    extendInfoMap["mch_id"] = "360634064"
    // 证书序列号
    extendInfoMap["cert_sn"] = "20220818883326714"
    // 服务商密钥
    extendInfoMap["secret_key"] = "RERE202208182319"
    // 证书类型标记
    extendInfoMap["cert_flag"] = ""
    return extendInfoMap
}

func getD3445b35252c4ebd855b4f1aeb23cfbc() string {
    dto := make(map[string]interface{})
    // 文件类型
    dto["file_type"] = "F53"
    // 文件jfileID
    dto["file_id"] = "9aec5b9e-816f-3ebf-8fe8-4146348ce2b0"
    // 文件名称
    dto["file_name"] = "证书1202208189390.crt"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

