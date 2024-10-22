/**
 * 分期交易退款查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePayafteruseInstallmentRefundqueryRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePayafteruseInstallmentRefundqueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePayafteruseInstallmentRefundqueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePayafteruseInstallmentRefundqueryRequest(reqParam)
}

func (bp *BsPay) V2TradePayafteruseInstallmentRefundqueryRequest(reqParam V2TradePayafteruseInstallmentRefundqueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYAFTERUSE_INSTALLMENT_REFUNDQUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}