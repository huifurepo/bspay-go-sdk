/**
 * 交易结算对账文件配置 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantBusiBillConfigRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantBusiBillConfigRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantBusiBillConfigRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 汇付机构编号
        HuifuId:"6666000109133323",
        // 对账文件生成开关
        ReconSendFlag:"Y",
        // 对账单类型
        FileType:"1,2,3,4",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantBusiBillConfigRequest(dgReq)
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
    // SFTP地址
    // extendInfoMap["ftp_addr"] = ""
    // SFTP用户名
    // extendInfoMap["ftp_user"] = ""
    // SFTP密码
    // extendInfoMap["ftp_pwd"] = ""
    // 是否跟随结算时间
    // extendInfoMap["as_settle_flag"] = ""
    // 跑批完成通知地址
    // extendInfoMap["notify_url"] = ""
    // 用户数据范围
    // extendInfoMap["contain_user"] = ""
    // 商户数据范围
    // extendInfoMap["contain_merchant"] = ""
    return extendInfoMap
}

