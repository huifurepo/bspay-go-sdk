/**
 * 提交申诉 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantAppealCommonSubmitRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantAppealCommonSubmitRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantAppealCommonSubmitRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户经营模式
        BusinessPattern:"03",
        // 协查单号
        AssistId:"202407021808060672701923328",
        // 申诉单号
        AppealId:"202407021808060674786492416",
        // 商户类型
        MerType:"01",
        // 申诉人姓名
        AppealPersonName:"张三",
        // 申诉人身份证号
        AppealPersonCertNo:"41162719213519",
        // 申诉人联系电话
        AppealPersonPhoneNo:"186234508",
        // 法人姓名
        LegalName:"张三",
        // 法人身份证号
        LegalCertNo:"411627199509123",
        // 法人联系电话
        LegalPhoneNo:"186234502",
        // 商户主营业务
        MainBusiness:"批发零食饮料",
        // 申诉理由
        AppealDesc:"因XX申诉",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantAppealCommonSubmitRequest(dgReq)
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
    // 商户营业执照号
    extendInfoMap["mer_business_license_no"] = "916501042135"
    // app名称
    extendInfoMap["app_name"] = "app名称"
    // app下载链接
    extendInfoMap["app_url"] = "XXX"
    // 电商网址
    extendInfoMap["commerce_url"] = "www.baidu.com"
    // ICP备案号
    extendInfoMap["icp_record_no"] = "京ICP21003632-7"
    // 实际经营地址
    extendInfoMap["business_address"] = "实际经营地址"
    // 营业用地性质
    extendInfoMap["business_location_nature"] = "01"
    // 经营时长
    extendInfoMap["business_time"] = "09:00:00-21:30:00"
    // 经营地段
    extendInfoMap["business_location_type"] = "01"
    // 员工人数
    extendInfoMap["employee_cnt"] = "10"
    // 申诉文件列表
    extendInfoMap["appeal_file_list"] = get0c986afe8b424cad8c63F947666296c8()
    return extendInfoMap
}

func get0c986afe8b424cad8c63F947666296c8() string {
    dto := make(map[string]interface{})
    // 申诉文件名称
    dto["item_name"] = "法人身份证正面"
    // 申诉文件Id
    dto["file_id"] = "fede0ba8-4994-386c-966a-sd23"
    // 申诉文件类型
    dto["file_code"] = "F01"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

