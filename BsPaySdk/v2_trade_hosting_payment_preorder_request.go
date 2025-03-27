/**
 * 支付宝小程序预下单接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeHostingPaymentPreorderRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    PreOrderType string `json:"pre_order_type" structs:"pre_order_type"` // 预下单类型
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 交易金额
    GoodsDesc string `json:"goods_desc" structs:"goods_desc"` // 商品描述
    AppData string `json:"app_data" structs:"app_data"` // app扩展参数集合

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeHostingPaymentPreorderRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeHostingPaymentPreorderRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeHostingPaymentPreorderRequest(reqParam)
}

func (bp *BsPay) V2TradeHostingPaymentPreorderRequest(reqParam V2TradeHostingPaymentPreorderRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_HOSTING_PAYMENT_PREORDER
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
