/**
 * 不明来账处理
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePaymentZxeUnknownincomeDisposeRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    BankSerialNo string `json:"bank_serial_no" structs:"bank_serial_no"` // 银行侧交易流水号
    OperateType string `json:"operate_type" structs:"operate_type"` // 操作类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePaymentZxeUnknownincomeDisposeRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePaymentZxeUnknownincomeDisposeRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePaymentZxeUnknownincomeDisposeRequest(reqParam)
}

func (bp *BsPay) V2TradePaymentZxeUnknownincomeDisposeRequest(reqParam V2TradePaymentZxeUnknownincomeDisposeRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYMENT_ZXE_UNKNOWNINCOME_DISPOSE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
