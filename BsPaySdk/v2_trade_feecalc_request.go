/**
 * 手续费试算
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeFeecalcRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    TradeType string `json:"trade_type" structs:"trade_type"` // 交易类型
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 交易金额

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeFeecalcRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeFeecalcRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeFeecalcRequest(reqParam)
}

func (bp *BsPay) V2TradeFeecalcRequest(reqParam V2TradeFeecalcRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_FEECALC
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
