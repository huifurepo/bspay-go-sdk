/**
 * 红字发票开具接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2InvoiceRedopenRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    OriIvcNumber string `json:"ori_ivc_number" structs:"ori_ivc_number"` // 原发票号码
    RedApplyReason string `json:"red_apply_reason" structs:"red_apply_reason"` // 红冲原因
    RedApplySource string `json:"red_apply_source" structs:"red_apply_source"` // 红冲申请来源

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2InvoiceRedopenRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2InvoiceRedopenRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2InvoiceRedopenRequest(reqParam)
}

func (bp *BsPay) V2InvoiceRedopenRequest(reqParam V2InvoiceRedopenRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_INVOICE_REDOPEN
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
