/**
 * 上传短信验证码
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2InvoiceMerVerifycodeUploadRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 开票方汇付ID
    VerifyType string `json:"verify_type" structs:"verify_type"` // 校验类型
    SerialNum string `json:"serial_num" structs:"serial_num"` // 流水号
    VerifyCode string `json:"verify_code" structs:"verify_code"` // 验证码

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2InvoiceMerVerifycodeUploadRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2InvoiceMerVerifycodeUploadRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2InvoiceMerVerifycodeUploadRequest(reqParam)
}

func (bp *BsPay) V2InvoiceMerVerifycodeUploadRequest(reqParam V2InvoiceMerVerifycodeUploadRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_INVOICE_MER_VERIFYCODE_UPLOAD
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
