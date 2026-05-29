/**
 * 开票员删除
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2InvoiceClerkDeleteRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    LoginAccount string `json:"login_account" structs:"login_account"` // 登录账号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2InvoiceClerkDeleteRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2InvoiceClerkDeleteRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2InvoiceClerkDeleteRequest(reqParam)
}

func (bp *BsPay) V2InvoiceClerkDeleteRequest(reqParam V2InvoiceClerkDeleteRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_INVOICE_CLERK_DELETE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
