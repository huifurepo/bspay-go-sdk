/**
 * 三方门店绑定（二阶段）
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2LinkappStoreBindRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户号
    PlatformType string `json:"platform_type" structs:"platform_type"` // 平台类型
    OpenShopUuid string `json:"open_shop_uuid" structs:"open_shop_uuid"` // 门店ID

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2LinkappStoreBindRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2LinkappStoreBindRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2LinkappStoreBindRequest(reqParam)
}

func (bp *BsPay) V2LinkappStoreBindRequest(reqParam V2LinkappStoreBindRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_LINKAPP_STORE_BIND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
