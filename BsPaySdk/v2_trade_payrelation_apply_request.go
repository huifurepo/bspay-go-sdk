/**
 * 付款关系提交
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePayrelationApplyRequest struct {
    OutHuifuId string `json:"out_huifu_id" structs:"out_huifu_id"` // 出款方商户号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    PayRelations string `json:"pay_relations" structs:"pay_relations"` // 付款关系明细

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePayrelationApplyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePayrelationApplyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePayrelationApplyRequest(reqParam)
}

func (bp *BsPay) V2TradePayrelationApplyRequest(reqParam V2TradePayrelationApplyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYRELATION_APPLY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
