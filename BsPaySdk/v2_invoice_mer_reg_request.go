/**
 * 商户注册
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2InvoiceMerRegRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 开票方汇付ID
    TaxPayerId string `json:"tax_payer_id" structs:"tax_payer_id"` // 纳税人识别号
    TaxPayerName string `json:"tax_payer_name" structs:"tax_payer_name"` // 销方名称
    TelNo string `json:"tel_no" structs:"tel_no"` // 公司电话
    RegAddress string `json:"reg_address" structs:"reg_address"` // 公司地址
    BankName string `json:"bank_name" structs:"bank_name"` // 开户银行
    AccountNo string `json:"account_no" structs:"account_no"` // 开户账户
    ContactPhoneNo string `json:"contact_phone_no" structs:"contact_phone_no"` // 联系人手机号
    OpenMode string `json:"open_mode" structs:"open_mode"` // 开票方式
    ProvId string `json:"prov_id" structs:"prov_id"` // 所属省
    CityId string `json:"city_id" structs:"city_id"` // 所属市

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2InvoiceMerRegRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2InvoiceMerRegRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2InvoiceMerRegRequest(reqParam)
}

func (bp *BsPay) V2InvoiceMerRegRequest(reqParam V2InvoiceMerRegRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_INVOICE_MER_REG
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
