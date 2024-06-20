/**
 * 电子发票业务开通查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2InvoicePermissionQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付Id
    IncludeSubFlag string `json:"include_sub_flag" structs:"include_sub_flag"` // 是否包含下级
    PageNum string `json:"page_num" structs:"page_num"` // 当前页
    PageSize string `json:"page_size" structs:"page_size"` // 分页大小

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2InvoicePermissionQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2InvoicePermissionQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2InvoicePermissionQueryRequest(reqParam)
}

func (bp *BsPay) V2InvoicePermissionQueryRequest(reqParam V2InvoicePermissionQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_INVOICE_PERMISSION_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
