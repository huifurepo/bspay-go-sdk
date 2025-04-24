/**
 * 银联统一在线收银台签约接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeOnlinepaymentUnionsignRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    PayScene string `json:"pay_scene" structs:"pay_scene"` // 支付场景
    NotifyUrl string `json:"notify_url" structs:"notify_url"` // 异步通知地址
    TerminalDeviceData string `json:"terminal_device_data" structs:"terminal_device_data"` // 设备信息
    ThirdPayData string `json:"third_pay_data" structs:"third_pay_data"` // 三方支付数据jsonObject；&lt;br/&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeOnlinepaymentUnionsignRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeOnlinepaymentUnionsignRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeOnlinepaymentUnionsignRequest(reqParam)
}

func (bp *BsPay) V2TradeOnlinepaymentUnionsignRequest(reqParam V2TradeOnlinepaymentUnionsignRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ONLINEPAYMENT_UNIONSIGN
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
