/**
 * 发票邮件重发接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2InvoiceResendmailRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    IvcNumber string `json:"ivc_number" structs:"ivc_number"` // 发票号码
    MailAddr string `json:"mail_addr" structs:"mail_addr"` // 重发邮箱地址

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2InvoiceResendmailRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2InvoiceResendmailRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2InvoiceResendmailRequest(reqParam)
}

func (bp *BsPay) V2InvoiceResendmailRequest(reqParam V2InvoiceResendmailRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_INVOICE_RESENDMAIL
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
