/**
 * 分期单创建
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePayafteruseInstallmentCreateRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    FqAmt string `json:"fq_amt" structs:"fq_amt"` // 分期金额

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePayafteruseInstallmentCreateRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePayafteruseInstallmentCreateRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePayafteruseInstallmentCreateRequest(reqParam)
}

func (bp *BsPay) V2TradePayafteruseInstallmentCreateRequest(reqParam V2TradePayafteruseInstallmentCreateRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYAFTERUSE_INSTALLMENT_CREATE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
