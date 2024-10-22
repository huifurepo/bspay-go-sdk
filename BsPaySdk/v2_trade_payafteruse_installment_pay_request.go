/**
 * 分期扣款
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePayafteruseInstallmentPayRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 客户号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 交易金额
    GoodsDesc string `json:"goods_desc" structs:"goods_desc"` // 商品描述
    RiskCheckData string `json:"risk_check_data" structs:"risk_check_data"` // 风控信息
    TimeExpire string `json:"time_expire" structs:"time_expire"` // 交易有效期
    AlipayData string `json:"alipay_data" structs:"alipay_data"` // 支付宝扩展参数集合

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePayafteruseInstallmentPayRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePayafteruseInstallmentPayRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePayafteruseInstallmentPayRequest(reqParam)
}

func (bp *BsPay) V2TradePayafteruseInstallmentPayRequest(reqParam V2TradePayafteruseInstallmentPayRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYAFTERUSE_INSTALLMENT_PAY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
