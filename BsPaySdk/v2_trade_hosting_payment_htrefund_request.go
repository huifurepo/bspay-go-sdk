/**
 * 托管交易退款
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeHostingPaymentHtrefundRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    OrdAmt string `json:"ord_amt" structs:"ord_amt"` // 申请退款金额
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原交易请求日期
    RiskCheckData string `json:"risk_check_data" structs:"risk_check_data"` // 安全信息线上交易退款必填，参见线上退款接口；jsonObject字符串
    TerminalDeviceData string `json:"terminal_device_data" structs:"terminal_device_data"` // 设备信息线上交易退款必填，参见线上退款接口；jsonObject字符串

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeHostingPaymentHtrefundRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeHostingPaymentHtrefundRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeHostingPaymentHtrefundRequest(reqParam)
}

func (bp *BsPay) V2TradeHostingPaymentHtrefundRequest(reqParam V2TradeHostingPaymentHtrefundRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_HOSTING_PAYMENT_HTREFUND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
