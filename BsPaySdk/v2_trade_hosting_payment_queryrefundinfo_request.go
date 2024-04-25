/**
 * 托管交易退款查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeHostingPaymentQueryrefundinfoRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 退款请求日期
    OrgHfSeqId string `json:"org_hf_seq_id" structs:"org_hf_seq_id"` // 退款全局流水号退款请求流水号/退款全局流水号二选一不能都为空；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0030default220825182711P099ac1f343f00000&lt;/font&gt;
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 退款请求流水号退款请求流水号/退款全局流水号二选一不能都为空；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：202110210012100005&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeHostingPaymentQueryrefundinfoRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeHostingPaymentQueryrefundinfoRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeHostingPaymentQueryrefundinfoRequest(reqParam)
}

func (bp *BsPay) V2TradeHostingPaymentQueryrefundinfoRequest(reqParam V2TradeHostingPaymentQueryrefundinfoRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_HOSTING_PAYMENT_QUERYREFUNDINFO
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
