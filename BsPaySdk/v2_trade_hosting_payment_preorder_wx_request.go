/**
 * 微信小程序预下单接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeHostingPaymentPreorderWxRequest struct {
    PreOrderType string `json:"pre_order_type" structs:"pre_order_type"` // 预下单类型
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 交易金额
    GoodsDesc string `json:"goods_desc" structs:"goods_desc"` // 商品描述
    MiniappData string `json:"miniapp_data" structs:"miniapp_data"` // 微信小程序扩展参数集合

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeHostingPaymentPreorderWxRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeHostingPaymentPreorderWxRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeHostingPaymentPreorderWxRequest(reqParam)
}

func (bp *BsPay) V2TradeHostingPaymentPreorderWxRequest(reqParam V2TradeHostingPaymentPreorderWxRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_HOSTING_PAYMENT_PREORDER
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
