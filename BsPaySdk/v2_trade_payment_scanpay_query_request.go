/**
 * 交易查询接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePaymentScanpayQueryRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原机构请求日期
    OutOrdId string `json:"out_ord_id" structs:"out_ord_id"` // 汇付服务订单号out_ord_id,org_hf_seq_id,org_req_seq_id 必填其一；汇付生成的服务订单号；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1234323JKHDFE1243252&lt;/font&gt;
    OrgHfSeqId string `json:"org_hf_seq_id" structs:"org_hf_seq_id"` // 创建服务订单返回的汇付全局流水号out_ord_id,org_hf_seq_id,org_req_seq_id 必填其一；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：00290TOP1GR210919004230P853ac13262200000&lt;/font&gt;
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 服务订单创建请求流水号out_ord_id,org_hf_seq_id,org_req_seq_id 必填其一；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：202110210012100005&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePaymentScanpayQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePaymentScanpayQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePaymentScanpayQueryRequest(reqParam)
}

func (bp *BsPay) V2TradePaymentScanpayQueryRequest(reqParam V2TradePaymentScanpayQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYMENT_SCANPAY_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
