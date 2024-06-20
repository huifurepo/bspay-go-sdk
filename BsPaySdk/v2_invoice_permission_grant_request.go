/**
 * 电子发票业务开通
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2InvoicePermissionGrantRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 开票方汇付ID
    Status string `json:"status" structs:"status"` // 开通类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2InvoicePermissionGrantRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2InvoicePermissionGrantRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2InvoicePermissionGrantRequest(reqParam)
}

func (bp *BsPay) V2InvoicePermissionGrantRequest(reqParam V2InvoicePermissionGrantRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_INVOICE_PERMISSION_GRANT
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
