/**
 * 应用场景
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3TradePaymentJspayRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    GoodsDesc string `json:"goods_desc" structs:"goods_desc"` // 商品描述
    TradeType string `json:"trade_type" structs:"trade_type"` // 交易类型
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 交易金额

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3TradePaymentJspayRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3TradePaymentJspayRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3TradePaymentJspayRequest(reqParam)
}

func (bp *BsPay) V3TradePaymentJspayRequest(reqParam V3TradePaymentJspayRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_TRADE_PAYMENT_JSPAY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
