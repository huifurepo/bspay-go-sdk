/**
 * 个人用户基本信息开户 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2UserBasicdataIndvRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2UserBasicdataIndvRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2UserBasicdataIndvRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 个人姓名
        Name:"个人用户名称3657",
        // 个人证件类型
        CertType:"00",
        // 个人证件号码
        CertNo:"321084198912066512",
        // 个人证件有效期类型
        CertValidityType:"1",
        // 个人证件有效期开始日期
        CertBeginDate:"20200117",
        // 手机号
        MobileNo:"13764462205",
        // 地址开通中信E管家必填
        // Address:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2UserBasicdataIndvRequest(dgReq)
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
    // 个人证件有效期截止日期
    extendInfoMap["cert_end_date"] = "20400117"
    // 电子邮箱
    extendInfoMap["email"] = "jeff.peng@huifu.com"
    // 管理员账号
    extendInfoMap["login_name"] = "Lg2022022201394910571"
    // 是否发送短信标识
    extendInfoMap["sms_send_flag"] = "1"
    // 拓展方字段
    extendInfoMap["expand_id"] = ""
    // 文件列表
    extendInfoMap["file_list"] = getFileList()
    return extendInfoMap
}

func getFileList() string {
    dto := make(map[string]interface{})
    // 文件类型
    dto["file_type"] = "F04"
    // 文件jfileID
    dto["file_id"] = "2022022201394949297117211"
    // 文件名称
    dto["file_name"] = "企业营业执照1.jpg"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

