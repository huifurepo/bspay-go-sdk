/**
 * 全渠道资金付款到账户退款
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2EfpAcctpaymentRefundRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    OrgHfSeqId string `json:"org_hf_seq_id" structs:"org_hf_seq_id"` // 原交易全局流水号org_hf_seq_id和org_req_seq_id二选一； &lt;font color&#x3D;&quot;green&quot;&gt;示例值：00470topo1A211015160805P090ac132fef00000&lt;/font&gt;
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原交易请求流水号org_hf_seq_id和org_req_seq_id二选一；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：2021091708126665002&lt;/font&gt;
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原交易请求日期
    RefundAmt string `json:"refund_amt" structs:"refund_amt"` // 退款金额
    AcctSplitBunch string `json:"acct_split_bunch" structs:"acct_split_bunch"` // 接收方退款对象

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2EfpAcctpaymentRefundRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2EfpAcctpaymentRefundRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2EfpAcctpaymentRefundRequest(reqParam)
}

func (bp *BsPay) V2EfpAcctpaymentRefundRequest(reqParam V2EfpAcctpaymentRefundRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_EFP_ACCTPAYMENT_REFUND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
