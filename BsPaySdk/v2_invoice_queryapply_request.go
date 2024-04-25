/**
 * 发票开具申请查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2InvoiceQueryapplyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    ChannelId string `json:"channel_id" structs:"channel_id"` // 渠道号汇付商户号为空时，必传；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000109812124&lt;/font&gt;
    SeqId string `json:"seq_id" structs:"seq_id"` // 流水号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2InvoiceQueryapplyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2InvoiceQueryapplyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2InvoiceQueryapplyRequest(reqParam)
}

func (bp *BsPay) V2InvoiceQueryapplyRequest(reqParam V2InvoiceQueryapplyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_INVOICE_QUERYAPPLY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
