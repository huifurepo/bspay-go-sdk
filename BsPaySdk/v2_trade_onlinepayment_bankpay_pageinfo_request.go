/**
 * 网银支付
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeOnlinepaymentBankpayPageinfoRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 订单金额
    GoodsDesc string `json:"goods_desc" structs:"goods_desc"` // 商品描述
    ExtendPayData string `json:"extend_pay_data" structs:"extend_pay_data"` // 网联扩展数据
    TerminalDeviceData string `json:"terminal_device_data" structs:"terminal_device_data"` // 设备信息
    RiskCheckData string `json:"risk_check_data" structs:"risk_check_data"` // 安全信息
    NotifyUrl string `json:"notify_url" structs:"notify_url"` // 异步通知地址

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeOnlinepaymentBankpayPageinfoRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeOnlinepaymentBankpayPageinfoRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeOnlinepaymentBankpayPageinfoRequest(reqParam)
}

func (bp *BsPay) V2TradeOnlinepaymentBankpayPageinfoRequest(reqParam V2TradeOnlinepaymentBankpayPageinfoRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ONLINEPAYMENT_BANKPAY_PAGEINFO
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc, false)
}
