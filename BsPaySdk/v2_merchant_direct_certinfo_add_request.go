/**
 * 证书登记
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantDirectCertinfoAddRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 渠道商汇付Id
    PayWay string `json:"pay_way" structs:"pay_way"` // 开通类型
    AppId string `json:"app_id" structs:"app_id"` // 开发者的应用ID
    FileList string `json:"file_list" structs:"file_list"` // 文件列表

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantDirectCertinfoAddRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantDirectCertinfoAddRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantDirectCertinfoAddRequest(reqParam)
}

func (bp *BsPay) V2MerchantDirectCertinfoAddRequest(reqParam V2MerchantDirectCertinfoAddRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_DIRECT_CERTINFO_ADD
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
