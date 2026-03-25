/**
 * 账单计划变更
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3BillpayPlanModifyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    PlanNo string `json:"plan_no" structs:"plan_no"` // 账单计划编号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3BillpayPlanModifyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3BillpayPlanModifyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3BillpayPlanModifyRequest(reqParam)
}

func (bp *BsPay) V3BillpayPlanModifyRequest(reqParam V3BillpayPlanModifyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_BILLPAY_PLAN_MODIFY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
