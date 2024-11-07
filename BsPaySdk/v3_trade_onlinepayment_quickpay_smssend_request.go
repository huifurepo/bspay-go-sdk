/**
 * 快捷短信发送
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3TradeOnlinepaymentQuickpaySmssendRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    UserHuifuId string `json:"user_huifu_id" structs:"user_huifu_id"` // 用户客户号
    CardBindId string `json:"card_bind_id" structs:"card_bind_id"` // 绑卡id
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 订单金额
    NotifyUrl string `json:"notify_url" structs:"notify_url"` // 异步通知地址
    NuccData string `json:"nucc_data" structs:"nucc_data"` // 网联数据
    TerminalDeviceData string `json:"terminal_device_data" structs:"terminal_device_data"` // 设备数据
    RiskCheckData string `json:"risk_check_data" structs:"risk_check_data"` // 安全信息

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3TradeOnlinepaymentQuickpaySmssendRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3TradeOnlinepaymentQuickpaySmssendRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3TradeOnlinepaymentQuickpaySmssendRequest(reqParam)
}

func (bp *BsPay) V3TradeOnlinepaymentQuickpaySmssendRequest(reqParam V3TradeOnlinepaymentQuickpaySmssendRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_TRADE_ONLINEPAYMENT_QUICKPAY_SMSSEND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
