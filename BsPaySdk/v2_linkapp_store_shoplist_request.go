/**
 * 三方门店查询（三方门店绑定一阶段）
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2LinkappStoreShoplistRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    PlatformType string `json:"platform_type" structs:"platform_type"` // 平台类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2LinkappStoreShoplistRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2LinkappStoreShoplistRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2LinkappStoreShoplistRequest(reqParam)
}

func (bp *BsPay) V2LinkappStoreShoplistRequest(reqParam V2LinkappStoreShoplistRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_LINKAPP_STORE_SHOPLIST
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
