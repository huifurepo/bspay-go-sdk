/**
 * 托管交易查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeHostingPaymentQueryorderinfoRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原交易请求日期
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原交易请求流水号与**party_order_id**二选一，必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：rQ2021121311173944&lt;/font&gt;
    PartyOrderId string `json:"party_order_id" structs:"party_order_id"` // 用户账单上的商户订单号与**org_req_seq_id**二选一，必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：03232109190255105603561&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeHostingPaymentQueryorderinfoRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeHostingPaymentQueryorderinfoRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeHostingPaymentQueryorderinfoRequest(reqParam)
}

func (bp *BsPay) V2TradeHostingPaymentQueryorderinfoRequest(reqParam V2TradeHostingPaymentQueryorderinfoRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_HOSTING_PAYMENT_QUERYORDERINFO
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
