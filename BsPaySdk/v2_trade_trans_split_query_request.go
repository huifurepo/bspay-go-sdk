/**
 * 交易分账明细查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeTransSplitQueryRequest struct {
    HfSeqId string `json:"hf_seq_id" structs:"hf_seq_id"` // 分账交易汇付全局流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    OrdType string `json:"ord_type" structs:"ord_type"` // 交易类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeTransSplitQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeTransSplitQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeTransSplitQueryRequest(reqParam)
}

func (bp *BsPay) V2TradeTransSplitQueryRequest(reqParam V2TradeTransSplitQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_TRANS_SPLIT_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
