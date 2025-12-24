/**
 * 创建批量账单数据
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3BillpayOrderBatchAddRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    ProjectNo string `json:"project_no" structs:"project_no"` // 账单项目编号
    UserDocInfoList string `json:"user_doc_info_list" structs:"user_doc_info_list"` // 用户资料信息列表
    PaymentInfoList string `json:"payment_info_list" structs:"payment_info_list"` // 账单收费项信息列表

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3BillpayOrderBatchAddRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3BillpayOrderBatchAddRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3BillpayOrderBatchAddRequest(reqParam)
}

func (bp *BsPay) V3BillpayOrderBatchAddRequest(reqParam V3BillpayOrderBatchAddRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_BILLPAY_ORDER_BATCH_ADD
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
