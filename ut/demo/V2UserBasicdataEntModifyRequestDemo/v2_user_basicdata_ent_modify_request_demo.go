/**
 * 企业用户基本信息修改 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2UserBasicdataEntModifyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2UserBasicdataEntModifyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2UserBasicdataEntModifyRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 汇付客户Id
        HuifuId:"6666000103862211",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2UserBasicdataEntModifyRequest(dgReq)
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
    // 企业用户名称
    // extendInfoMap["reg_name"] = ""
    // 经营简称
    extendInfoMap["short_name"] = "企业商户测试22"
    // 公司类型
    // extendInfoMap["ent_type"] = ""
    // 法人姓名
    extendInfoMap["legal_name"] = "陈立一"
    // 法人证件类型
    extendInfoMap["legal_cert_type"] = "00"
    // 法人证件号码
    extendInfoMap["legal_cert_no"] = "370684198903061000"
    // 法人证件有效期类型
    extendInfoMap["legal_cert_validity_type"] = "0"
    // 法人证件有效期开始日期
    extendInfoMap["legal_cert_begin_date"] = "20121010"
    // 法人证件有效期截止日期
    extendInfoMap["legal_cert_end_date"] = "20301010"
    // 联系人姓名
    extendInfoMap["contact_name"] = "花朵"
    // 联系人电子邮箱
    extendInfoMap["contact_email"] = "chang@huifu.com"
    // 联系人手机号
    extendInfoMap["contact_mobile"] = "13764462000"
    // 证照有效期类型
    extendInfoMap["license_validity_type"] = "1"
    // 证照有效期起始日期
    extendInfoMap["license_begin_date"] = "20200117"
    // 证照有效期结束日期
    extendInfoMap["license_end_date"] = "20400117"
    // 注册地址(省)
    extendInfoMap["reg_prov_id"] = "310000"
    // 注册地址(市)
    extendInfoMap["reg_area_id"] = "310100"
    // 注册地址(区)
    extendInfoMap["reg_district_id"] = "310104"
    // 注册地址(详细信息)
    extendInfoMap["reg_detail"] = "上海市宜山路"
    // 文件列表
    extendInfoMap["file_list"] = getFileList()
    return extendInfoMap
}

func getFileList() string {
    dto := make(map[string]interface{})
    // 文件类型
    dto["file_type"] = "F01"
    // 文件jfileID
    dto["file_id"] = "71da066c-5d15-3658-a86d-4e85ee67808a"
    // 文件名称
    dto["file_name"] = "企业营业执照1.jpg"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

