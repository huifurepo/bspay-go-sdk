/**
 * 门店绑定列表查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2LinkappStoreQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    PlatformType string `json:"platform_type" structs:"platform_type"` // 平台类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2LinkappStoreQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2LinkappStoreQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2LinkappStoreQueryRequest(reqParam)
}

func (bp *BsPay) V2LinkappStoreQueryRequest(reqParam V2LinkappStoreQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_LINKAPP_STORE_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
