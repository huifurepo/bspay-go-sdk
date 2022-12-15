/**
 * 支付宝申诉请求凭证
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantComplaintRequestCertificatesRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求汇付流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求汇付时间
    RiskBizId string `json:"risk_biz_id" structs:"risk_biz_id"` // 支付宝推送流水号
    MerchantType string `json:"merchant_type" structs:"merchant_type"` // 商户类型
    OperationType string `json:"operation_type" structs:"operation_type"` // 商户经营模式
    PaymentScene string `json:"payment_scene" structs:"payment_scene"` // 收款应用场景

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantComplaintRequestCertificatesRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantComplaintRequestCertificatesRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantComplaintRequestCertificatesRequest(reqParam)
}

func (bp *BsPay) V2MerchantComplaintRequestCertificatesRequest(reqParam V2MerchantComplaintRequestCertificatesRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_COMPLAINT_REQUEST_CERTIFICATES
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
