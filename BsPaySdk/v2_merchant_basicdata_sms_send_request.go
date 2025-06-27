/**
 * 商户短信发送
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBasicdataSmsSendRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付Id
    Phone string `json:"phone" structs:"phone"` // 手机号verify_type&#x3D;&#39;elecAcctSign&#39;时，手机号为空，系统自动取联系人手机号; &lt;font color&#x3D;&quot;green&quot;&gt;示例值：13911111111&lt;/font&gt;
    VerifyType string `json:"verify_type" structs:"verify_type"` // 验证类型
    OperationType string `json:"operation_type" structs:"operation_type"` // 操作类型verify_type&#x3D;&#39;elecAcctSign&#39;时必填；枚举值：sendSmsCode-发送验证码；identitySmsCode-验证码核实；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：sendSmsCode&lt;/font&gt;
    VerifyCode string `json:"verify_code" structs:"verify_code"` // 验证码verify_type&#x3D;&#39;elecAcctSign&#39;且operation_type&#x3D;&#39;identitySmsCode&#39;时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：123456&lt;/font&gt;
    ElecAcctSignSeqId string `json:"elec_acct_sign_seq_id" structs:"elec_acct_sign_seq_id"` // 中信签约流水号verify_type&#x3D;&#39;elecAcctSign&#39;且operation_type&#x3D;&#39;identitySmsCode&#39;时必填；值为中信E管家签约发送短信时返回值；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：EMSSBPG2504284593690058431260676&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBasicdataSmsSendRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBasicdataSmsSendRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBasicdataSmsSendRequest(reqParam)
}

func (bp *BsPay) V2MerchantBasicdataSmsSendRequest(reqParam V2MerchantBasicdataSmsSendRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BASICDATA_SMS_SEND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
