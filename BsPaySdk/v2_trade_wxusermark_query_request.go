/**
 * 微信用户标识查询接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeWxusermarkQueryRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    AuthCode string `json:"auth_code" structs:"auth_code"` // 支付授权码

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeWxusermarkQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeWxusermarkQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeWxusermarkQueryRequest(reqParam)
}

func (bp *BsPay) V2TradeWxusermarkQueryRequest(reqParam V2TradeWxusermarkQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_WXUSERMARK_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
