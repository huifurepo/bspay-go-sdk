/**
 * 交易确认退款查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePaymentDelaytransConfirmrefundqueryRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原交易请求日期
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原交易请求流水号指交易确认退款请求流水号，org_req_seq_id和org_hf_seq_id二选一；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：2021091708126665002&lt;/font&gt;
    OrgHfSeqId string `json:"org_hf_seq_id" structs:"org_hf_seq_id"` // 原退款全局流水号原交易确认退款全局流水号。org_req_seq_id和org_hf_seq_id二选一；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：003500TOP2B211021163242P447ac132fd200000&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePaymentDelaytransConfirmrefundqueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePaymentDelaytransConfirmrefundqueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePaymentDelaytransConfirmrefundqueryRequest(reqParam)
}

func (bp *BsPay) V2TradePaymentDelaytransConfirmrefundqueryRequest(reqParam V2TradePaymentDelaytransConfirmrefundqueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYMENT_DELAYTRANS_CONFIRMREFUNDQUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
