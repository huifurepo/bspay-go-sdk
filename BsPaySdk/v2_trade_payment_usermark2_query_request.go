/**
 * 获取银联用户标识接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePaymentUsermark2QueryRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    AuthCode string `json:"auth_code" structs:"auth_code"` // 支付授权码

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePaymentUsermark2QueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePaymentUsermark2QueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePaymentUsermark2QueryRequest(reqParam)
}

func (bp *BsPay) V2TradePaymentUsermark2QueryRequest(reqParam V2TradePaymentUsermark2QueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYMENT_USERMARK2_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
