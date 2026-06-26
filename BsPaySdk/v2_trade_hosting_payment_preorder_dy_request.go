/**
 * 抖音直连下单接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeHostingPaymentPreorderDyRequest struct {
    PreOrderType string `json:"pre_order_type" structs:"pre_order_type"` // 预下单类型
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 交易金额
    GoodsDesc string `json:"goods_desc" structs:"goods_desc"` // 商品描述
    DyData string `json:"dy_data" structs:"dy_data"` // 抖音参数集合

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeHostingPaymentPreorderDyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeHostingPaymentPreorderDyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeHostingPaymentPreorderDyRequest(reqParam)
}

func (bp *BsPay) V2TradeHostingPaymentPreorderDyRequest(reqParam V2TradeHostingPaymentPreorderDyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_HOSTING_PAYMENT_PREORDER
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
