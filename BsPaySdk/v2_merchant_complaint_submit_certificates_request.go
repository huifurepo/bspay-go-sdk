/**
 * 支付宝申诉提交凭证
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantComplaintSubmitCertificatesRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求汇付流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求汇付时间
    RiskBizId string `json:"risk_biz_id" structs:"risk_biz_id"` // 支付宝推送流水号
    RelievingId string `json:"relieving_id" structs:"relieving_id"` // 申诉解限的唯一ID
    RelieveRiskType string `json:"relieve_risk_type" structs:"relieve_risk_type"` // 解限风险类型
    RelieveCertDataList string `json:"relieve_cert_data_list" structs:"relieve_cert_data_list"` // 提交的凭证数据

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantComplaintSubmitCertificatesRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantComplaintSubmitCertificatesRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantComplaintSubmitCertificatesRequest(reqParam)
}

func (bp *BsPay) V2MerchantComplaintSubmitCertificatesRequest(reqParam V2MerchantComplaintSubmitCertificatesRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_COMPLAINT_SUBMIT_CERTIFICATES
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
