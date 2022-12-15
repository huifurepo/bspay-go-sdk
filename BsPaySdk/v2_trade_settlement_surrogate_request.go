/**
 * 银行卡代发
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeSettlementSurrogateRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    CashAmt string `json:"cash_amt" structs:"cash_amt"` // 代发金额
    PurposeDesc string `json:"purpose_desc" structs:"purpose_desc"` // 代发用途描述
    Province string `json:"province" structs:"province"` // 省份
    Area string `json:"area" structs:"area"` // 地区
    BankCode string `json:"bank_code" structs:"bank_code"` // 银行编号
    BankAccountName string `json:"bank_account_name" structs:"bank_account_name"` // 银行卡用户名
    CardAcctType string `json:"card_acct_type" structs:"card_acct_type"` // 对公对私标识
    BankCardNoCrypt string `json:"bank_card_no_crypt" structs:"bank_card_no_crypt"` // 银行账号密文
    CertificateNoCrypt string `json:"certificate_no_crypt" structs:"certificate_no_crypt"` // 证件号密文
    CertificateType string `json:"certificate_type" structs:"certificate_type"` // 证件类型
    MobileNoCrypt string `json:"mobile_no_crypt" structs:"mobile_no_crypt"` // 手机号密文
    IntoAcctDateType string `json:"into_acct_date_type" structs:"into_acct_date_type"` // 到账日期类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeSettlementSurrogateRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeSettlementSurrogateRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeSettlementSurrogateRequest(reqParam)
}

func (bp *BsPay) V2TradeSettlementSurrogateRequest(reqParam V2TradeSettlementSurrogateRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_SETTLEMENT_SURROGATE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
