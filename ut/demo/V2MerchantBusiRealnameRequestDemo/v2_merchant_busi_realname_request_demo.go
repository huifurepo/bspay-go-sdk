/**
 * 微信实名认证 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantBusiRealnameRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantBusiRealnameRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantBusiRealnameRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付ID
        HuifuId:"6666000104854510",
        // 联系人姓名
        Name:"小枫",
        // 联系人手机号
        Mobile:"17521205027",
        // 联系人身份证号码
        IdCardNumber:"130224198806083798",
        // 联系人类型
        ContactType:"SUPER",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantBusiRealnameRequest(dgReq)
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
    // 子渠道号
    extendInfoMap["pay_channel_id"] = "JP00001"
    // 支付场景
    extendInfoMap["pay_scene"] = "01"
    // 经营者/法人是否为受益人
    extendInfoMap["owner"] = "N"
    // 法人证件居住地址
    extendInfoMap["identification_address"] = "上海市徐汇区宜山路789号789室"
    // 受益人信息
    extendInfoMap["ubo_info_list"] = getUboInfoList()
    // 联系人证件类型
    extendInfoMap["contact_id_doc_type"] = "01"
    // 联系人证件有效期开始时间
    extendInfoMap["contact_period_begin"] = "1990-03-07"
    // 联系人证件有效期结束时间
    extendInfoMap["contact_period_end"] = "长期"
    // 证书类型
    extendInfoMap["cert_type"] = "CERTIFICATE_TYPE_2389"
    // 证书编号
    extendInfoMap["cert_number"] = "1234567892"
    // 证书照片
    extendInfoMap["cert_copy"] = ""
    // 小微经营类型
    extendInfoMap["micro_biz_type"] = ""
    // 门店名称
    extendInfoMap["store_name"] = ""
    // 门店门头照片
    extendInfoMap["store_header_copy"] = ""
    // 店内环境照片
    extendInfoMap["store_indoor_copy"] = ""
    // 门店省市编码
    extendInfoMap["store_address_code"] = ""
    // 门店地址
    extendInfoMap["store_address"] = ""
    // 身份证件正面照片
    extendInfoMap["identification_front_copy"] = "c7faf0e6-39e9-3c35-9075-2312ad6f4ea4"
    // 身份证件反面照片
    extendInfoMap["identification_back_copy"] = "c7faf0e6-39e9-3c35-9075-2312ad6f4ea4"
    // 单位证明函照片
    extendInfoMap["company_prove_copy"] = ""
    // 是否金融机构
    extendInfoMap["finance_institution_flag"] = "N"
    // 金融机构类型
    extendInfoMap["finance_type"] = ""
    // 特殊行业Id
    extendInfoMap["category_id"] = ""
    // 文件列表
    extendInfoMap["special_file_info_list"] = getSpecialFileInfoList()
    return extendInfoMap
}

func getUboInfoList() string {
    dto := make(map[string]interface{})
    // 受益人证件类型
    dto["ubo_id_doc_type"] = "00"
    // 证件正面照片
    dto["ubo_id_doc_copy"] = "c7faf0e6-39e9-3c35-9075-2312ad6f4ea4"
    // 受益人证件姓名
    dto["ubo_id_doc_name"] = "杨雷"
    // 受益人证件号码
    dto["ubo_id_doc_number"] = "110101199003072631"
    // 证件居住地址
    dto["ubo_id_doc_address"] = "上海市徐汇区宜山路789号"
    // 证件有效期开始时间
    dto["ubo_period_begin"] = "19900307"
    // 证件有效期结束时间
    dto["ubo_period_end"] = "长期"
    // 证件反面照片
    dto["ubo_id_doc_copy_back"] = "c7faf0e6-39e9-3c35-9075-2312ad6f4ea4"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getSpecialFileInfoList() string {
    dto := make(map[string]interface{})
    // 文件类型
    dto["file_type"] = "F33"
    // 文件jfileID
    dto["file_id"] = "49ac7d9b-851c-31b4-8b21-2983ea97b98c"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

