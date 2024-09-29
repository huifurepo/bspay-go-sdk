/**
 * 申诉文件url查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantAppealFileurlQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    FileId string `json:"file_id" structs:"file_id"` // 申诉文件Id

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantAppealFileurlQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantAppealFileurlQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantAppealFileurlQueryRequest(reqParam)
}

func (bp *BsPay) V2MerchantAppealFileurlQueryRequest(reqParam V2MerchantAppealFileurlQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_APPEAL_FILEURL_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
