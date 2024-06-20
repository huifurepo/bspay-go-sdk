/**
 * 商户注册信息修改 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2InvoiceMerModifyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2InvoiceMerModifyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2InvoiceMerModifyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 开票方汇付ID
        HuifuId:"6666000149801800",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2InvoiceMerModifyRequest(dgReq)
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
    // 销方名称
    extendInfoMap["tax_payer_name"] = "汇付网络技术（上海）有限公司"
    // 公司电话
    extendInfoMap["tel_no"] = "021-33323999"
    // 公司地址
    extendInfoMap["reg_address"] = "徐汇宜山路700号C5栋021-33323999"
    // 开户银行
    extendInfoMap["bank_name"] = "民生银行徐汇支行"
    // 开户账户
    extendInfoMap["account_no"] = "0206014180001609"
    // 联系人手机号
    extendInfoMap["contact_phone_no"] = "17621100776"
    // 联系人
    extendInfoMap["contact"] = "王姗"
    // 联系人身份证号
    extendInfoMap["id_card_no"] = "210123198702122747"
    // 所属省
    extendInfoMap["prov_id"] = "310000"
    // 所属市
    extendInfoMap["city_id"] = "310100"
    return extendInfoMap
}

