/**
 * 手机WAP支付
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeOnlinepaymentWappayRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 交易金额
    ExtendPayData string `json:"extend_pay_data" structs:"extend_pay_data"` // 网联扩展数据
    RiskCheckData string `json:"risk_check_data" structs:"risk_check_data"` // 安全信息
    TerminalDeviceData string `json:"terminal_device_data" structs:"terminal_device_data"` // 设备信息
    FrontUrl string `json:"front_url" structs:"front_url"` // 页面跳转地址
    NotifyUrl string `json:"notify_url" structs:"notify_url"` // 异步通知地址

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeOnlinepaymentWappayRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeOnlinepaymentWappayRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeOnlinepaymentWappayRequest(reqParam)
}

func (bp *BsPay) V2TradeOnlinepaymentWappayRequest(reqParam V2TradeOnlinepaymentWappayRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ONLINEPAYMENT_WAPPAY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc, false)
}
