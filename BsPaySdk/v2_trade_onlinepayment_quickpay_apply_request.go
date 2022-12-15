/**
 * 快捷支付申请
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeOnlinepaymentQuickpayApplyRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 订单金额
    CardBindId string `json:"card_bind_id" structs:"card_bind_id"` // 绑卡id
    NotifyUrl string `json:"notify_url" structs:"notify_url"` // 异步通知地址
    UserHuifuId string `json:"user_huifu_id" structs:"user_huifu_id"` // 用户客户号
    RiskCheckData string `json:"risk_check_data" structs:"risk_check_data"` // 安全信息
    TerminalDeviceData string `json:"terminal_device_data" structs:"terminal_device_data"` // 设备数据
    ExtendPayData string `json:"extend_pay_data" structs:"extend_pay_data"` // 银行扩展字段

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeOnlinepaymentQuickpayApplyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeOnlinepaymentQuickpayApplyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeOnlinepaymentQuickpayApplyRequest(reqParam)
}

func (bp *BsPay) V2TradeOnlinepaymentQuickpayApplyRequest(reqParam V2TradeOnlinepaymentQuickpayApplyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ONLINEPAYMENT_QUICKPAY_APPLY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
