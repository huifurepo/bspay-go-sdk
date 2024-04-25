/**
 * 开通下级商户权限配置接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBusiHeadConfigRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    ProductId string `json:"product_id" structs:"product_id"` // 产品编号
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 直属渠道号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBusiHeadConfigRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBusiHeadConfigRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBusiHeadConfigRequest(reqParam)
}

func (bp *BsPay) V2MerchantBusiHeadConfigRequest(reqParam V2MerchantBusiHeadConfigRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BUSI_HEAD_CONFIG
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
