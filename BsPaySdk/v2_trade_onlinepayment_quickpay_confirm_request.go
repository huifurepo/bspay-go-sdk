/**
 * 快捷支付确认
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeOnlinepaymentQuickpayConfirmRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    SmsCode string `json:"sms_code" structs:"sms_code"` // 短信验证码
    NotifyUrl string `json:"notify_url" structs:"notify_url"` // 外部地址
    GoodsDesc string `json:"goods_desc" structs:"goods_desc"` // 商品描述

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeOnlinepaymentQuickpayConfirmRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeOnlinepaymentQuickpayConfirmRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeOnlinepaymentQuickpayConfirmRequest(reqParam)
}

func (bp *BsPay) V2TradeOnlinepaymentQuickpayConfirmRequest(reqParam V2TradeOnlinepaymentQuickpayConfirmRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ONLINEPAYMENT_QUICKPAY_CONFIRM
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
