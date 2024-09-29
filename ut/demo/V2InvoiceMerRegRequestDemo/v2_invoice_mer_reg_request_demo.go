/**
 * 商户注册 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2InvoiceMerRegRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2InvoiceMerRegRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2InvoiceMerRegRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 开票方汇付ID
        HuifuId:"6666000149801800",
        // 纳税人识别号
        TaxPayerId:"91310000795605184T",
        // 销方名称
        TaxPayerName:"汇付网络技术（上海）有限公司",
        // 公司电话
        TelNo:"021-33323999",
        // 公司地址
        RegAddress:"徐汇宜山路700号C5栋021-33323999",
        // 开户银行
        BankName:"民生银行徐汇支行",
        // 开户账户
        AccountNo:"0206014180001609",
        // 联系人手机号
        ContactPhoneNo:"17621100776",
        // 开票方式
        OpenMode:"2",
        // 所属省
        ProvId:"310000",
        // 所属市
        CityId:"310100",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2InvoiceMerRegRequest(dgReq)
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
    // 联系人
    extendInfoMap["contact"] = "王姗"
    // 联系人身份证号
    extendInfoMap["id_card_no"] = "210123198702122747"
    // 业务到期年限
    extendInfoMap["valid_period"] = "1"
    // 自动续约
    extendInfoMap["auto_renewal"] = "Y"
    // 开票结果异步通知地址
    extendInfoMap["callback_url"] = "http: //service.example.com/to/path"
    return extendInfoMap
}

