/**
 * 全渠道资金付款到账户退款查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2EfpAcctpaymentRefundQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 退款交易请求流水号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 退款交易请求日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2EfpAcctpaymentRefundQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2EfpAcctpaymentRefundQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2EfpAcctpaymentRefundQueryRequest(reqParam)
}

func (bp *BsPay) V2EfpAcctpaymentRefundQueryRequest(reqParam V2EfpAcctpaymentRefundQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_EFP_ACCTPAYMENT_REFUND_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
