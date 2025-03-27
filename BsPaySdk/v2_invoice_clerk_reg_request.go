/**
 * 开票员登记
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2InvoiceClerkRegRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付Id
    ClerkIdentity string `json:"clerk_identity" structs:"clerk_identity"` // 开票员登录身份
    LoginAccount string `json:"login_account" structs:"login_account"` // 登录账号
    LoginPassword string `json:"login_password" structs:"login_password"` // 登录密码
    ClerkPhoneNo string `json:"clerk_phone_no" structs:"clerk_phone_no"` // 开票员手机号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2InvoiceClerkRegRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2InvoiceClerkRegRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2InvoiceClerkRegRequest(reqParam)
}

func (bp *BsPay) V2InvoiceClerkRegRequest(reqParam V2InvoiceClerkRegRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_INVOICE_CLERK_REG
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
