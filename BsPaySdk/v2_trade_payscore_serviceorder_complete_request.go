/**
 * 完结支付分订单
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePayscoreServiceorderCompleteRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    OutOrderNo string `json:"out_order_no" structs:"out_order_no"` // 汇付订单号
    OrdAmt string `json:"ord_amt" structs:"ord_amt"` // 完结金额
    TimeRange string `json:"time_range" structs:"time_range"` // 服务时间

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePayscoreServiceorderCompleteRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePayscoreServiceorderCompleteRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePayscoreServiceorderCompleteRequest(reqParam)
}

func (bp *BsPay) V2TradePayscoreServiceorderCompleteRequest(reqParam V2TradePayscoreServiceorderCompleteRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYSCORE_SERVICEORDER_COMPLETE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
