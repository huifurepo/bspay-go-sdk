/**
 * 微信直连-修改微信结算帐号
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantDirectWechatSettlementinfoModifyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付ID
    AppId string `json:"app_id" structs:"app_id"` // 开发者的应用ID
    MchId string `json:"mch_id" structs:"mch_id"` // 微信商户号
    SubMchid string `json:"sub_mchid" structs:"sub_mchid"` // 特约商户号
    AccountType string `json:"account_type" structs:"account_type"` // 账户类型
    AccountBank string `json:"account_bank" structs:"account_bank"` // 开户银行
    BankAddressCode string `json:"bank_address_code" structs:"bank_address_code"` // 开户银行省市编码
    AccountNumber string `json:"account_number" structs:"account_number"` // 银行账号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantDirectWechatSettlementinfoModifyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantDirectWechatSettlementinfoModifyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantDirectWechatSettlementinfoModifyRequest(reqParam)
}

func (bp *BsPay) V2MerchantDirectWechatSettlementinfoModifyRequest(reqParam V2MerchantDirectWechatSettlementinfoModifyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_DIRECT_WECHAT_SETTLEMENTINFO_MODIFY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
