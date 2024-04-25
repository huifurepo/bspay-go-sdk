/**
 * 钱包关单
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2WalletTradeOrderCloseRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    OrgHfSeqId string `json:"org_hf_seq_id" structs:"org_hf_seq_id"` // 原交易全局流水号org_hf_seq_id，org_req_seq_id二选一；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：00470topo1A221019132207P068ac1362af00000&lt;/font&gt;
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原交易请求流水号org_hf_seq_id，org_req_seq_id二选一；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：rQ2021121311173944134649875651&lt;/font&gt;
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原交易请求日期

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2WalletTradeOrderCloseRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2WalletTradeOrderCloseRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2WalletTradeOrderCloseRequest(reqParam)
}

func (bp *BsPay) V2WalletTradeOrderCloseRequest(reqParam V2WalletTradeOrderCloseRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_WALLET_TRADE_ORDER_CLOSE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
