/**
 * 一键绑卡 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2QuickbuckleOnekeyCardbindRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2QuickbuckleOnekeyCardbindRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2QuickbuckleOnekeyCardbindRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 汇付Id
        HuifuId:"6666000003078984",
        // 顾客用户号 
        OutCustId:"asdtrw123123123",
        // 银行号
        BankId:"01030000",
        // 银行卡开户姓名 
        CardName:"YTRf65hBDkH9UU1AwG16r4Nlc/X1rH6ejKbvmqT80exJ6whdHI1zB+izBtNBOJfhRNbIOhi1FrRuE5b7wnt/03Q+vwWQQLDGJXWZf92yp+eIRDHg8JdbjOgxKvF2q4Qw5704jbsjQm4UJW5fqRhzRPtnnAL9zzTSgVhuQ0KCwc8=",
        // 银行卡绑定身份证
        CertId:"gk8zqa+zvIUAvzV3Bjzzw+vRgq2LgTzQI8PRoqUdbeuFMbFZ6LllQpcOhWIz6F82VtFBKzLd5kPOaCwlQCiwsXhSqUFB11zgzKUtVIuiv9lHY/EsjRwqDBhgeR5f2H9FXr3wQ9f7bI7t8ca9o93QxrXr/1MDBq7fGok0xu2ytsM=",
        // 银行卡绑定证件类型
        CertType:"00",
        // 证件有效期到期日长期有效不填.格式：yyyyMMdd。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20311111&lt;/font&gt;
        CertEndDate:"20370121",
        // 银行卡绑定手机号 
        CardMp:"e2zKkJ6PE6UtZhgz42pqgPLQh6p83/WJsG7EVSgYYgN+7MIiCzXnjTpmpv0Cgv7cYTbQTBf/NF5jqeI8BpFjP7C6gg+0cjqW2tgdVxyfqZLu2fEJRth7NgZKgS2ZKbkZ8PNfUk7V+aHbqkAVKY92bdcSQSNIuWaJCeIF34w+l+k=",
        // 卡的借贷类型
        DcType:"D",
        // 异步请求地址
        AsyncReturnUrl:"http://192.168.85.157:30031/sspm/testVirgo",
        // 设备信息域
        TrxDeviceInf:getTrxDeviceInf(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2QuickbuckleOnekeyCardbindRequest(dgReq)
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
    // 证件有效期类型
    extendInfoMap["cert_validity_type"] = "0"
    // 证件有效期起始日
    extendInfoMap["cert_begin_date"] = "20170121"
    // 挂网协议编号
    extendInfoMap["protocol_no"] = ""
    // 风控信息
    // extendInfoMap["risk_info"] = getRiskInfo()
    // 回调页面地址
    // extendInfoMap["verify_front_url"] = ""
    return extendInfoMap
}

func getRiskInfo() string {
    dto := make(map[string]interface{})
    // IP类型
    // dto["ip_type"] = "test"
    // IP值
    // dto["source_ip"] = "test"
    // 设备标识
    // dto["device_id"] = ""
    // 设备类型
    // dto["device_type"] = ""
    // 应用提供方账户ID
    // dto["account_id_hash"] = ""
    // 银行预留手机号
    // dto["mobile"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getTrxDeviceInf() string {
    dto := make(map[string]interface{})
    // 银行预留手机号
    dto["trx_mobile_num"] = "15556622368"
    // 设备类型
    dto["trx_device_type"] = "1"
    // 交易设备IP
    dto["trx_device_ip"] = "10.10.0.1"
    // 交易设备MAC
    dto["trx_device_mac"] = "10.10.0.1"
    // 交易设备IMEI
    dto["trx_device_imei"] = "030147441006000182623"
    // 交易设备IMSI
    dto["trx_device_imsi"] = "030147441006000182623"
    // 交易设备ICCID
    dto["trx_device_icc_id"] = "030147441006000182623"
    // 交易设备WIFIMAC
    dto["trx_device_wfifi_mac"] = "030147441006000182623"
    // 交易设备GPS
    dto["trx_device_gps"] = "030147441006000182623"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

