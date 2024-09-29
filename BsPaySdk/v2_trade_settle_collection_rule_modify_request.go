/**
 * 修改归集配置
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeSettleCollectionRuleModifyRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    OutHuifuId string `json:"out_huifu_id" structs:"out_huifu_id"` // 转出方商户号
    OutAcctId string `json:"out_acct_id" structs:"out_acct_id"` // 转出方账户号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeSettleCollectionRuleModifyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeSettleCollectionRuleModifyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeSettleCollectionRuleModifyRequest(reqParam)
}

func (bp *BsPay) V2TradeSettleCollectionRuleModifyRequest(reqParam V2TradeSettleCollectionRuleModifyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_SETTLE_COLLECTION_RULE_MODIFY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
