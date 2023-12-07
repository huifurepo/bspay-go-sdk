/**
 * 企业用户基本信息开户 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2UserBasicdataEntRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2UserBasicdataEntRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2UserBasicdataEntRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 企业用户名称
        RegName:"企业商户名称8225",
        // 营业执照编号
        LicenseCode:"20220222013747149",
        // 证照有效期类型
        LicenseValidityType:"1",
        // 证照有效期起始日期
        LicenseBeginDate:"20200117",
        // 证照有效期结束日期日期格式：yyyyMMdd; 非长期有效时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20320905&lt;/font&gt;
        LicenseEndDate:"20400117",
        // 注册地址(省)
        RegProvId:"310000",
        // 注册地址(市)
        RegAreaId:"310100",
        // 注册地址(区)
        RegDistrictId:"310104",
        // 注册地址(详细信息)
        RegDetail:"上海市宜山路",
        // 法人姓名
        LegalName:"陈立五",
        // 法人证件类型
        LegalCertType:"00",
        // 法人证件号码
        LegalCertNo:"321084198912066582",
        // 法人证件有效期类型
        LegalCertValidityType:"1",
        // 法人证件有效期开始日期
        LegalCertBeginDate:"20120801",
        // 法人证件有效期截止日期日期格式：yyyyMMdd; 非长期有效时必填，长期有效为空；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20320905&lt;/font&gt;
        LegalCertEndDate:"20300801",
        // 联系人姓名
        ContactName:"小的",
        // 联系人手机号
        ContactMobile:"13764462211",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2UserBasicdataEntRequest(dgReq)
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
    // 经营简称
    extendInfoMap["short_name"] = "企业商户"
    // 联系人电子邮箱
    extendInfoMap["contact_email"] = "jeff.peng@huifu.com"
    // 管理员账号
    extendInfoMap["login_name"] = "Lg2022022201374721361"
    // 操作员
    extendInfoMap["operator_id"] = ""
    // 是否发送短信标识
    extendInfoMap["sms_send_flag"] = "1"
    // 扩展方字段
    extendInfoMap["expand_id"] = ""
    // 文件列表
    // extendInfoMap["file_list"] = getFileList()
    // 公司类型
    // extendInfoMap["ent_type"] = ""
    return extendInfoMap
}

func getFileList() string {
    dto := make(map[string]interface{})
    // 文件类型
    // dto["file_type"] = "test"
    // 文件jfileID
    // dto["file_id"] = "test"
    // 文件名称
    // dto["file_name"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

