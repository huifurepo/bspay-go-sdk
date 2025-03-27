/**
 * 个人用户基本信息修改 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2UserBasicdataIndvModifyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2UserBasicdataIndvModifyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2UserBasicdataIndvModifyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付客户Id
        HuifuId:"6666000103854106",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2UserBasicdataIndvModifyRequest(dgReq)
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
    // 个人证件有效期类型
    extendInfoMap["cert_validity_type"] = "2"
    // 个人证件有效期开始日期
    extendInfoMap["cert_begin_date"] = "20200111"
    // 个人证件有效期截止日期
    extendInfoMap["cert_end_date"] = "20400111"
    // 电子邮箱
    extendInfoMap["email"] = "jeff.peng@huifu.com"
    // 手机号
    extendInfoMap["mobile_no"] = "15556622000"
    // 文件列表
    // extendInfoMap["file_list"] = get8ace1c78E20e4c36B860Bf175485748a()
    // 地址
    // extendInfoMap["address"] = ""
    return extendInfoMap
}

func get8ace1c78E20e4c36B860Bf175485748a() string {
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

