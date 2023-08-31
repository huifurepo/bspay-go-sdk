/**
 * 查询扣款信息
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePayscoreDeductQueryRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePayscoreDeductQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePayscoreDeductQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePayscoreDeductQueryRequest(reqParam)
}

func (bp *BsPay) V2TradePayscoreDeductQueryRequest(reqParam V2TradePayscoreDeductQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYSCORE_DEDUCT_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
