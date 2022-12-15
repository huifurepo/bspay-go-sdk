/**
 * 交易结算对账单查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeCheckFilequeryRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    FileDate string `json:"file_date" structs:"file_date"` // 文件生成日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeCheckFilequeryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeCheckFilequeryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeCheckFilequeryRequest(reqParam)
}

func (bp *BsPay) V2TradeCheckFilequeryRequest(reqParam V2TradeCheckFilequeryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_CHECK_FILEQUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
