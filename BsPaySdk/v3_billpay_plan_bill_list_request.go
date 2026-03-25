/**
 * 查询账单计划下已生成账单数据
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3BillpayPlanBillListRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 客户请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 客户请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    PlanNo string `json:"plan_no" structs:"plan_no"` // 账单计划编号
    PageNum string `json:"page_num" structs:"page_num"` // 页码

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3BillpayPlanBillListRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3BillpayPlanBillListRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3BillpayPlanBillListRequest(reqParam)
}

func (bp *BsPay) V3BillpayPlanBillListRequest(reqParam V3BillpayPlanBillListRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_BILLPAY_PLAN_BILL_LIST
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
