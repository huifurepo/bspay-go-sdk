/**
 * 资金归集明细查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeSettleCollectionDetailQueryRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    BeginDate string `json:"begin_date" structs:"begin_date"` // 开始日期
    EndDate string `json:"end_date" structs:"end_date"` // 结束日期
    OutHuifuId string `json:"out_huifu_id" structs:"out_huifu_id"` // 转出方商户号转出方商户号和转入方商户号必填一个；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123124&lt;/font&gt;
    InHuifuId string `json:"in_huifu_id" structs:"in_huifu_id"` // 转入方商户号转出方商户号和转入方商户号必填一个；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123124&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeSettleCollectionDetailQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeSettleCollectionDetailQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeSettleCollectionDetailQueryRequest(reqParam)
}

func (bp *BsPay) V2TradeSettleCollectionDetailQueryRequest(reqParam V2TradeSettleCollectionDetailQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_SETTLE_COLLECTION_DETAIL_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
