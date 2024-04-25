/**
 * 余额支付退款
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeAcctpaymentRefundRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原余额支付请求日期
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原余额支付请求流水号org_hf_seq_id和orgReqSeqId二选一；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：2021091708126665001&lt;/font&gt;
    OrgHfSeqId string `json:"org_hf_seq_id" structs:"org_hf_seq_id"` // 原余额支付支付全局流水号org_hf_seq_id和orgReqSeqId二选一；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：00470topo1A211015160805P090ac132fef00000&lt;/font&gt;
    OrdAmt string `json:"ord_amt" structs:"ord_amt"` // 退款金额

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeAcctpaymentRefundRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeAcctpaymentRefundRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeAcctpaymentRefundRequest(reqParam)
}

func (bp *BsPay) V2TradeAcctpaymentRefundRequest(reqParam V2TradeAcctpaymentRefundRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ACCTPAYMENT_REFUND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
