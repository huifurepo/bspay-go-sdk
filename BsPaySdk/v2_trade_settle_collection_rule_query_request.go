/**
 * 归集配置查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeSettleCollectionRuleQueryRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    OutHuifuId string `json:"out_huifu_id" structs:"out_huifu_id"` // 转出方商户号转出方商户号和转入方商户号二选一必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123124&lt;/font&gt;
    InHuifuId string `json:"in_huifu_id" structs:"in_huifu_id"` // 转入方商户号转出方商户号和转入方商户号二选一必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123124&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeSettleCollectionRuleQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeSettleCollectionRuleQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeSettleCollectionRuleQueryRequest(reqParam)
}

func (bp *BsPay) V2TradeSettleCollectionRuleQueryRequest(reqParam V2TradeSettleCollectionRuleQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_SETTLE_COLLECTION_RULE_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
