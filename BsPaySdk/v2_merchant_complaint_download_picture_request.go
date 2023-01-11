/**
 * 图片下载
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantComplaintDownloadPictureRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    MediaUrl string `json:"media_url" structs:"media_url"` // 下载图片的url
    MchId string `json:"mch_id" structs:"mch_id"` // 微信商户号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantComplaintDownloadPictureRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantComplaintDownloadPictureRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantComplaintDownloadPictureRequest(reqParam)
}

func (bp *BsPay) V2MerchantComplaintDownloadPictureRequest(reqParam V2MerchantComplaintDownloadPictureRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_COMPLAINT_DOWNLOAD_PICTURE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
