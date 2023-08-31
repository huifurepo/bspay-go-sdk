/**
 * 登记扣款信息
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePayscoreDeductRegitsterRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 商户申请单号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePayscoreDeductRegitsterRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePayscoreDeductRegitsterRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePayscoreDeductRegitsterRequest(reqParam)
}

func (bp *BsPay) V2TradePayscoreDeductRegitsterRequest(reqParam V2TradePayscoreDeductRegitsterRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYSCORE_DEDUCT_REGITSTER
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
