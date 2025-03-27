/**
 * 全渠道资金付款申请
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2EfpSurrogateRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付id
    CashAmt string `json:"cash_amt" structs:"cash_amt"` // 交易金额.单位:元，2位小数
    CardNo string `json:"card_no" structs:"card_no"` // 银行账号
    BankCode string `json:"bank_code" structs:"bank_code"` // 银行编号
    CardName string `json:"card_name" structs:"card_name"` // 银行卡用户名
    CardAcctType string `json:"card_acct_type" structs:"card_acct_type"` // 对公对私标识
    ProvId string `json:"prov_id" structs:"prov_id"` // 省份
    AreaId string `json:"area_id" structs:"area_id"` // 地区
    MobileNo string `json:"mobile_no" structs:"mobile_no"` // 手机号对私必填，使用斗拱系统的公钥对手机号进行RSA加密得到秘文；  示例值：b9LE5RccVVLChrHgo9lvp……PhWhjKrWg2NPfbe0mkUDd
    CertType string `json:"cert_type" structs:"cert_type"` // 证件类型
    CertNo string `json:"cert_no" structs:"cert_no"` // 证件号
    LicenceCode string `json:"licence_code" structs:"licence_code"` // 统一社会信用代码对公必填
    AgreementUrl string `json:"agreement_url" structs:"agreement_url"` // 挂网协议地址

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2EfpSurrogateRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2EfpSurrogateRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2EfpSurrogateRequest(reqParam)
}

func (bp *BsPay) V2EfpSurrogateRequest(reqParam V2EfpSurrogateRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_EFP_SURROGATE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
