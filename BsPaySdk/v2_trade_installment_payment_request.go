/**
 * 分期支付
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeInstallmentPaymentRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 交易金额
    InstallmentNum string `json:"installment_num" structs:"installment_num"` // 分期数
    GoodsDesc string `json:"goods_desc" structs:"goods_desc"` // 商品描述
    RiskCheckData string `json:"risk_check_data" structs:"risk_check_data"` // 安全信息
    JdbtData string `json:"jdbt_data" structs:"jdbt_data"` // 京东白条分期信息trans_type&#x3D;JDBT时，必填jsonObject字符串，京东白条分期相关信息通过该参数集上送
    YljfqData string `json:"yljfq_data" structs:"yljfq_data"` // 银联聚分期信息trans_type&#x3D;YLJFQ-银联聚分期时，必填jsonObject字符串，银联聚分期相关信息通过该参数集上送

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeInstallmentPaymentRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeInstallmentPaymentRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeInstallmentPaymentRequest(reqParam)
}

func (bp *BsPay) V2TradeInstallmentPaymentRequest(reqParam V2TradeInstallmentPaymentRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_INSTALLMENT_PAYMENT
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
