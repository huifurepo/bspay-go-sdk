/**
 * 查询账单计划详情
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3BillpayPlanDetailRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    PlanNo string `json:"plan_no" structs:"plan_no"` // 账单计划编号与原请求流水号编号二选一必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：BP202412270001&lt;/font&gt;
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原请求流水号原请求流水号，同一商户号当天唯一；与账单计划编号二选一必填
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原请求日期原请求日期格式：yyyyMMdd，以北京时间为准；与账单编号二选一必填

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3BillpayPlanDetailRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3BillpayPlanDetailRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3BillpayPlanDetailRequest(reqParam)
}

func (bp *BsPay) V3BillpayPlanDetailRequest(reqParam V3BillpayPlanDetailRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_BILLPAY_PLAN_DETAIL
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
