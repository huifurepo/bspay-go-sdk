/**
 * 灵工支付查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2FlexibleTradeQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原请求流水号原请求流水号与原请求全局流水号二选一必填，示例值：2021091708126665001
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原请求日期原请求流水号必填则原请求日期必填，格式：yyyyMMdd；示例值：20210917
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    OrgHfSeqId string `json:"org_hf_seq_id" structs:"org_hf_seq_id"` // 原交易全局流水号原请求流水号与原请求全局流水号二选一必填，  &lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值： 003100TOP1A230816150903P990ac139c0600000&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2FlexibleTradeQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2FlexibleTradeQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2FlexibleTradeQueryRequest(reqParam)
}

func (bp *BsPay) V2FlexibleTradeQueryRequest(reqParam V2FlexibleTradeQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_FLEXIBLE_TRADE_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
