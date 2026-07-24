/**
 * 电子合同申请列表查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2InvoiceQueryecmsapplylistRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    PageNum string `json:"page_num" structs:"page_num"` // 当前页
    PageSize string `json:"page_size" structs:"page_size"` // 分页大小

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2InvoiceQueryecmsapplylistRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2InvoiceQueryecmsapplylistRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2InvoiceQueryecmsapplylistRequest(reqParam)
}

func (bp *BsPay) V2InvoiceQueryecmsapplylistRequest(reqParam V2InvoiceQueryecmsapplylistRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_INVOICE_QUERYECMSAPPLYLIST
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
