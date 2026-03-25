/**
 * 创建账单计划
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3BillpayPlanAddRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    ProjectNo string `json:"project_no" structs:"project_no"` // 账单项目编号
    PlanCycle string `json:"plan_cycle" structs:"plan_cycle"` // 账单周期
    BillDay string `json:"bill_day" structs:"bill_day"` // 账单日
    ReissueBillFlag string `json:"reissue_bill_flag" structs:"reissue_bill_flag"` // 补发当前周期账单标志枚举:Y-是、N-否；指定账单日时，必填；若填写是，则立即生成当前系统时间所在周期的账单； 滚动账单日时，此字段无效
    WithholdInfoData string `json:"withhold_info_data" structs:"withhold_info_data"` // 代扣信息jsonObject格式；账单计划需自动代扣时必填
    UserDocInfoList string `json:"user_doc_info_list" structs:"user_doc_info_list"` // 用户资料信息列表
    PaymentInfoList string `json:"payment_info_list" structs:"payment_info_list"` // 账单收费项信息列表

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3BillpayPlanAddRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3BillpayPlanAddRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3BillpayPlanAddRequest(reqParam)
}

func (bp *BsPay) V3BillpayPlanAddRequest(reqParam V3BillpayPlanAddRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_BILLPAY_PLAN_ADD
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
