/**
 * 查询支付分订单
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePayscoreServiceorderQueryRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePayscoreServiceorderQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePayscoreServiceorderQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePayscoreServiceorderQueryRequest(reqParam)
}

func (bp *BsPay) V2TradePayscoreServiceorderQueryRequest(reqParam V2TradePayscoreServiceorderQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYSCORE_SERVICEORDER_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
