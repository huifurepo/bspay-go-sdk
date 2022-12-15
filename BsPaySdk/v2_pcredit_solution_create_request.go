/**
 * 创建花呗分期方案
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2PcreditSolutionCreateRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    ActivityName string `json:"activity_name" structs:"activity_name"` // 花呗分期商家贴息活动名称
    StartTime string `json:"start_time" structs:"start_time"` // 活动开始时间
    EndTime string `json:"end_time" structs:"end_time"` // 活动结束时间
    MinMoneyLimit string `json:"min_money_limit" structs:"min_money_limit"` // 免息金额下限(元)
    MaxMoneyLimit string `json:"max_money_limit" structs:"max_money_limit"` // 免息金额上限(元)
    AmountBudget string `json:"amount_budget" structs:"amount_budget"` // 花呗分期贴息预算金额
    InstallNumStrList string `json:"install_num_str_list" structs:"install_num_str_list"` // 花呗分期数集合
    BudgetWarningMoney string `json:"budget_warning_money" structs:"budget_warning_money"` // 预算提醒金额(元)
    BudgetWarningMailList string `json:"budget_warning_mail_list" structs:"budget_warning_mail_list"` // 预算提醒邮件列表
    BudgetWarningMobileNoList string `json:"budget_warning_mobile_no_list" structs:"budget_warning_mobile_no_list"` // 预算提醒手机号列表
    SubShopInfoList string `json:"sub_shop_info_list" structs:"sub_shop_info_list"` // 子门店信息集合

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2PcreditSolutionCreateRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2PcreditSolutionCreateRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2PcreditSolutionCreateRequest(reqParam)
}

func (bp *BsPay) V2PcreditSolutionCreateRequest(reqParam V2PcreditSolutionCreateRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_PCREDIT_SOLUTION_CREATE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
