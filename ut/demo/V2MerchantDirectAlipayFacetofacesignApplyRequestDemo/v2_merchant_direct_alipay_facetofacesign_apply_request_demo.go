/**
 * 支付宝直连-申请当面付代签约 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantDirectAlipayFacetofacesignApplyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantDirectAlipayFacetofacesignApplyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantDirectAlipayFacetofacesignApplyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付客户Id
        HuifuId:"6666000003080750",
        // 上级主体ID
        UpperHuifuId:"6666000003078903",
        // 支付宝经营类目
        DirectCategory:"A_A01_4119",
        // 开发者的应用ID
        AppId:"AE150003019",
        // 联系人姓名
        ContactName:"hqqTEST",
        // 联系人手机号
        ContactMobileNo:"15800718101",
        // 联系人电子邮箱
        ContactEmail:"24324@qq.com",
        // 商户账号
        Account:"288000000345345",
        // 服务费率（%）0.38~3之间，精确到0.01。当签约且授权sign_and_auth&#x3D;Y时，必填。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.38&lt;/font&gt;
        Rate:"0.38",
        // 文件列表
        FileList:getFileList(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantDirectAlipayFacetofacesignApplyRequest(dgReq)
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
    // 订单授权凭证
    extendInfoMap["order_ticket"] = "werwe234234234"
    // 签约且授权标识
    extendInfoMap["sign_and_auth"] = "Y"
    // 应用授权令牌
    extendInfoMap["app_auth_token"] = "test0004"
    // 营业执照编号
    extendInfoMap["license_code"] = ""
    // 营业执照有效期类型
    extendInfoMap["license_validity_type"] = "0"
    // 营业执照有效期开始日期
    extendInfoMap["license_begin_date"] = "20200429"
    // 营业执照有效期截止日期
    extendInfoMap["license_end_date"] = "29200429"
    return extendInfoMap
}

func getFileList() string {
    dto := make(map[string]interface{})
    // 文件类型
    dto["file_type"] = "F50"
    // 文件jfileID
    dto["file_id"] = "b53e18b3-f933-357f-9a6f-952c6a021ba5"
    // 文件名称
    dto["file_name"] = "360huxi.jpg"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

