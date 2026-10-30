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
        // 纳税人识别号支付宝和微信乐企模式非必填，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：91510108MA7CQPFFO6&lt;/font&gt;
        TaxPayerId:"91310000795605184T",
        // 销方名称支付宝和微信乐企模式非必填，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：汇付数电测试&lt;/font&gt;
        TaxPayerName:"汇付网络技术（上海）有限公司",
        // 公司电话支付宝和微信乐企模式非必填，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：021-33323999&lt;/font&gt;
        TelNo:"021-33323999",
        // 公司地址支付宝和微信乐企模式非必填，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：徐汇宜山路700号C5栋&lt;/font&gt;
        RegAddress:"徐汇宜山路700号C5栋021-33323999",
        // 开户银行支付宝和微信乐企模式非必填，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：民生银行徐汇支行&lt;/font&gt;
        BankName:"民生银行徐汇支行",
        // 开户账户支付宝和微信乐企模式非必填，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0206014180001609&lt;/font&gt;
        AccountNo:"0206014180001609",
        // 联系人手机号支付宝和微信乐企模式非必填，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：18011212123&lt;/font&gt;
        ContactPhoneNo:"17621100776",
        // 开票方式
        OpenMode:"2",
        // 所属省支付宝和微信乐企模式非必填，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：310000&lt;/font&gt;，参见[地区编码](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_dqbm)
        ProvId:"310000",
        // 所属市支付宝和微信乐企模式非必填，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：310100&lt;/font&gt;，参见[地区编码](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_dqbm)
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
    // 开票模式
    extendInfoMap["issue_mode"] = "LE_QI_LINK"
    // 区域代码
    // extendInfoMap["area_code"] = ""
    // 协议附件
    // extendInfoMap["net_url"] = ""
    // 签约商户名称
    // extendInfoMap["mer_name"] = ""
    // 生效日期
    // extendInfoMap["start_date"] = ""
    return extendInfoMap
}

