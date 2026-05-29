/**
 * 设置默认开票员
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2InvoiceClerkDefaultModifyRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    LoginAccount string `json:"login_account" structs:"login_account"` // 登录账号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2InvoiceClerkDefaultModifyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2InvoiceClerkDefaultModifyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2InvoiceClerkDefaultModifyRequest(reqParam)
}

func (bp *BsPay) V2InvoiceClerkDefaultModifyRequest(reqParam V2InvoiceClerkDefaultModifyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_INVOICE_CLERK_DEFAULT_MODIFY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
