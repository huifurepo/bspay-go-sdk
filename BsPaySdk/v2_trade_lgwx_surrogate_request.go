/**
 * 灵工微信代发
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeLgwxSurrogateRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 出款方商户号
    CashAmt string `json:"cash_amt" structs:"cash_amt"` // 支付金额(元)
    SalaryModleType string `json:"salary_modle_type" structs:"salary_modle_type"` // 代发模式
    BmemberId string `json:"bmember_id" structs:"bmember_id"` // 落地公司商户号
    SubAppid string `json:"sub_appid" structs:"sub_appid"` // 子商户应用ID
    NotifyUrl string `json:"notify_url" structs:"notify_url"` // 异步通知地址
    AcctSplitBunch string `json:"acct_split_bunch" structs:"acct_split_bunch"` // 分账明细

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeLgwxSurrogateRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeLgwxSurrogateRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeLgwxSurrogateRequest(reqParam)
}

func (bp *BsPay) V2TradeLgwxSurrogateRequest(reqParam V2TradeLgwxSurrogateRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_LGWX_SURROGATE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
