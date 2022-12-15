/**
 * 微信代发
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeTransWxSurrogateRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    OutHuifuId string `json:"out_huifu_id" structs:"out_huifu_id"` // 出账商户号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 代发金额
    OpenId string `json:"open_id" structs:"open_id"` // 收款用户openid
    UserName string `json:"user_name" structs:"user_name"` // 微信收款用户姓名
    Remark string `json:"remark" structs:"remark"` // 代发备注

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeTransWxSurrogateRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeTransWxSurrogateRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeTransWxSurrogateRequest(reqParam)
}

func (bp *BsPay) V2TradeTransWxSurrogateRequest(reqParam V2TradeTransWxSurrogateRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_TRANS_WX_SURROGATE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
