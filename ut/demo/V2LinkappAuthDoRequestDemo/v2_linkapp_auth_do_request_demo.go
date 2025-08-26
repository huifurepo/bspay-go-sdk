/**
 * 商户公域授权 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2LinkappAuthDoRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2LinkappAuthDoRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2LinkappAuthDoRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付商户号
        HuifuId:"6666000107203801",
        // 平台类型
        PlatformType:"21",
        // 协议地址
        ContractUrl:"https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/spin/files/斗拱增值业务服务协议 V1.020231120.docx",
        // 签约商户名称
        ContractMerName:"于云飞",
        // 签约时间
        ContractTime:"1744008692000",
        // 登录用手机号第一次登录有需要手机验证码的情况;（需要授权手机安装一个转发短信的应用）
        // PhoneNumber:"test",
        // 商户类型商户类型：0个人店 1企业 2个体工商户 3其他(目前固定填3即可)
        // MerchantType:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2LinkappAuthDoRequest(dgReq)
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
    // 回调地址
    // extendInfoMap["callback_url"] = ""
    // 
    // extendInfoMap["以下仅RPA授权输入"] = ""
    // 登录账号名称
    // extendInfoMap["account_name"] = ""
    // 登录账号密码
    // extendInfoMap["account_password"] = ""
    // 门店名称
    // extendInfoMap["shop_name"] = ""
    // 管理员手机号
    // extendInfoMap["admin_phone_num"] = ""
    return extendInfoMap
}

