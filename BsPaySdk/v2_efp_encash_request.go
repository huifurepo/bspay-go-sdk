/**
 * 全渠道资金提现申请
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2EfpEncashRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付id
    CashAmt string `json:"cash_amt" structs:"cash_amt"` // 交易金额.单位:元，2位小数
    TokenNo string `json:"token_no" structs:"token_no"` // 取现卡序列号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2EfpEncashRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2EfpEncashRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2EfpEncashRequest(reqParam)
}

func (bp *BsPay) V2EfpEncashRequest(reqParam V2EfpEncashRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_EFP_ENCASH
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
