/**
 * 电子协议拒签
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2AgreementElectRejectsignRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    AgreementId string `json:"agreement_id" structs:"agreement_id"` // 协议编号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2AgreementElectRejectsignRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2AgreementElectRejectsignRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2AgreementElectRejectsignRequest(reqParam)
}

func (bp *BsPay) V2AgreementElectRejectsignRequest(reqParam V2AgreementElectRejectsignRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_AGREEMENT_ELECT_REJECTSIGN
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
