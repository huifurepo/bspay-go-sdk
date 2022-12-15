/**
 * 支付宝直连-换取应用授权令牌
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantDirectAlipayAppauthtokenExchangeRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付ID
    AppId string `json:"app_id" structs:"app_id"` // 开发者的应用ID
    OperType string `json:"oper_type" structs:"oper_type"` // 操作类型
    AppAuthCode string `json:"app_auth_code" structs:"app_auth_code"` // 授权码授权码，操作类型为0-换取令牌时必填，其它选填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：123&lt;/font&gt;
    AppAuthToken string `json:"app_auth_token" structs:"app_auth_token"` // 应用授权令牌应用授权令牌，操作类型为1-刷新令牌时，且该字段有值，将与数据库值进行校验；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：202208200312104378&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantDirectAlipayAppauthtokenExchangeRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantDirectAlipayAppauthtokenExchangeRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantDirectAlipayAppauthtokenExchangeRequest(reqParam)
}

func (bp *BsPay) V2MerchantDirectAlipayAppauthtokenExchangeRequest(reqParam V2MerchantDirectAlipayAppauthtokenExchangeRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_DIRECT_ALIPAY_APPAUTHTOKEN_EXCHANGE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
