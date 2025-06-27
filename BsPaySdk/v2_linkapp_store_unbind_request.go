/**
 * 三方门店解绑
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2LinkappStoreUnbindRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    BindId string `json:"bind_id" structs:"bind_id"` // 绑定流水号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2LinkappStoreUnbindRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2LinkappStoreUnbindRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2LinkappStoreUnbindRequest(reqParam)
}

func (bp *BsPay) V2LinkappStoreUnbindRequest(reqParam V2LinkappStoreUnbindRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_LINKAPP_STORE_UNBIND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
