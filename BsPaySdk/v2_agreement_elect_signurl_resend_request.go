/**
 * 电子协议签署链接重发
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2AgreementElectSignurlResendRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    AgreementId string `json:"agreement_id" structs:"agreement_id"` // 协议编号
    SignUserEmail string `json:"sign_user_email" structs:"sign_user_email"` // 个人邮箱

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2AgreementElectSignurlResendRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2AgreementElectSignurlResendRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2AgreementElectSignurlResendRequest(reqParam)
}

func (bp *BsPay) V2AgreementElectSignurlResendRequest(reqParam V2AgreementElectSignurlResendRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_AGREEMENT_ELECT_SIGNURL_RESEND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
