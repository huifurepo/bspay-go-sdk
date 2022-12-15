/**
 * 交易结算对账文件重新生成
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeCheckReplayRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 交易日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付机构编号
    FileType string `json:"file_type" structs:"file_type"` // 文件类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeCheckReplayRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeCheckReplayRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeCheckReplayRequest(reqParam)
}

func (bp *BsPay) V2TradeCheckReplayRequest(reqParam V2TradeCheckReplayRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_CHECK_REPLAY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
