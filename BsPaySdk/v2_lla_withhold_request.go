/**
 * 代运营佣金代扣
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2LlaWithholdRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    AgencyHuifuId string `json:"agency_huifu_id" structs:"agency_huifu_id"` // 代运营汇付id
    MerchantHuifuId string `json:"merchant_huifu_id" structs:"merchant_huifu_id"` // 商家汇付id
    PlatformType string `json:"platform_type" structs:"platform_type"` // 平台
    EncashSeqId string `json:"encash_seq_id" structs:"encash_seq_id"` // 提现id
    TokenNo string `json:"token_no" structs:"token_no"` // 绑卡id
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 抽佣金额
    ExtendPayData string `json:"extend_pay_data" structs:"extend_pay_data"` // 银行扩展数据
    TerminalDeviceData string `json:"terminal_device_data" structs:"terminal_device_data"` // 设备信息
    RiskCheckData string `json:"risk_check_data" structs:"risk_check_data"` // 安全信息

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2LlaWithholdRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2LlaWithholdRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2LlaWithholdRequest(reqParam)
}

func (bp *BsPay) V2LlaWithholdRequest(reqParam V2LlaWithholdRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_LLA_WITHHOLD
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
