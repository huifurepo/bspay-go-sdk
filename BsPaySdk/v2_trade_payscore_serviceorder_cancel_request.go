/**
 * 取消支付分订单
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePayscoreServiceorderCancelRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    Reason string `json:"reason" structs:"reason"` // 取消服务订单原因

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePayscoreServiceorderCancelRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePayscoreServiceorderCancelRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePayscoreServiceorderCancelRequest(reqParam)
}

func (bp *BsPay) V2TradePayscoreServiceorderCancelRequest(reqParam V2TradePayscoreServiceorderCancelRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYSCORE_SERVICEORDER_CANCEL
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
