/**
 * 余额支付
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeAcctpaymentPayRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    OutHuifuId string `json:"out_huifu_id" structs:"out_huifu_id"` // 出款方商户号
    OrdAmt string `json:"ord_amt" structs:"ord_amt"` // 支付金额
    AcctSplitBunch string `json:"acct_split_bunch" structs:"acct_split_bunch"` // 分账对象
    RiskCheckData string `json:"risk_check_data" structs:"risk_check_data"` // 安全信息
    FundType string `json:"fund_type" structs:"fund_type"` // 资金类型资金类型。支付渠道为中信E管家时，资金类型必填（[详见说明](https://paas.huifu.com/open/doc/api/#/yuer/api_zxegjzllx)）
    TransFeeTakeFlag string `json:"trans_fee_take_flag" structs:"trans_fee_take_flag"` // 手续费承担方标识余额支付手续费承担方标识；商户余额支付扣收规则为接口指定承担方时必填！枚举值：&lt;br/&gt;OUT：出款方；&lt;br/&gt;IN：分账接受方。&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：IN&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeAcctpaymentPayRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeAcctpaymentPayRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeAcctpaymentPayRequest(reqParam)
}

func (bp *BsPay) V2TradeAcctpaymentPayRequest(reqParam V2TradeAcctpaymentPayRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ACCTPAYMENT_PAY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
