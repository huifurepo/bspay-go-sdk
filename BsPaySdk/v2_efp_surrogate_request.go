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
    CardNo string `json:"card_no" structs:"card_no"` // 银行账号使用斗拱系统的公钥对银行账号进行RSA加密得到秘文；  示例值：b9LE5RccVVLChrHgo9lvp……PhWhjKrWg2NPfbe0mkQ&#x3D;&#x3D; 到账类型标识为E或P时必填
    BankCode string `json:"bank_code" structs:"bank_code"` // 银行编号银行编号 到账类型标识为E或P时必填
    CardName string `json:"card_name" structs:"card_name"` // 银行卡用户名银行卡用户名 到账类型标识为E或P时必填
    CardAcctType string `json:"card_acct_type" structs:"card_acct_type"` // 到账类型标识
    ProvId string `json:"prov_id" structs:"prov_id"` // 省份到账类型标识为E或P时必填
    AreaId string `json:"area_id" structs:"area_id"` // 地区到账类型标识为E或P时必填
    MobileNo string `json:"mobile_no" structs:"mobile_no"` // 手机号对私必填，使用斗拱系统的公钥对手机号进行RSA加密得到秘文；  示例值：b9LE5RccVVLChrHgo9lvp……PhWhjKrWg2NPfbe0mkUDd
    CertType string `json:"cert_type" structs:"cert_type"` // 证件类型证件类型01：身份证  03：护照  06：港澳通行证  07：台湾通行证  09：外国人居留证  11：营业执照  12：组织机构代码证  14：统一社会信用代码  99：其他  示例值：14 到账类型标识为E或P时必填
    CertNo string `json:"cert_no" structs:"cert_no"` // 证件号使用斗拱系统的公钥对证件号进行RSA加密得到秘文；  示例值：b9LE5RccVVLChrHgo9lvp……PhWhjKrWg2NPfbe0mkQ 到账类型标识为P时必填
    LicenceCode string `json:"licence_code" structs:"licence_code"` // 统一社会信用代码到账类型标识为E时必填
    AcctSplitBunch string `json:"acct_split_bunch" structs:"acct_split_bunch"` // 入账接收方对象json格式,到账类型标识为H时必填

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
