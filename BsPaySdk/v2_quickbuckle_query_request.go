/**
 * 快捷绑卡查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2QuickbuckleQueryRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户Id
    OutCustId string `json:"out_cust_id" structs:"out_cust_id"` // 用户id

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2QuickbuckleQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2QuickbuckleQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2QuickbuckleQueryRequest(reqParam)
}

func (bp *BsPay) V2QuickbuckleQueryRequest(reqParam V2QuickbuckleQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_QUICKBUCKLE_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
