/**
 * 申请开票
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2HycInvoiceApplyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付id
    BatchList string `json:"batch_list" structs:"batch_list"` // 交易流水列表
    ReceiveMobile string `json:"receive_mobile" structs:"receive_mobile"` // 接收人手机号
    ReceiveName string `json:"receive_name" structs:"receive_name"` // 接收人姓名
    CourierAddress string `json:"courier_address" structs:"courier_address"` // 快递地址
    InvoiceCategory string `json:"invoice_category" structs:"invoice_category"` // 开票类目

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2HycInvoiceApplyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2HycInvoiceApplyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2HycInvoiceApplyRequest(reqParam)
}

func (bp *BsPay) V2HycInvoiceApplyRequest(reqParam V2HycInvoiceApplyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_HYC_INVOICE_APPLY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
